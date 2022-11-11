package connect

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	LENGTH_METHOD = 1
	LENGTH_REP    = 1

	// LENGTH_BIND_ADDR = Variable
	LENGTH_BIND_ADDR_IPv4 = 4
	LENGTH_BIND_ADDR_IPv6 = 6
	// LENGTH_BIND_ADDR_DOMAIN = Variable

	LENGTH_BIND_PORT = 2
)

type Response struct {
	VER       uint8
	REP       uint8
	RSV       uint8
	ATYP      uint8
	BIND_ADDR string
	BIND_PORT int
}

func (r *Response) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.WriteByte(r.VER)
	buf.WriteByte(r.REP)
	buf.WriteByte(r.RSV)
	buf.WriteByte(r.ATYP)

	switch r.ATYP {
	case ATYP_IPv4:
		// 4 Bytes
		parts := strings.Split(r.BIND_ADDR, ".")
		parts0, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		parts1, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		parts2, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, err
		}
		parts3, err := strconv.Atoi(parts[3])
		if err != nil {
			return nil, err
		}

		buf.Write([]byte{byte(parts0), byte(parts1), byte(parts2), byte(parts3)})
	case ATYP_IPv6:
		// 16 Bytes
		// buf.Write(r.BINDAddr)
		return nil, fmt.Errorf("unsupported ATYP_IPv6")
	case ATYP_DOMAIN:
		// Variable
		LENGTH_BIND_DOMAIN := len(r.BIND_ADDR)
		buf.WriteByte(byte(LENGTH_BIND_DOMAIN))
		buf.WriteString(r.BIND_ADDR)
	default:
		return nil, fmt.Errorf("unsupported ATYP: %d", r.ATYP)
	}

	bufPort := make([]byte, 2)
	binary.BigEndian.PutUint16(bufPort, uint16(r.BIND_PORT))
	buf.Write(bufPort)

	return buf.Bytes(), nil
}

func (r *Response) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// VER
	buf := make([]byte, LENGTH_VER)
	n, err := io.ReadFull(reader, buf)
	if n != LENGTH_VER || err != nil {
		return fmt.Errorf("failed to read ver:  %s", err)
	}
	r.VER = buf[0]

	// REP
	buf = make([]byte, LENGTH_REP)
	n, err = io.ReadFull(reader, buf)
	if n != LENGTH_REP || err != nil {
		return fmt.Errorf("failed to read REP:  %s", err)
	}
	r.REP = buf[0]

	// RSV
	buf = make([]byte, LENGTH_RSV)
	n, err = io.ReadFull(reader, buf)
	if n != LENGTH_RSV || err != nil {
		return fmt.Errorf("failed to read RSV:  %s", err)
	}
	r.RSV = buf[0]

	// ATYP
	buf = make([]byte, LENGTH_ATYP)
	n, err = io.ReadFull(reader, buf)
	if n != LENGTH_ATYP || err != nil {
		return fmt.Errorf("failed to read ATYP:  %s", err)
	}
	r.ATYP = buf[0]

	// BIND_ADDR
	switch r.ATYP {
	case ATYP_IPv4:
		// 4 Bytes
		buf = make([]byte, LENGTH_BIND_ADDR_IPv4)
		n, err = io.ReadFull(reader, buf)
		if n != LENGTH_BIND_ADDR_IPv4 || err != nil {
			return fmt.Errorf("failed to read BIND_ADDR(IPv4):  %s", err)
		}
		r.BIND_ADDR = fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3])
	case ATYP_IPv6:
		// // 16 Bytes
		// buf = make([]byte, LENGTH_BIND_ADDR_IPv6)
		// n, err = io.ReadFull(reader, buf)
		// if n!= LENGTH_BIND_ADDR_IPv6 || err!= nil {
		// 	return fmt.Errorf("failed to read BIND_ADDR(IPv6):  %s", err)
		// }
		return fmt.Errorf("unsupported ATYP_IPv6")
	case ATYP_DOMAIN:
		buf = make([]byte, 1)
		n, err = io.ReadFull(reader, buf)
		if n != 1 || err != nil {
			return fmt.Errorf("failed to read BIND_ADDR(IPv4):  %s", err)
		}
		LENGTH_BIND_ADDR_DOMAIN := int(buf[0])

		buf = make([]byte, LENGTH_BIND_ADDR_DOMAIN)
		n, err = io.ReadFull(reader, buf)
		if n != LENGTH_BIND_ADDR_DOMAIN || err != nil {
			return fmt.Errorf("failed to read BIND_ADDR(IPv4):  %s", err)
		}
		r.BIND_ADDR = string(buf)
	default:
		return fmt.Errorf("unsupported ATYP: %d", r.ATYP)
	}

	// BIND_PORT
	buf = make([]byte, LENGTH_BIND_PORT)
	n, err = io.ReadFull(reader, buf)
	if n != LENGTH_BIND_PORT || err != nil {
		return fmt.Errorf("failed to read BIND_PORT:  %s", err)
	}
	r.BIND_PORT = int(binary.BigEndian.Uint16(buf))

	return nil
}
