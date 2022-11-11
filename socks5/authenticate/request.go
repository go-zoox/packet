package authenticate

import (
	"bytes"
	"fmt"
	"io"
)

const (
	// LengthVer is the byte length of VER
	LengthVer = 1
	// LengthNMethods is the byte length of NMETHODS
	LengthNMethods = 1
	// LENGTH_METHODS  = 1 ~ 255
)

// Request is the request for authenticate
type Request struct {
	VER      byte
	NMETHODS byte
	METHODS  []byte
}

// Encode encodes the request
func (r *Request) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.WriteByte(r.VER)
	buf.WriteByte(r.NMETHODS)
	buf.Write(r.METHODS)
	return buf.Bytes(), nil
}

// Decode decodes the request
func (r *Request) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// VER
	buf := make([]byte, LengthVer)
	n, err := io.ReadFull(reader, buf)
	if n != LengthVer || err != nil {
		return fmt.Errorf("failed to read ver:  %s", err)
	}
	r.VER = buf[0]

	// NMETHODS
	buf = make([]byte, LengthNMethods)
	n, err = io.ReadFull(reader, buf)
	if n != LengthNMethods || err != nil {
		return fmt.Errorf("failed to read nmethods:  %s", err)
	}
	r.NMETHODS = buf[0]

	// METHODS
	LengthMethods := int(r.NMETHODS)
	buf = make([]byte, LengthMethods)
	n, err = io.ReadFull(reader, buf)
	if n != LengthMethods || err != nil {
		return fmt.Errorf("failed to read METHODS:  %s", err)
	}
	r.METHODS = buf

	return nil
}
