package authenticate

import (
	"bytes"
	"fmt"
	"io"
)

// DATA Protocol:
//
// AUTHENTICATE DATA:
// request:  USER_CLIENT_ID | TIMESTAMP | NONCE | SIGNATURE
//             10           |    13     |   6   |  64 HMAC_SHA256

const (
	LENGTH_USER_CLIENT_ID = 10
	LENGTH_TIMESTAMP      = 13
	LENGTH_NONCE          = 6
	LENGTH_SIGNATURE      = 64
)

type Request struct {
	USER_CLIENT_ID string
	TIMESTAMP      string
	NONCE          string
	SIGNATURE      string
}

func (r *Request) Encode() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(r.USER_CLIENT_ID)
	buf.WriteString(r.TIMESTAMP)
	buf.WriteString(r.NONCE)
	buf.WriteString(r.SIGNATURE)
	return buf.Bytes(), nil
}

func (r *Request) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// USER_CLIENT_ID
	buf := make([]byte, LENGTH_USER_CLIENT_ID)
	n, err := io.ReadFull(reader, buf)
	if n != LENGTH_USER_CLIENT_ID || err != nil {
		return fmt.Errorf("failed to read user client id:  %s", err)
	}
	r.USER_CLIENT_ID = string(buf)

	// TIMESTAMP
	buf = make([]byte, LENGTH_TIMESTAMP)
	n, err = io.ReadFull(reader, buf)
	if n != LENGTH_TIMESTAMP || err != nil {
		return fmt.Errorf("failed to read timestamp:  %s", err)
	}
	r.TIMESTAMP = string(buf)

	// NONCE
	buf = make([]byte, LENGTH_NONCE)
	n, err = io.ReadFull(reader, buf)
	if n != LENGTH_NONCE || err != nil {
		return fmt.Errorf("failed to read nonce:  %s", err)
	}
	r.NONCE = string(buf)

	// SIGNATURE
	buf = make([]byte, LENGTH_SIGNATURE)
	n, err = io.ReadFull(reader, buf)
	if n != LENGTH_SIGNATURE || err != nil {
		return fmt.Errorf("failed to read signature:  %s", err)
	}
	r.SIGNATURE = string(buf)

	return nil
}
