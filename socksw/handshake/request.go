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
	LENGTH_CONNECTION_ID              = 21
	LENGTH_TARGET_USER_CLIENT_ID      = 10
	LENGTH_TARGET_USER_PAIR_SIGNATURE = 64
	LENGTH_NETWORK                    = 1
	LENGTH_ATYP                       = 1
	// LENGTH_DST_ADDR = 4
	LENGTH_DST_PORT = 2

	ATYP_IPv4   = 0x01
	ATYP_IPv6   = 0x04
	ATYP_DOMAIN = 0x03

	NETWORK_TCP = 0x01
	NETWORK_UDP = 0x02
)

type Request struct {
	CONNECTION_ID              string
	TARGET_USER_CLIENT_ID      string
	TARGET_USER_PAIR_SIGNATURE string
	NETWORK                    uint8
	ATYP                       uint8
	DST_ADDR                   string
	DST_PORT                   uint16
}

func (r *Request) Encode() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(r.CONNECTION_ID)
	buf.WriteString(r.TARGET_USER_CLIENT_ID)
	buf.WriteString(r.TARGET_USER_PAIR_SIGNATURE)
	buf.WriteByte(r.NETWORK)
	buf.WriteByte(r.ATYP)

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

	DSTAddrLength := len(r.DST_ADDR)
	buf.WriteByte(byte(DSTAddrLength))
	buf.WriteString(r.DST_ADDR)
	portBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(portBytes, r.DST_PORT)
	buf.Write(portBytes)

	return buf.Bytes(), nil
}

func (r *Request) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// CONNECTION_ID
	buf := make([]byte, LENGTH_CONNECTION_ID)
	n, err := io.ReadFull(reader, buf)
	if n != LENGTH_CONNECTION_ID || err != nil {
		return fmt.Errorf("failed to read connection id:  %s", err)
	}
	r.CONNECTION_ID = string(buf)

	// TARGET_USER_CLIENT_ID
	buf = make([]byte, LENGTH_TARGET_USER_CLIENT_ID)
	n, err = io.ReadFull(reader, buf)
	if n != LENGTH_TARGET_USER_CLIENT_ID || err != nil {
		return fmt.Errorf("failed to read target user client id:  %s", err)
	}
	r.TARGET_USER_CLIENT_ID = string(buf)

	// TARGET_USER_PAIR_KEY
	buf = make([]byte, LENGTH_TARGET_USER_PAIR_SIGNATURE)
	n, err = io.ReadFull(reader, buf)
	if n != LENGTH_TARGET_USER_PAIR_SIGNATURE || err != nil {
		return fmt.Errorf("failed to read target user pair key:  %s", err)
	}
	r.TARGET_USER_PAIR_SIGNATURE = string(buf)

	// NETWORK
	buf = make([]byte, LENGTH_NETWORK)
	n, err = io.ReadFull(reader, buf)
	if n != LENGTH_NETWORK || err != nil {
		return fmt.Errorf("failed to read signature:  %s", err)
	}
	r.NETWORK = uint8(buf[0])

	// ATYP
	buf = make([]byte, LENGTH_ATYP)
	n, err = io.ReadFull(reader, buf)
	if n != LENGTH_ATYP || err != nil {
		return fmt.Errorf("failed to read atyp:  %s", err)
	}
	r.ATYP = uint8(buf[0])

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
	r.DST_ADDR = string(buf)

	// DSTAddrPort
	buf = make([]byte, LENGTH_DST_PORT)
	n, err = io.ReadFull(reader, buf)
	if n != LENGTH_DST_PORT || err != nil {
		return fmt.Errorf("failed to read atyp:  %s", err)
	}
	r.DST_PORT = binary.BigEndian.Uint16(buf[:2])

	return nil
}
