package authenticate

import (
	"bytes"
	"fmt"
	"io"
)

const (
	// LengthMethod is the byte length of METHOD
	LengthMethod = 1
)

// Response is the response for authenticate
type Response struct {
	VER    byte
	METHOD byte
}

// Encode encodes the response for authenticate
func (r *Response) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.WriteByte(r.VER)
	buf.WriteByte(r.METHOD)
	return buf.Bytes(), nil
}

// Decode decodes the response for authenticate
func (r *Response) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// VER
	buf := make([]byte, LengthVer)
	n, err := io.ReadFull(reader, buf)
	if n != LengthVer || err != nil {
		return fmt.Errorf("failed to read ver:  %s", err)
	}
	r.VER = buf[0]

	// METHOD
	buf = make([]byte, LengthMethod)
	n, err = io.ReadFull(reader, buf)
	if n != LengthMethod || err != nil {
		return fmt.Errorf("failed to read method:  %s", err)
	}
	r.METHOD = buf[0]

	return nil
}
