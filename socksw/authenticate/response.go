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
	LENGTH_STATUS = 1
)

type Response struct {
	STATUS  uint8
	MESSAGE string
}

func (r *Response) Encode() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteByte(r.STATUS)
	buf.WriteString(r.MESSAGE)
	return buf.Bytes(), nil
}

func (r *Response) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// STATUS
	buf := make([]byte, LENGTH_STATUS)
	n, err := io.ReadFull(reader, buf)
	if n != LENGTH_STATUS || err != nil {
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