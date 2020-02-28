package main

import (
	"io"
	"log"
)

// https://sourceforge.net/p/wsjt/wsjtx/ci/master/tree/NetworkMessage.hpp

// NetworkMessage defines messages from wsjt-x
// I moved MessageType here where the wsjt docs show it in the individual message definitions
type NetworkMessage struct {
	Magic       uint32
	Schema      uint32
	MessageType uint32
}

// Read deserializes a NetworkMessage from io.Reader
func (networkmessage *NetworkMessage) Read(r io.Reader) error {
	// read all values from stream
	magic, err := ReadQUint32(r)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	schema, err := ReadQUint32(r)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	messagetype, err := ReadQUint32(r)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	// set in object
	networkmessage.Magic = magic
	networkmessage.Schema = schema
	networkmessage.MessageType = messagetype

	return nil
}

// Write serializes a NetworkMessage to io.Writer
func (networkmessage *NetworkMessage) Write(w io.Writer) error {
	err := WriteQUint32(w, networkmessage.Magic)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	err = WriteQUint32(w, networkmessage.Schema)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	err = WriteQUint32(w, networkmessage.MessageType)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	return nil
}

// Decode defines Decode messages from wsjt-x
type Decode struct {
	ID             string
	New            bool
	Time           uint32
	SNR            int32
	DeltaTime      float64
	DeltaFrequency uint32
	Mode           string
	Message        string
	LowConfidence  bool
	OffAir         bool
}

// Read deserializes a Decode from io.Reader
func (decode *Decode) Read(r io.Reader) error {
	// read all values from stream
	id, err := ReadQUtf8(r)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	new, err := ReadQBool(r)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	tm, err := ReadQTime(r)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	snr, err := ReadQInt32(r)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	deltatime, err := ReadQFloat(r)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	deltafrequency, err := ReadQUint32(r)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	mode, err := ReadQUtf8(r)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	message, err := ReadQUtf8(r)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	lowconfidence, err := ReadQBool(r)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	offair, err := ReadQBool(r)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	// set in object
	decode.ID = id
	decode.New = new
	decode.Time = tm
	decode.SNR = snr
	decode.DeltaTime = deltatime
	decode.DeltaFrequency = deltafrequency
	decode.Mode = mode
	decode.Message = message
	decode.LowConfidence = lowconfidence
	decode.OffAir = offair

	return nil
}

type Reply struct {
	ID             string
	Time           uint32
	SNR            int32
	DeltaTime      float64
	DeltaFrequency uint32
	Mode           string
	Message        string
	LowConfidence  bool
	Modifiers      uint8
}

// Write serializes a Reply to io.Writer
func (reply *Reply) Write(w io.Writer) error {
	err := WriteQUtf8(w, reply.ID)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	err = WriteQTime(w, reply.Time)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	err = WriteQInt32(w, reply.SNR)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	err = WriteQFloat(w, reply.DeltaTime)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	err = WriteQUint32(w, reply.DeltaFrequency)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	err = WriteQUtf8(w, reply.Mode)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	err = WriteQUtf8(w, reply.Message)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	err = WriteQBool(w, reply.LowConfidence)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	err = WriteQUint8(w, reply.Modifiers)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	return nil
}
