package forward

import (
	"bytes"
	"fmt"
	"io"

	"github.com/go-zoox/packet/socksz"
)

// DATA Protocol:
//
// TRANSMISIION DATA:
// request:  CONNECTION_ID | DATA
//					       21      |  -

// Forward ...
type Forward struct {
	ConnectionID string
	Data         []byte
}

// Encode encodes the forward data
func (r *Forward) Encode() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	n, err := buf.WriteString(r.ConnectionID)
	if n != socksz.LengthConnectionID || err != nil {
		return nil, fmt.Errorf("failed to write ConnectionID: %s", err)
	}

	_, err = buf.Write(r.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to write Data: %s", err)
	}

	return buf.Bytes(), nil
}

// Decode decodes the forward data
func (r *Forward) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// CONNECTION_ID
	buf := make([]byte, socksz.LengthConnectionID)
	n, err := io.ReadFull(reader, buf)
	if n != socksz.LengthConnectionID || err != nil {
		return fmt.Errorf("failed to read connection id:  %s", err)
	}
	r.ConnectionID = string(buf)

	r.Data, err = io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("failed to read data:  %s", err)
	}

	return nil
}
