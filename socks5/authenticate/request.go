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
	Ver      byte
	NMethods byte
	Methods  []byte
}

// Encode encodes the request
func (r *Request) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)

	err := buf.WriteByte(r.Ver)
	if err != nil {
		return nil, fmt.Errorf("failed to write Ver: %s", err)
	}

	err = buf.WriteByte(r.NMethods)
	if err != nil {
		return nil, fmt.Errorf("failed to write NMethods: %s", err)
	}

	var n int
	n, err = buf.Write(r.Methods)
	if n != int(r.NMethods) || err != nil {
		return nil, fmt.Errorf("failed to write Methods: %s", err)
	}

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
	r.Ver = buf[0]

	// NMETHODS
	buf = make([]byte, LengthNMethods)
	n, err = io.ReadFull(reader, buf)
	if n != LengthNMethods || err != nil {
		return fmt.Errorf("failed to read nmethods:  %s", err)
	}
	r.NMethods = buf[0]

	// METHODS
	LengthMethods := int(r.NMethods)
	buf = make([]byte, LengthMethods)
	n, err = io.ReadFull(reader, buf)
	if n != LengthMethods || err != nil {
		return fmt.Errorf("failed to read METHODS:  %s", err)
	}
	r.Methods = buf

	return nil
}
