package close

import (
	"bytes"
	"fmt"
	"io"

	"github.com/go-zoox/packet/socksz"
)

// DATA Protocol:
//
// CONNECTION CLOSE DATA:
// request:  CONNECTION_ID
//                 13

// Close ...
type Close struct {
	ConnectionID string
}

// Encode encodes the data
func (r *Close) Encode() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	n, err := buf.WriteString(r.ConnectionID)
	if n != socksz.LengthConnectionID || err != nil {
		return nil, fmt.Errorf("failed to write ConnectionID: %s", err)
	}

	return buf.Bytes(), nil
}

// Decode decodes the data
func (r *Close) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// CONNECTION_ID
	buf := make([]byte, socksz.LengthConnectionID)
	n, err := io.ReadFull(reader, buf)
	if n != socksz.LengthConnectionID || err != nil {
		return fmt.Errorf("failed to read connection id:  %s", err)
	}
	r.ConnectionID = string(buf)

	return nil
}
