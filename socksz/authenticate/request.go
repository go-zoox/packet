package authenticate

import (
	"bytes"
	"fmt"
	"io"

	"github.com/go-zoox/crypto/hmac"
	"github.com/go-zoox/packet/socksz"
	"github.com/go-zoox/random"
)

// DATA Protocol:
//
// AUTHENTICATE DATA:
// request:  USER_CLIENT_ID_LENGTH | USER_CLIENT_ID | TIMESTAMP | NONCE | SIGNATURE
//               1                 |    -           |    13     |   6   |  64 HMAC_SHA256

// Request is the request for authenticate
type Request struct {
	UserClientID string
	Timestamp    string
	Nonce        string
	Signature    string
	//
	Secret string
}

// Encode encodes the request
func (r *Request) Encode() ([]byte, error) {
	if r.Secret == "" {
		return nil, fmt.Errorf("secret is required")
	}

	// generates start
	if r.Nonce == "" {
		r.Nonce = random.String(socksz.LengthNonce)
	}
	r.Signature = hmac.Sha256(fmt.Sprintf("%s_%s_%s", r.UserClientID, r.Timestamp, r.Nonce), r.Secret, "hex")
	// generates done

	buf := bytes.NewBuffer([]byte{})

	lengthUserClientID := len(r.UserClientID)
	err := buf.WriteByte(byte(lengthUserClientID))
	if err != nil {
		return nil, fmt.Errorf("failed to write user client id length:  %s", err)
	}

	n, err := buf.WriteString(r.UserClientID)
	if n != lengthUserClientID || err != nil {
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

	// USER_CLIENT_ID_LENGTH
	buf := make([]byte, socksz.LengthUserClientIDLength)
	n, err := io.ReadFull(reader, buf)
	if n != socksz.LengthUserClientIDLength || err != nil {
		return fmt.Errorf("failed to read user client id length:  %s", err)
	}
	lengthUserClientID := int(buf[0])

	// USER_CLIENT_ID
	buf = make([]byte, lengthUserClientID)
	n, err = io.ReadFull(reader, buf)
	if n != lengthUserClientID || err != nil {
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

func (r *Request) Verify() error {
	if r.Secret == "" {
		return fmt.Errorf("secret is required")
	}

	// verify signature
	if r.Signature != hmac.Sha256(fmt.Sprintf("%s_%s_%s", r.UserClientID, r.Timestamp, r.Nonce), r.Secret, "hex") {
		return fmt.Errorf("invalid signature")
	}

	return nil
}
