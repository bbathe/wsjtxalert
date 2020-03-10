package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

// configuration holds the application configuration
type configuration struct {
	WSJTXServer struct {
		Port string
		IP   string
	}
	Prefixes struct {
		Callsign   []string
		Gridsquare []string
	}
}

var (
	// application configuration
	config configuration
)

// HasAnyPrefix tests whether the string s begins with any of the prefixes
func HasAnyPrefix(s string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

func CompactSlice(s []string) []string {
	c := make([]string, 0, len(s))

	for _, e := range s {
		if len(e) > 0 {
			c = append(c, e)
		}
	}

	return c
}

func main() {
	// show file & location, date & time
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// log & config files are in the same directory as the executable with the same base name
	fn, err := os.Executable()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	basefn := strings.TrimSuffix(fn, path.Ext(fn))

	// log to file
	f, err := os.OpenFile(basefn+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	// read config
	// #nosec G304
	b, err := ioutil.ReadFile(basefn + ".yaml")
	if err != nil {
		log.Fatalf("%+v", err)
	}

	err = yaml.Unmarshal(b, &config)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// read from UDP port
	port, err := strconv.Atoi(config.WSJTXServer.Port)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	ip := net.ParseIP(config.WSJTXServer.IP)
	if ip == nil {
		log.Fatalf("%+v", err)
	}
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   ip,
		Port: port,
	})
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer conn.Close()

	msg := make([]byte, 10240)
	for {
		n, _, err := conn.ReadFromUDP(msg)
		if err != nil {
			log.Fatalf("%+v", err)
		}

		r := bytes.NewReader(msg[:n])

		// read as if it's a wsjt-x NetworkMessage
		var nm NetworkMessage
		err = nm.Read(r)
		if err != nil {
			log.Printf("%+v", err)
			continue
		}

		// from wsjt-x?
		if nm.Magic != 0xadbccbda {
			log.Print("failed magic test")
			continue
		}

		// check for schema version 2
		if nm.Schema != 2 {
			log.Print("wrong schema version")
			continue
		}

		// only handle Decode messages
		if nm.MessageType == 2 {
			var d Decode

			// reconstitute Decode message
			err = d.Read(r)
			if err != nil {
				log.Printf("%+v", err)
				continue
			}

			// process message for callsign & gridsquare
			f := CompactSlice(strings.Split(d.Message, " "))
			l := len(f)
			if l > 2 {
				// CQ KR0OT DN40
				// CQ DX K4PI EM73
				// DU3CQ N2NL EM50
				// KW0G KR0OT -02
				// W4EJY WW5SS RR73
				callerCallsign := f[l-2]
				callerGridsquare := f[l-1]

				// test callsign
				if HasAnyPrefix(callerCallsign, config.Prefixes.Callsign) {
					color.Red(fmt.Sprintf("%s matches callsign prefix alert", callerCallsign))
					alertSound()
				}

				// filter out invalid gridsquares
				_, err := strconv.Atoi(callerGridsquare[1:])
				if err != nil && callerGridsquare != "RR73" {
					// test gridsquare
					if HasAnyPrefix(callerGridsquare, config.Prefixes.Gridsquare) {
						color.Red(fmt.Sprintf("%s matches gridsquare prefix alert", callerGridsquare))
						alertSound()
					}
				}
			}
			continue
		}
	}
}
