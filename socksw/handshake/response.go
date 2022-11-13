package handshake

import (
	"bytes"
	"fmt"
	"io"
)

// DATA Protocol:
//
// AUTHENTICATE DATA:
// response: STATUS | MESSAGE
//            1     |  -

const (
	// LengthStatus ...
	LengthStatus = 1
)

// Response represents the response
type Response struct {
	ConnectionID string
	Status       uint8
	Message      string
}

// Encode encodes the data
func (r *Response) Encode() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	n, err := buf.WriteString(r.ConnectionID)
	if n != LengthConnectionID || err != nil {
		return nil, fmt.Errorf("failed to write ConnectionID: %s", err)
	}

	err = buf.WriteByte(r.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to write Status: %s", err)
	}

	_, err = buf.WriteString(r.Message)
	if err != nil {
		return nil, fmt.Errorf("failed to write Message: %s", err)
	}

	return buf.Bytes(), nil
}

// Decode decodes the data
func (r *Response) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// CONNECTION_ID
	buf := make([]byte, LengthConnectionID)
	n, err := io.ReadFull(reader, buf)
	if n != LengthConnectionID || err != nil {
		return fmt.Errorf("failed to read id:  %s", err)
	}
	r.ConnectionID = string(buf)

	// STATUS
	buf = make([]byte, LengthStatus)
	n, err = io.ReadFull(reader, buf)
	if n != LengthStatus || err != nil {
		return fmt.Errorf("failed to read status:  %s", err)
	}
	r.Status = uint8(buf[0])

	// Message
	buf, err = io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("failed to read message:  %s", err)
	}
	r.Message = string(buf)

	return nil
}
