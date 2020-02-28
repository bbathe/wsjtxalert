package main

import (
	"encoding/binary"
	"io"
	"log"
)

// ReadQUint32 deserializes a Quint32 from io.Reader
func ReadQUint32(r io.Reader) (uint32, error) {
	var u uint32

	err := binary.Read(r, binary.BigEndian, &u)
	if err != nil {
		log.Printf("%+v", err)
		return 0, err
	}
	return u, nil
}

// WriteQUint32 serializes a Quint32 to io.Writer
func WriteQUint32(w io.Writer, u uint32) error {
	err := binary.Write(w, binary.BigEndian, &u)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}
	return nil
}

// ReadQInt32 deserializes a Qint32 from io.Reader
func ReadQInt32(r io.Reader) (int32, error) {
	var i int32

	err := binary.Read(r, binary.BigEndian, &i)
	if err != nil {
		log.Printf("%+v", err)
		return 0, err
	}
	return i, nil
}

// WriteQInt32 serializes a Qint32 to io.Writer
func WriteQInt32(w io.Writer, t int32) error {
	err := binary.Write(w, binary.BigEndian, &t)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}
	return nil
}

// ReadQUtf8 deserializes a Qutf8 from io.Reader
func ReadQUtf8(r io.Reader) (string, error) {
	var n uint32

	// get string size to read
	err := binary.Read(r, binary.BigEndian, &n)
	if err != nil {
		log.Printf("%+v", err)
		return "", err
	}

	// empty  strings have a length of zero and null strings have a length field of 0xffffffff
	if n == 0xffffffff || n == 0 {
		return "", nil
	}

	// get string bytes
	bs := make([]byte, n)
	_, err = io.ReadFull(r, bs)
	if err != nil {
		log.Printf("%+v", err)
		return "", err
	}
	return string(bs), nil
}

// WriteQUtf8 serializes a QUtf8 to io.Writer
func WriteQUtf8(w io.Writer, s string) error {
	n := uint32(len(s))

	err := binary.Write(w, binary.BigEndian, &n)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	b := []byte(s[:n])

	err = binary.Write(w, binary.BigEndian, &b)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	return nil
}

// ReadQBool deserializes a Qbool from io.Reader
func ReadQBool(r io.Reader) (bool, error) {
	var b uint8

	err := binary.Read(r, binary.BigEndian, &b)
	if err != nil {
		log.Printf("%+v", err)
		return false, err
	}
	return b != 0, nil
}

// WriteQBool serializes a QBool to io.Writer
func WriteQBool(w io.Writer, b bool) error {
	bs := uint8(0)
	if b {
		bs = 1
	}

	err := binary.Write(w, binary.BigEndian, &bs)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}
	return nil
}

// ReadQUint8 deserializes a QUint8 from io.Reader
func ReadQUint8(r io.Reader) (uint8, error) {
	var u uint8

	err := binary.Read(r, binary.BigEndian, &u)
	if err != nil {
		log.Printf("%+v", err)
		return 0, err
	}
	return u, nil
}

// WriteQUint8 serializes a QUint8 to io.Writer
func WriteQUint8(w io.Writer, u uint8) error {
	err := binary.Write(w, binary.BigEndian, &u)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}
	return nil
}

// ReadQFloat deserializes a Qfloat from io.Reader
func ReadQFloat(r io.Reader) (float64, error) {
	var f float64

	err := binary.Read(r, binary.BigEndian, &f)
	if err != nil {
		log.Printf("%+v", err)
		return 0, err
	}
	return f, nil
}

// WriteQFloat serializes a QFloat to io.Writer
func WriteQFloat(w io.Writer, f float64) error {
	err := binary.Write(w, binary.BigEndian, &f)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}
	return nil
}

// ReadQTime deserializes a Qtime from io.Reader
func ReadQTime(r io.Reader) (uint32, error) {
	var t uint32

	err := binary.Read(r, binary.BigEndian, &t)
	if err != nil {
		log.Printf("%+v", err)
		return 0, err
	}
	return t, nil
}

// WriteQTime serializes a QTime to io.Writer
func WriteQTime(w io.Writer, t uint32) error {
	err := binary.Write(w, binary.BigEndian, &t)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}
	return nil
}
