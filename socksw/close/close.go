package close

import (
	"bytes"
	"fmt"
	"io"
)

// DATA Protocol:
//
// CONNECTION CLOSE DATA:
// request:  CONNECTION_ID
//                 21

const (
	LENGTH_CONNECTION_ID = 21
)

type Close struct {
	CONNECTION_ID string
}

func (r *Close) Encode() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(r.CONNECTION_ID)
	return buf.Bytes(), nil
}

func (r *Close) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// CONNECTION_ID
	buf := make([]byte, LENGTH_CONNECTION_ID)
	n, err := io.ReadFull(reader, buf)
	if n != LENGTH_CONNECTION_ID || err != nil {
		return fmt.Errorf("failed to read connection id:  %s", err)
	}
	r.CONNECTION_ID = string(buf)

	return nil
}
