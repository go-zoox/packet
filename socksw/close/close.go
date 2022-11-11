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
	// LengthConnectionID ...
	LengthConnectionID = 21
)

// Close ...
type Close struct {
	ConnectionID string
}

// Encode encodes the data
func (r *Close) Encode() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(r.ConnectionID)
	return buf.Bytes(), nil
}

// Decode decodes the data
func (r *Close) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// CONNECTION_ID
	buf := make([]byte, LengthConnectionID)
	n, err := io.ReadFull(reader, buf)
	if n != LengthConnectionID || err != nil {
		return fmt.Errorf("failed to read connection id:  %s", err)
	}
	r.ConnectionID = string(buf)

	return nil
}
