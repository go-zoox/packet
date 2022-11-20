package forward

import (
	"bytes"
	"fmt"
	"io"

	"github.com/go-zoox/crypto/aes"
	"github.com/go-zoox/packet/socksz"
)

// DATA Protocol:
//
// TRANSMISIION DATA:
// request:  CONNECTION_ID | DATA
//					       13      |  -

// Forward ...
type Forward struct {
	Secret string
	Crypto uint8
	//
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

	cipher, err := r.encrypt(r.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt Data: %s", err)
	}
	_, err = buf.Write(cipher)
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

	cipher, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("failed to read data:  %s", err)
	}

	r.Data, err = r.descript(cipher)
	if err != nil {
		return fmt.Errorf("failed to descript DATA:  %s", err)
	}

	return nil
}

func (r *Forward) encrypt(plain []byte) ([]byte, error) {
	switch r.Crypto {
	case 0x00:
		return plain, nil
	case 0x01: // aes-128-cfb
		enc, _ := aes.NewCFB(128, &aes.Base64Encoding{}, nil)
		return enc.Encrypt(plain, aes.PaddingKey([]byte(r.Secret), 16))
	case 0x02: // aes-192-cfb
		enc, _ := aes.NewCFB(192, &aes.Base64Encoding{}, nil)
		return enc.Encrypt(plain, aes.PaddingKey([]byte(r.Secret), 24))
	case 0x03: // aes-256-cfb
		enc, _ := aes.NewCFB(256, &aes.Base64Encoding{}, nil)
		return enc.Encrypt(plain, aes.PaddingKey([]byte(r.Secret), 32))
	case 0x04: // aes-128-cbc
		enc, _ := aes.NewCBC(128, nil, &aes.Base64Encoding{}, nil)
		return enc.Encrypt(plain, aes.PaddingKey([]byte(r.Secret), 16))
	case 0x05: // aes-192-cbc
		enc, _ := aes.NewCBC(192, nil, &aes.Base64Encoding{}, nil)
		return enc.Encrypt(plain, aes.PaddingKey([]byte(r.Secret), 24))
	case 0x06: // aes-256-cbc
		enc, _ := aes.NewCBC(256, nil, &aes.Base64Encoding{}, nil)
		return enc.Encrypt(plain, aes.PaddingKey([]byte(r.Secret), 32))
	case 0x07: // aes-128-gcm
		enc, _ := aes.NewGCM(128, &aes.Base64Encoding{}, nil)
		return enc.Encrypt(plain, aes.PaddingKey([]byte(r.Secret), 16))
	case 0x08: // aes-192-gcm
		enc, _ := aes.NewGCM(192, &aes.Base64Encoding{}, nil)
		return enc.Encrypt(plain, aes.PaddingKey([]byte(r.Secret), 24))
	case 0x09: // aes-256-gcm
		enc, _ := aes.NewGCM(256, &aes.Base64Encoding{}, nil)
		return enc.Encrypt(plain, aes.PaddingKey([]byte(r.Secret), 32))
	case 0x10: // chacha20-poly1305
		enc, _ := aes.NewChacha20Poly1305(&aes.Base64Encoding{}, nil)
		return enc.Encrypt(plain, aes.PaddingKey([]byte(r.Secret), 32))
	default:
		return nil, fmt.Errorf("unsupported encryption algorithm: %d", r.Crypto)
	}
}

func (r *Forward) descript(cipher []byte) ([]byte, error) {
	switch r.Crypto {
	case 0x00:
		return cipher, nil
	case 0x01: // aes-128-cfb
		enc, _ := aes.NewCFB(128, &aes.Base64Encoding{}, nil)
		return enc.Decrypt(cipher, aes.PaddingKey([]byte(r.Secret), 16))
	case 0x02: // aes-192-cfb
		enc, _ := aes.NewCFB(192, &aes.Base64Encoding{}, nil)
		return enc.Decrypt(cipher, aes.PaddingKey([]byte(r.Secret), 24))
	case 0x03: // aes-256-cfb
		enc, _ := aes.NewCFB(256, &aes.Base64Encoding{}, nil)
		return enc.Decrypt(cipher, aes.PaddingKey([]byte(r.Secret), 32))
	case 0x04: // aes-128-cfc
		enc, _ := aes.NewCBC(128, nil, &aes.Base64Encoding{}, nil)
		return enc.Decrypt(cipher, aes.PaddingKey([]byte(r.Secret), 16))
	case 0x05: // aes-192-cfc
		enc, _ := aes.NewCBC(192, nil, &aes.Base64Encoding{}, nil)
		return enc.Decrypt(cipher, aes.PaddingKey([]byte(r.Secret), 24))
	case 0x06: // aes-256-cfc
		enc, _ := aes.NewCBC(256, nil, &aes.Base64Encoding{}, nil)
		return enc.Decrypt(cipher, aes.PaddingKey([]byte(r.Secret), 32))
	case 0x07: // aes-128-gcm
		enc, _ := aes.NewGCM(128, &aes.Base64Encoding{}, nil)
		return enc.Decrypt(cipher, aes.PaddingKey([]byte(r.Secret), 16))
	case 0x08: // aes-192-gcm
		enc, _ := aes.NewGCM(192, &aes.Base64Encoding{}, nil)
		return enc.Decrypt(cipher, aes.PaddingKey([]byte(r.Secret), 24))
	case 0x09: // aes-256-gcm
		enc, _ := aes.NewGCM(256, &aes.Base64Encoding{}, nil)
		return enc.Decrypt(cipher, aes.PaddingKey([]byte(r.Secret), 32))
	case 0x10: // chacha20-poly1305
		enc, _ := aes.NewChacha20Poly1305(&aes.Base64Encoding{}, nil)
		return enc.Decrypt(cipher, aes.PaddingKey([]byte(r.Secret), 32))
	default:
		return nil, fmt.Errorf("unsupported description algorithm: %d", r.Crypto)
	}
}
