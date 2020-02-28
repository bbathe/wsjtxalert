// +build windows

package main

import (
	"log"
	"syscall"

	"golang.org/x/sys/windows"
)

var (
	dll   = windows.NewLazySystemDLL("Kernel32.dll")
	pBeep = dll.NewProc("Beep")
)

// alertSound generates an aubible alert
func alertSound() {
	go func() {
		for i := 5; i > 0; i-- {
			r1, _, err := syscall.Syscall(pBeep.Addr(),
				2,
				uintptr(750),
				uintptr(300),
				0,
			)
			if r1 == 0 {
				log.Printf("%+v", err)
				return
			}
		}
	}()
}
