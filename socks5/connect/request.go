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
	// LengthVER is the byte length of VER
	LengthVER = 1
	// LengthCMD is the byte length of CMD
	LengthCMD = 1
	// LengthRSV is the byte length of RSV
	LengthRSV = 1
	// LengthATYP is the byte length of ATYP
	LengthATYP = 1

	// LengthDST_ADDR = Variable

	// LengthDSTAddrIPv4 is the byte length of DST_ADDR_IPv4
	LengthDSTAddrIPv4 = 4
	// LengthDSTAddrIPv6 is the byte length of DST_ADDR_IPv6
	LengthDSTAddrIPv6 = 6
	// LengthDST_ADDR_DOMAIN = Variable

	// LengthDSTPort is the byte length of DST_PORT
	LengthDSTPort = 2
)

const (
	// ATypIPv4 means IPv4 address
	ATypIPv4 = 0x01
	// ATypDOMAIN means domain name
	ATypDOMAIN = 0x03
	// ATypIPv6 means IPv6 address
	ATypIPv6 = 0x04
)

// Request is the request for connect
type Request struct {
	VER     uint8
	CMD     uint8
	RSV     uint8
	ATYP    uint8
	DSTADDR string
	DSTPORT int
}

// Encode encodes the request
func (r *Request) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.WriteByte(r.VER)
	buf.WriteByte(r.CMD)
	buf.WriteByte(r.RSV)
	buf.WriteByte(r.ATYP)

	switch r.ATYP {
	case ATypIPv4:
		// 4 Bytes
		parts := strings.Split(r.DSTADDR, ".")
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
		// buf.Write(r.DSTAddr)
		return nil, fmt.Errorf("unsupported ATypIPv6")
	case ATypDOMAIN:
		// Variable
		LengthDSTDomain := len(r.DSTADDR)
		buf.WriteByte(byte(LengthDSTDomain))
		buf.WriteString(r.DSTADDR)
	default:
		return nil, fmt.Errorf("unsupported ATYP: %d", r.ATYP)
	}

	bufPort := make([]byte, 2)
	binary.BigEndian.PutUint16(bufPort, uint16(r.DSTPORT))
	buf.Write(bufPort)

	return buf.Bytes(), nil
}

// Decode decodes the request
func (r *Request) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// VER
	buf := make([]byte, LengthVER)
	n, err := io.ReadFull(reader, buf)
	if n != LengthVER || err != nil {
		return fmt.Errorf("failed to read ver:  %s", err)
	}
	r.VER = buf[0]

	// CMD
	buf = make([]byte, LengthCMD)
	n, err = io.ReadFull(reader, buf)
	if n != LengthCMD || err != nil {
		return fmt.Errorf("failed to read CMD:  %s", err)
	}
	r.CMD = buf[0]

	// RSV
	buf = make([]byte, LengthRSV)
	n, err = io.ReadFull(reader, buf)
	if n != LengthRSV || err != nil {
		return fmt.Errorf("failed to read RSV:  %s", err)
	}
	r.RSV = buf[0]

	// ATYP
	buf = make([]byte, LengthATYP)
	n, err = io.ReadFull(reader, buf)
	if n != LengthATYP || err != nil {
		return fmt.Errorf("failed to read ATYP:  %s", err)
	}
	r.ATYP = buf[0]

	// DST_ADDR
	switch r.ATYP {
	case ATypIPv4:
		// 4 Bytes
		buf = make([]byte, LengthDSTAddrIPv4)
		n, err = io.ReadFull(reader, buf)
		if n != LengthDSTAddrIPv4 || err != nil {
			return fmt.Errorf("failed to read DST_ADDR(IPv4):  %s", err)
		}
		r.DSTADDR = fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3])
	case ATypIPv6:
		// // 16 Bytes
		// buf = make([]byte, LengthDST_ADDR_IPv6)
		// n, err = io.ReadFull(reader, buf)
		// if n!= LengthDST_ADDR_IPv6 || err!= nil {
		// 	return fmt.Errorf("failed to read DST_ADDR(IPv6):  %s", err)
		// }
		return fmt.Errorf("unsupported ATypIPv6")
	case ATypDOMAIN:
		buf = make([]byte, 1)
		n, err = io.ReadFull(reader, buf)
		if n != 1 || err != nil {
			return fmt.Errorf("failed to read DST_ADDR(IPv4):  %s", err)
		}
		LengthDSTAddrDomain := int(buf[0])

		buf = make([]byte, LengthDSTAddrDomain)
		n, err = io.ReadFull(reader, buf)
		if n != LengthDSTAddrDomain || err != nil {
			return fmt.Errorf("failed to read DST_ADDR(IPv4):  %s", err)
		}
		r.DSTADDR = string(buf)
	default:
		return fmt.Errorf("unsupported ATYP: %d", r.ATYP)
	}

	// DST_PORT
	buf = make([]byte, LengthDSTPort)
	n, err = io.ReadFull(reader, buf)
	if n != LengthDSTPort || err != nil {
		return fmt.Errorf("failed to read DST_PORT:  %s", err)
	}
	r.DSTPORT = int(binary.BigEndian.Uint16(buf))

	return nil
}
