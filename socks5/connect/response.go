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
	// LengthMETHOD is the byte length of METHOD
	LengthMETHOD = 1
	// LengthREP is the byte length of REP
	LengthREP = 1

	// LengthBindAddr = Variable

	// LengthBindAddrIPv4 is the byte length of BIND_ADDR_IPv4
	LengthBindAddrIPv4 = 4
	// LengthBindAddrIPv6 is the byte length of BIND_ADDR_IPv6
	LengthBindAddrIPv6 = 6
	// LengthBIND_ADDR_DOMAIN = Variable

	// LengthBindPort is the byte length of BIND_PORT
	LengthBindPort = 2
)

// Response is the response for connect
type Response struct {
	Ver      uint8
	Rep      uint8
	Rsv      uint8
	ATyp     uint8
	BindAddr string
	BindPort int
}

// Encode encodes the response
func (r *Response) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)

	err := buf.WriteByte(r.Ver)
	if err != nil {
		return nil, fmt.Errorf("failed to write Ver: %s", err)
	}

	err = buf.WriteByte(r.Rep)
	if err != nil {
		return nil, fmt.Errorf("failed to write Rep: %s", err)
	}

	err = buf.WriteByte(r.Rsv)
	if err != nil {
		return nil, fmt.Errorf("failed to write Rsv: %s", err)
	}

	err = buf.WriteByte(r.ATyp)
	if err != nil {
		return nil, fmt.Errorf("failed to write ATyp: %s", err)
	}

	switch r.ATyp {
	case ATypIPv4:
		// 4 Bytes
		parts := strings.Split(r.BindAddr, ".")
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
	case ATypIPv6:
		// 16 Bytes
		// buf.Write(r.BINDAddr)
		return nil, fmt.Errorf("unsupported ATypIPv6")
	case ATypDOMAIN:
		// Variable
		LengthBindDomain := len(r.BindAddr)
		buf.WriteByte(byte(LengthBindDomain))
		buf.WriteString(r.BindAddr)
	default:
		return nil, fmt.Errorf("unsupported ATYP: %d", r.ATyp)
	}

	bufPort := make([]byte, 2)
	binary.BigEndian.PutUint16(bufPort, uint16(r.BindPort))
	buf.Write(bufPort)

	return buf.Bytes(), nil
}

// Decode decodes the response
func (r *Response) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// VER
	buf := make([]byte, LengthVER)
	n, err := io.ReadFull(reader, buf)
	if n != LengthVER || err != nil {
		return fmt.Errorf("failed to read ver:  %s", err)
	}
	r.Ver = buf[0]

	// REP
	buf = make([]byte, LengthREP)
	n, err = io.ReadFull(reader, buf)
	if n != LengthREP || err != nil {
		return fmt.Errorf("failed to read REP:  %s", err)
	}
	r.Rep = buf[0]

	// RSV
	buf = make([]byte, LengthRSV)
	n, err = io.ReadFull(reader, buf)
	if n != LengthRSV || err != nil {
		return fmt.Errorf("failed to read RSV:  %s", err)
	}
	r.Rsv = buf[0]

	// ATYP
	buf = make([]byte, LengthATYP)
	n, err = io.ReadFull(reader, buf)
	if n != LengthATYP || err != nil {
		return fmt.Errorf("failed to read ATYP:  %s", err)
	}
	r.ATyp = buf[0]

	// BIND_ADDR
	switch r.ATyp {
	case ATypIPv4:
		// 4 Bytes
		buf = make([]byte, LengthBindAddrIPv4)
		n, err = io.ReadFull(reader, buf)
		if n != LengthBindAddrIPv4 || err != nil {
			return fmt.Errorf("failed to read BIND_ADDR(IPv4):  %s", err)
		}
		r.BindAddr = fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3])
	case ATypIPv6:
		// // 16 Bytes
		// buf = make([]byte, LengthBIND_ADDR_IPv6)
		// n, err = io.ReadFull(reader, buf)
		// if n!= LengthBIND_ADDR_IPv6 || err!= nil {
		// 	return fmt.Errorf("failed to read BIND_ADDR(IPv6):  %s", err)
		// }
		return fmt.Errorf("unsupported ATypIPv6")
	case ATypDOMAIN:
		buf = make([]byte, 1)
		n, err = io.ReadFull(reader, buf)
		if n != 1 || err != nil {
			return fmt.Errorf("failed to read BIND_ADDR(IPv4):  %s", err)
		}
		LengthBindAddrDomain := int(buf[0])

		buf = make([]byte, LengthBindAddrDomain)
		n, err = io.ReadFull(reader, buf)
		if n != LengthBindAddrDomain || err != nil {
			return fmt.Errorf("failed to read BIND_ADDR(IPv4):  %s", err)
		}
		r.BindAddr = string(buf)
	default:
		return fmt.Errorf("unsupported ATYP: %d", r.ATyp)
	}

	// BIND_PORT
	buf = make([]byte, LengthBindPort)
	n, err = io.ReadFull(reader, buf)
	if n != LengthBindPort || err != nil {
		return fmt.Errorf("failed to read BIND_PORT:  %s", err)
	}
	r.BindPort = int(binary.BigEndian.Uint16(buf))

	return nil
}
