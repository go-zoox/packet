package forward

import (
	"bytes"
	"fmt"
	"io"
)

// DATA Protocol:
//
// TRANSMISIION DATA:
// request:  CONNECTION_ID | DATA
//					       21      |  -

const (
	// LengthConnectionID ...
	LengthConnectionID = 21
)

// Forward ...
type Forward struct {
	ConnectionID string
	Data         []byte
}

// Encode encodes the forward data
func (r *Forward) Encode() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(r.ConnectionID)
	buf.Write(r.Data)
	return buf.Bytes(), nil
}

// Decode decodes the forward data
func (r *Forward) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// CONNECTION_ID
	buf := make([]byte, LengthConnectionID)
	n, err := io.ReadFull(reader, buf)
	if n != LengthConnectionID || err != nil {
		return fmt.Errorf("failed to read connection id:  %s", err)
	}
	r.ConnectionID = string(buf)

	r.Data, err = io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("failed to read data:  %s", err)
	}

	return nil
}
