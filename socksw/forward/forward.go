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
	LENGTH_CONNECTION_ID = 21
)

type Forward struct {
	CONNECTION_ID string
	DATA          []byte
}

func (r *Forward) Encode() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(r.CONNECTION_ID)
	buf.Write(r.DATA)
	return buf.Bytes(), nil
}

func (r *Forward) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// CONNECTION_ID
	buf := make([]byte, LENGTH_CONNECTION_ID)
	n, err := io.ReadFull(reader, buf)
	if n != LENGTH_CONNECTION_ID || err != nil {
		return fmt.Errorf("failed to read connection id:  %s", err)
	}
	r.CONNECTION_ID = string(buf)

	r.DATA, err = io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("failed to read data:  %s", err)
	}

	return nil
}
