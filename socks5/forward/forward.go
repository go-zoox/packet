package forward

import (
	"bytes"
	"fmt"
	"io"
)

// Forward ...
type Forward struct {
	DATA []byte
}

// Encode encodes the forward data
func (r *Forward) Encode() ([]byte, error) {
	return r.DATA, nil
}

// Decode decodes the forward data
func (r *Forward) Decode(raw []byte) (err error) {
	reader := bytes.NewReader(raw)

	// DATA
	r.DATA, err = io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("failed to read DATA:  %s", err)
	}

	return nil
}
