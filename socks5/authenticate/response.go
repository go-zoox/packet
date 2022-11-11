package authenticate

import (
	"bytes"
	"fmt"
	"io"
)

const (
	LENGTH_METHOD = 1
)

type Response struct {
	VER    byte
	METHOD byte
}

func (r *Response) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.WriteByte(r.VER)
	buf.WriteByte(r.METHOD)
	return buf.Bytes(), nil
}

func (r *Response) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// VER
	buf := make([]byte, LENGTH_VER)
	n, err := io.ReadFull(reader, buf)
	if n != LENGTH_VER || err != nil {
		return fmt.Errorf("failed to read ver:  %s", err)
	}
	r.VER = buf[0]

	// METHOD
	buf = make([]byte, LENGTH_METHOD)
	n, err = io.ReadFull(reader, buf)
	if n != LENGTH_METHOD || err != nil {
		return fmt.Errorf("failed to read method:  %s", err)
	}
	r.METHOD = buf[0]

	return nil
}
