package data

import (
	"bytes"
	"fmt"
	"io"
)

type Request struct {
	DATA []byte
}

func (r *Request) Encode() ([]byte, error) {
	return r.DATA, nil
}

func (r *Request) Decode(raw []byte) (err error) {
	reader := bytes.NewReader(raw)

	// DATA
	r.DATA, err = io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("failed to read DATA:  %s", err)
	}

	return nil
}
