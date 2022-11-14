package authenticate

import (
	"bytes"
	"fmt"
	"io"

	"github.com/go-zoox/packet/socksz"
)

// DATA Protocol:
//
// AUTHENTICATE DATA:
// request:  USER_CLIENT_ID | TIMESTAMP | NONCE | SIGNATURE
//             10           |    13     |   6   |  64 HMAC_SHA256

// Request is the request for authenticate
type Request struct {
	UserClientID string
	Timestamp    string
	Nonce        string
	Signature    string
}

// Encode encodes the request
func (r *Request) Encode() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	n, err := buf.WriteString(r.UserClientID)
	if n != socksz.LengthUserClientID || err != nil {
		return nil, fmt.Errorf("failed to write user client id:  %s", err)
	}

	n, err = buf.WriteString(r.Timestamp)
	if n != socksz.LengthTimestamp || err != nil {
		return nil, fmt.Errorf("failed to write timestamp:  %s", err)
	}

	n, err = buf.WriteString(r.Nonce)
	if n != socksz.LengthNonce || err != nil {
		return nil, fmt.Errorf("failed to write nonce:  %s", err)
	}

	n, err = buf.WriteString(r.Signature)
	if n != socksz.LengthSignature || err != nil {
		return nil, fmt.Errorf("failed to write signature:  %s", err)
	}

	return buf.Bytes(), nil
}

// Decode decodes the request
func (r *Request) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// USER_CLIENT_ID
	buf := make([]byte, socksz.LengthUserClientID)
	n, err := io.ReadFull(reader, buf)
	if n != socksz.LengthUserClientID || err != nil {
		return fmt.Errorf("failed to read user client id:  %s", err)
	}
	r.UserClientID = string(buf)

	// TIMESTAMP
	buf = make([]byte, socksz.LengthTimestamp)
	n, err = io.ReadFull(reader, buf)
	if n != socksz.LengthTimestamp || err != nil {
		return fmt.Errorf("failed to read timestamp:  %s", err)
	}
	r.Timestamp = string(buf)

	// NONCE
	buf = make([]byte, socksz.LengthNonce)
	n, err = io.ReadFull(reader, buf)
	if n != socksz.LengthNonce || err != nil {
		return fmt.Errorf("failed to read nonce:  %s", err)
	}
	r.Nonce = string(buf)

	// SIGNATURE
	buf = make([]byte, socksz.LengthSignature)
	n, err = io.ReadFull(reader, buf)
	if n != socksz.LengthSignature || err != nil {
		return fmt.Errorf("failed to read signature:  %s", err)
	}
	r.Signature = string(buf)

	return nil
}
