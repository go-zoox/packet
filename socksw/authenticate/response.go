package authenticate

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
	// LengthStatus is the byte length of STATUS
	LengthStatus = 1
)

// Response is the response for authenticate
type Response struct {
	STATUS  uint8
	MESSAGE string
}

// Encode encodes the response for authenticate
func (r *Response) Encode() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteByte(r.STATUS)
	buf.WriteString(r.MESSAGE)
	return buf.Bytes(), nil
}

// Decode decodes the response for authenticate
func (r *Response) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// STATUS
	buf := make([]byte, LengthStatus)
	n, err := io.ReadFull(reader, buf)
	if n != LengthStatus || err != nil {
		return fmt.Errorf("failed to read status:  %s", err)
	}
	r.STATUS = uint8(buf[0])

	// Message
	buf, err = io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("failed to read message:  %s", err)
	}
	r.MESSAGE = string(buf)

	return nil
}
