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
	LENGTH_VER         = 1
	LENGTH_CMD         = 1
	LENGTH_CRYPTO      = 1
	LENGTH_COMPRESSION = 1
)

type Base struct {
	VER         uint8
	CMD         uint8
	CRYPTO      uint8
	COMPRESSION uint8
	DATA        []byte
}

func (r *Base) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.WriteByte(r.VER)
	buf.WriteByte(r.CMD)
	buf.WriteByte(r.CRYPTO)
	buf.WriteByte(r.COMPRESSION)
	buf.Write(r.DATA)
	return buf.Bytes(), nil
}

func (r *Base) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// VER
	buf := make([]byte, LENGTH_VER)
	n, err := io.ReadFull(reader, buf)
	if n != LENGTH_VER || err != nil {
		return fmt.Errorf("failed to read VER:  %s", err)
	}
	r.VER = uint8(buf[0])

	// CMD
	buf = make([]byte, LENGTH_CMD)
	n, err = io.ReadFull(reader, buf)
	if n != LENGTH_CMD || err != nil {
		return fmt.Errorf("failed to read CMD:  %s", err)
	}
	r.CMD = uint8(buf[0])

	// CRYPTO
	buf = make([]byte, LENGTH_CRYPTO)
	n, err = io.ReadFull(reader, buf)
	if n != LENGTH_CRYPTO || err != nil {
		return fmt.Errorf("failed to read CRYPTO:  %s", err)
	}
	r.CRYPTO = uint8(buf[0])

	// COMPRESSION
	buf = make([]byte, LENGTH_COMPRESSION)
	n, err = io.ReadFull(reader, buf)
	if n != LENGTH_COMPRESSION || err != nil {
		return fmt.Errorf("failed to read COMPRESSION:  %s", err)
	}
	r.COMPRESSION = uint8(buf[0])

	// DATA
	r.DATA, err = io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("failed to read DATA:  %s", err)
	}

	return nil
}
