package joinAsAgent

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
// Request DATA:
//  Type | ID_LENGTH | ID | TIMESTAMP | NONCE | SIGNATURE
// 	  1  |       1   | -  |    13     |   6   |  64 (HMAC_SHA256)

// Request is the request for authenticate
type Request struct {
	Type      string
	ID        string
	Timestamp string
	Nonce     string
	Signature string
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
	r.Signature = hmac.Sha256(fmt.Sprintf("%s_%s_%s", r.ID, r.Timestamp, r.Nonce), r.Secret, "hex")
	// generates done

	buf := bytes.NewBuffer([]byte{})

	switch r.Type {
	case "user":
		err := buf.WriteByte(0x00)
		if err != nil {
			return nil, fmt.Errorf("failed to write type user: %s", err)
		}
	case "room":
		err := buf.WriteByte(0x01)
		if err != nil {
			return nil, fmt.Errorf("failed to write type room: %s", err)
		}
	default:
		return nil, fmt.Errorf("unsupport type: %s", r.Type)
	}

	lengthUserClientID := len(r.ID)
	err := buf.WriteByte(byte(lengthUserClientID))
	if err != nil {
		return nil, fmt.Errorf("failed to write user client id or room id length:  %s", err)
	}

	n, err := buf.WriteString(r.ID)
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

	// Type
	buf := make([]byte, socksz.LengthJoinAsRoomType)
	n, err := io.ReadFull(reader, buf)
	if n != socksz.LengthJoinAsRoomType || err != nil {
		return fmt.Errorf("failed to read type length:  %s", err)
	}
	typValue := int(buf[0])
	switch typValue {
	case 0x00:
		r.Type = "user"
	case 0x01:
		r.Type = "room"
	default:
		return fmt.Errorf("unsupport type: %d", typValue)
	}

	// ID_LENGTH
	buf = make([]byte, socksz.LengthUserClientIDLength)
	n, err = io.ReadFull(reader, buf)
	if n != socksz.LengthUserClientIDLength || err != nil {
		return fmt.Errorf("failed to read user client id length:  %s", err)
	}
	lengthUserClientID := int(buf[0])

	// ID
	buf = make([]byte, lengthUserClientID)
	n, err = io.ReadFull(reader, buf)
	if n != lengthUserClientID || err != nil {
		return fmt.Errorf("failed to read user client id:  %s", err)
	}
	r.ID = string(buf)

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
	if r.Signature != hmac.Sha256(fmt.Sprintf("%s_%s_%s", r.ID, r.Timestamp, r.Nonce), r.Secret, "hex") {
		return fmt.Errorf("invalid signature")
	}

	return nil
}
