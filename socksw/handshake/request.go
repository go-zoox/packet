package handshake

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// DATA Protocol:
//
// Handshake DATA:
// request:  CONNECTION_ID | TARGET_USER_CLIENT_ID | TARGET_USER_PAIR_SIGNATURE |  NETWORK   | ATYP                 | DST.ADDR 							 | DST.PORT
//					       21      |       10              |					64                | 1(tcp/udp) | 1(IPv4/IPv6/Domain)  |   4 or 16 or domain    |    2
// response: STATUS | MESSAGE
//            1     |  -

const (
	// LengthConnectionID ...
	LengthConnectionID = 21
	// LengthTargetUserClientID ...
	LengthTargetUserClientID = 10
	// LengthTargetUserPairSignature ...
	LengthTargetUserPairSignature = 64
	// LengthNetwork ...
	LengthNetwork = 1
	// LengthATyp ...
	LengthATyp = 1

	// LengthDSTAddr = 4

	// LengthDSTPort ...
	LengthDSTPort = 2

	// ATypIPv4 ...
	ATypIPv4 = 0x01
	// ATypIPv6 ...
	ATypIPv6 = 0x04
	// ATypDomain ...
	ATypDomain = 0x03

	// NetworkTCP ...
	NetworkTCP = 0x01
	// NetworkUDP ...
	NetworkUDP = 0x02
)

// Request is the request for handshake
type Request struct {
	ConnectionID            string
	TargetUserClientID      string
	TargetUserPairSignature string
	Network                 uint8
	ATyp                    uint8
	DSTAddr                 string
	DSTPort                 uint16
}

// Encode encodes the data
func (r *Request) Encode() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	n, err := buf.WriteString(r.ConnectionID)
	if n != LengthConnectionID || err != nil {
		return nil, fmt.Errorf("failed to write ConnectionID: %s", err)
	}

	n, err = buf.WriteString(r.TargetUserClientID)
	if n != LengthTargetUserClientID || err != nil {
		return nil, fmt.Errorf("failed to write TargetUserClientID: %s", err)
	}

	n, err = buf.WriteString(r.TargetUserPairSignature)
	if n != LengthTargetUserPairSignature || err != nil {
		return nil, fmt.Errorf("failed to write TargetUserPairSignature: %s", err)
	}

	err = buf.WriteByte(r.Network)
	if err != nil {
		return nil, fmt.Errorf("failed to write Network: %s", err)
	}

	err = buf.WriteByte(r.ATyp)
	if err != nil {
		return nil, fmt.Errorf("failed to write ATyp: %s", err)
	}

	// switch a.ATyp {
	// case ATYP_IPv4:
	// 	// 1.1.1.1
	// 	for _, p := range strings.Split(a.DSTAddr, ".") {
	// 		if v, err := strconv.Atoi(p); err != nil {
	// 			return fmt.Errorf("invalid atyp IPv4 dst addr(%s)")
	// 		} else {
	// 			buf.WriteByte(byte(v))
	// 		}
	// 	}
	// }

	LengthDSTAddr := len(r.DSTAddr)
	buf.WriteByte(byte(LengthDSTAddr))
	n, err = buf.WriteString(r.DSTAddr)
	if n != LengthDSTAddr || err != nil {
		return nil, fmt.Errorf("failed to write DSTAddr: %s", err)
	}

	portBytes := make([]byte, LengthDSTPort)
	binary.BigEndian.PutUint16(portBytes, r.DSTPort)
	n, err = buf.Write(portBytes)
	if n != LengthDSTPort || err != nil {
		return nil, fmt.Errorf("failed to write DSTPort: %s", err)
	}

	return buf.Bytes(), nil
}

// Decode decodes the data
func (r *Request) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// CONNECTION_ID
	buf := make([]byte, LengthConnectionID)
	n, err := io.ReadFull(reader, buf)
	if n != LengthConnectionID || err != nil {
		return fmt.Errorf("failed to read connection id:  %s", err)
	}
	r.ConnectionID = string(buf)

	// TARGET_USER_CLIENT_ID
	buf = make([]byte, LengthTargetUserClientID)
	n, err = io.ReadFull(reader, buf)
	if n != LengthTargetUserClientID || err != nil {
		return fmt.Errorf("failed to read target user client id:  %s", err)
	}
	r.TargetUserClientID = string(buf)

	// TARGET_USER_PAIR_KEY
	buf = make([]byte, LengthTargetUserPairSignature)
	n, err = io.ReadFull(reader, buf)
	if n != LengthTargetUserPairSignature || err != nil {
		return fmt.Errorf("failed to read target user pair key:  %s", err)
	}
	r.TargetUserPairSignature = string(buf)

	// NETWORK
	buf = make([]byte, LengthNetwork)
	n, err = io.ReadFull(reader, buf)
	if n != LengthNetwork || err != nil {
		return fmt.Errorf("failed to read signature:  %s", err)
	}
	r.Network = uint8(buf[0])

	// ATYP
	buf = make([]byte, LengthATyp)
	n, err = io.ReadFull(reader, buf)
	if n != LengthATyp || err != nil {
		return fmt.Errorf("failed to read atyp:  %s", err)
	}
	r.ATyp = uint8(buf[0])

	// DSTAddr
	buf = make([]byte, 1)
	n, err = io.ReadFull(reader, buf)
	if n != 1 || err != nil {
		return fmt.Errorf("failed to read dst addr length:  %s", err)
	}
	dstAddrLength := int(buf[0])
	buf = make([]byte, dstAddrLength)
	n, err = io.ReadFull(reader, buf)
	if n != dstAddrLength || err != nil {
		return fmt.Errorf("failed to read dst addr:  %s", err)
	}
	r.DSTAddr = string(buf)

	// DSTAddrPort
	buf = make([]byte, LengthDSTPort)
	n, err = io.ReadFull(reader, buf)
	if n != LengthDSTPort || err != nil {
		return fmt.Errorf("failed to read atyp:  %s", err)
	}
	r.DSTPort = binary.BigEndian.Uint16(buf[:2])

	return nil
}
