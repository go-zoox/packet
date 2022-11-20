package handshake

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/go-zoox/crypto/hmac"
	"github.com/go-zoox/packet/socksz"
)

// DATA Protocol:
//
// Handshake DATA:
// request:  CONNECTION_ID | TARGET_USER_CLIENT_ID | TARGET_USER_PAIR_SIGNATURE |  NETWORK   | ATYP                 | DST.ADDR 							 | DST.PORT
//					       13      |       10              |					64                | 1(tcp/udp) | 1(IPv4/IPv6/Domain)  |   4 or 16 or domain    |    2
// response: STATUS | MESSAGE
//            1     |  -

const (
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
	//
	Secret string
}

// Encode encodes the data
func (r *Request) Encode() ([]byte, error) {
	if r.Secret == "" {
		return nil, fmt.Errorf("secret is required")
	}

	// generates start
	r.TargetUserPairSignature = hmac.Sha256(fmt.Sprintf("%s_%s", r.ConnectionID, r.TargetUserClientID), r.Secret)
	// generates done

	buf := bytes.NewBuffer([]byte{})

	n, err := buf.WriteString(r.ConnectionID)
	if n != socksz.LengthConnectionID {
		return nil, fmt.Errorf("failed to write ConnectionID: length expect %d, but got %d", socksz.LengthConnectionID, n)
	} else if err != nil {
		return nil, fmt.Errorf("failed to write ConnectionID: %s", err)
	}

	n, err = buf.WriteString(r.TargetUserClientID)
	if n != socksz.LengthTargetUserClientID {
		return nil, fmt.Errorf("failed to write TargetUserClientID: length expect %d, but got %d", socksz.LengthTargetUserClientID, n)
	} else if err != nil {
		return nil, fmt.Errorf("failed to write TargetUserClientID: %s", err)
	}

	n, err = buf.WriteString(r.TargetUserPairSignature)
	if n != socksz.LengthTargetUserPairSignature {
		return nil, fmt.Errorf("failed to write TargetUserPairSignature: length expect %d, but got %d", socksz.LengthTargetUserPairSignature, n)
	} else if err != nil {
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

	portBytes := make([]byte, socksz.LengthDSTPort)
	binary.BigEndian.PutUint16(portBytes, r.DSTPort)
	n, err = buf.Write(portBytes)
	if n != socksz.LengthDSTPort || err != nil {
		return nil, fmt.Errorf("failed to write DSTPort: %s", err)
	}

	return buf.Bytes(), nil
}

// Decode decodes the data
func (r *Request) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// CONNECTION_ID
	buf := make([]byte, socksz.LengthConnectionID)
	n, err := io.ReadFull(reader, buf)
	if n != socksz.LengthConnectionID || err != nil {
		return fmt.Errorf("failed to read connection id:  %s", err)
	}
	r.ConnectionID = string(buf)

	// TARGET_USER_CLIENT_ID
	buf = make([]byte, socksz.LengthTargetUserClientID)
	n, err = io.ReadFull(reader, buf)
	if n != socksz.LengthTargetUserClientID || err != nil {
		return fmt.Errorf("failed to read target user client id:  %s", err)
	}
	r.TargetUserClientID = string(buf)

	// TARGET_USER_PAIR_KEY
	buf = make([]byte, socksz.LengthTargetUserPairSignature)
	n, err = io.ReadFull(reader, buf)
	if n != socksz.LengthTargetUserPairSignature || err != nil {
		return fmt.Errorf("failed to read target user pair key:  %s", err)
	}
	r.TargetUserPairSignature = string(buf)

	// NETWORK
	buf = make([]byte, socksz.LengthNetwork)
	n, err = io.ReadFull(reader, buf)
	if n != socksz.LengthNetwork || err != nil {
		return fmt.Errorf("failed to read signature:  %s", err)
	}
	r.Network = uint8(buf[0])

	// ATYP
	buf = make([]byte, socksz.LengthATyp)
	n, err = io.ReadFull(reader, buf)
	if n != socksz.LengthATyp || err != nil {
		return fmt.Errorf("failed to read atyp:  %s", err)
	}
	r.ATyp = uint8(buf[0])

	// DSTAddr
	buf = make([]byte, 1)
	n, err = io.ReadFull(reader, buf)
	if n != 1 || err != nil {
		return fmt.Errorf("failed to read dst addr socksz.length:  %s", err)
	}
	dstAddrLength := int(buf[0])
	buf = make([]byte, dstAddrLength)
	n, err = io.ReadFull(reader, buf)
	if n != dstAddrLength || err != nil {
		return fmt.Errorf("failed to read dst addr:  %s", err)
	}
	r.DSTAddr = string(buf)

	// DSTAddrPort
	buf = make([]byte, socksz.LengthDSTPort)
	n, err = io.ReadFull(reader, buf)
	if n != socksz.LengthDSTPort || err != nil {
		return fmt.Errorf("failed to read atyp:  %s", err)
	}
	r.DSTPort = binary.BigEndian.Uint16(buf[:2])

	return nil
}

func (r *Request) Verify() error {
	if r.Secret == "" {
		return fmt.Errorf("secret is required")
	}

	// verify
	if r.TargetUserPairSignature != hmac.Sha256(fmt.Sprintf("%s_%s", r.ConnectionID, r.TargetUserClientID), r.Secret) {
		return fmt.Errorf("invalid signature")
	}

	return nil
}
