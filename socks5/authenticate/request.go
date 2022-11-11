package authenticate

import (
	"bytes"
	"fmt"
	"io"
)

const (
	LENGTH_VER      = 1
	LENGTH_NMETHODS = 1
	// LENGTH_METHODS  = 1 ~ 255
)

type Request struct {
	VER      byte
	NMETHODS byte
	METHODS  []byte
}

func (r *Request) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.WriteByte(r.VER)
	buf.WriteByte(r.NMETHODS)
	buf.Write(r.METHODS)
	return buf.Bytes(), nil
}

func (r *Request) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// VER
	buf := make([]byte, LENGTH_VER)
	n, err := io.ReadFull(reader, buf)
	if n != LENGTH_VER || err != nil {
		return fmt.Errorf("failed to read ver:  %s", err)
	}
	r.VER = buf[0]

	// NMETHODS
	buf = make([]byte, LENGTH_NMETHODS)
	n, err = io.ReadFull(reader, buf)
	if n != LENGTH_NMETHODS || err != nil {
		return fmt.Errorf("failed to read nmethods:  %s", err)
	}
	r.NMETHODS = buf[0]

	// METHODS
	LENGTH_METHODS := int(r.NMETHODS)
	buf = make([]byte, LENGTH_METHODS)
	n, err = io.ReadFull(reader, buf)
	if n != LENGTH_METHODS || err != nil {
		return fmt.Errorf("failed to read METHODS:  %s", err)
	}
	r.METHODS = buf

	return nil
}
