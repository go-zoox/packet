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
	Ver    byte
	Method byte
}

// Encode encodes the response for authenticate
func (r *Response) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)

	err := buf.WriteByte(r.Ver)
	if err != nil {
		return nil, fmt.Errorf("failed to write Ver: %s", err)
	}

	err = buf.WriteByte(r.Method)
	if err != nil {
		return nil, fmt.Errorf("failed to write Method: %s", err)
	}

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
	r.Ver = buf[0]

	// METHOD
	buf = make([]byte, LengthMethod)
	n, err = io.ReadFull(reader, buf)
	if n != LengthMethod || err != nil {
		return fmt.Errorf("failed to read method:  %s", err)
	}
	r.Method = buf[0]

	return nil
}
