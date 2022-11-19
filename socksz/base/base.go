package base

import (
	"bytes"
	"fmt"
	"io"

	"github.com/go-zoox/crypto/aes"
	"github.com/go-zoox/packet/socksz"
)

// DATA Protocol:
//
// CONNECTION CLOSE DATA:
//  VER | CMD | CRYPTO | COMPRESS | DATA
//   1  |  1  |  1     |   1      | -

// Base ...
type Base struct {
	Ver         uint8
	Cmd         uint8
	Crypto      uint8
	Compression uint8
	Data        []byte
	//
	Secret string
}

// Encode encodes the base data
func (r *Base) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)

	err := buf.WriteByte(r.Ver)
	if err != nil {
		return nil, fmt.Errorf("failed to write Ver: %s", err)
	}

	err = buf.WriteByte(r.Cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to write Cmd: %s", err)
	}

	err = buf.WriteByte(r.Crypto)
	if err != nil {
		return nil, fmt.Errorf("failed to write Crypto: %s", err)
	}

	err = buf.WriteByte(r.Compression)
	if err != nil {
		return nil, fmt.Errorf("failed to write Compression: %s", err)
	}

	cipherData, err := r.encrypt(r.Crypto, r.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt Data: %s", err)
	}

	_, err = buf.Write(cipherData)
	if err != nil {
		return nil, fmt.Errorf("failed to write Data: %s", err)
	}

	return buf.Bytes(), nil
}

// Decode decodes the base data
func (r *Base) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// VER
	buf := make([]byte, socksz.LengthVer)
	n, err := io.ReadFull(reader, buf)
	if n != socksz.LengthVer || err != nil {
		return fmt.Errorf("failed to read VER:  %s", err)
	}
	r.Ver = uint8(buf[0])

	// CMD
	buf = make([]byte, socksz.LengthCmd)
	n, err = io.ReadFull(reader, buf)
	if n != socksz.LengthCmd || err != nil {
		return fmt.Errorf("failed to read CMD:  %s", err)
	}
	r.Cmd = uint8(buf[0])

	// CRYPTO
	buf = make([]byte, socksz.LengthCrypto)
	n, err = io.ReadFull(reader, buf)
	if n != socksz.LengthCrypto || err != nil {
		return fmt.Errorf("failed to read CRYPTO:  %s", err)
	}
	r.Crypto = uint8(buf[0])

	// COMPRESSION
	buf = make([]byte, socksz.LengthCompression)
	n, err = io.ReadFull(reader, buf)
	if n != socksz.LengthCompression || err != nil {
		return fmt.Errorf("failed to read COMPRESSION:  %s", err)
	}
	r.Compression = uint8(buf[0])

	// DATA
	cipherData, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("failed to read DATA:  %s", err)
	}

	r.Data, err = r.descript(r.Crypto, cipherData)
	if err != nil {
		return fmt.Errorf("failed to descript DATA:  %s", err)
	}

	return nil
}

func (r *Base) encrypt(crypto uint8, plain []byte) ([]byte, error) {
	switch crypto {
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
		return nil, fmt.Errorf("unsupported encryption algorithm: %d", crypto)
	}
}

func (r *Base) descript(crypto uint8, cipher []byte) ([]byte, error) {
	switch crypto {
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
		return nil, fmt.Errorf("unsupported description algorithm: %d", crypto)
	}
}
