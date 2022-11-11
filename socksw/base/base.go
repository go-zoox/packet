package close

import (
	"bytes"
	"fmt"
	"io"
)

// DATA Protocol:
//
// CONNECTION CLOSE DATA:
//  VER | CMD | CRYPTO | COMPRESS | DATA
//   1  |  1  |  1     |   1      | -

const (
	// LengthVer ...
	LengthVer = 1
	// LengthCmd ...
	LengthCmd = 1
	// LengthCrypto ...
	LengthCrypto = 1
	// LengthCompression ...
	LengthCompression = 1
)

// Base ...
type Base struct {
	Ver         uint8
	Cmd         uint8
	Crypto      uint8
	Compression uint8
	Data        []byte
}

// Encode encodes the base data
func (r *Base) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.WriteByte(r.Ver)
	buf.WriteByte(r.Cmd)
	buf.WriteByte(r.Crypto)
	buf.WriteByte(r.Compression)
	buf.Write(r.Data)
	return buf.Bytes(), nil
}

// Decode decodes the base data
func (r *Base) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// VER
	buf := make([]byte, LengthVer)
	n, err := io.ReadFull(reader, buf)
	if n != LengthVer || err != nil {
		return fmt.Errorf("failed to read VER:  %s", err)
	}
	r.Ver = uint8(buf[0])

	// CMD
	buf = make([]byte, LengthCmd)
	n, err = io.ReadFull(reader, buf)
	if n != LengthCmd || err != nil {
		return fmt.Errorf("failed to read CMD:  %s", err)
	}
	r.Cmd = uint8(buf[0])

	// CRYPTO
	buf = make([]byte, LengthCrypto)
	n, err = io.ReadFull(reader, buf)
	if n != LengthCrypto || err != nil {
		return fmt.Errorf("failed to read CRYPTO:  %s", err)
	}
	r.Crypto = uint8(buf[0])

	// COMPRESSION
	buf = make([]byte, LengthCompression)
	n, err = io.ReadFull(reader, buf)
	if n != LengthCompression || err != nil {
		return fmt.Errorf("failed to read COMPRESSION:  %s", err)
	}
	r.Compression = uint8(buf[0])

	// DATA
	r.Data, err = io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("failed to read DATA:  %s", err)
	}

	return nil
}
