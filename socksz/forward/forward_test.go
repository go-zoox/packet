package forward

import (
	"testing"

	"github.com/go-zoox/packet/socksz"
	"github.com/go-zoox/random"
)

func TestEncodeDecode(t *testing.T) {
	packet := &Forward{
		ConnectionID: random.String(socksz.LengthConnectionID),
		Data:         []byte("hello world"),
	}

	encoded, err := packet.Encode()
	if err != nil {
		t.Fatalf("failed to encode %s", err)
	}

	decoded := &Forward{}
	if err := decoded.Decode(encoded); err != nil {
		t.Fatalf("failed to decode %s", err)
	}

	if decoded.ConnectionID != packet.ConnectionID {
		t.Fatalf("ConnectionID not match, expect %s, but got %s", packet.ConnectionID, decoded.ConnectionID)
	}

	if string(decoded.Data) != string(packet.Data) {
		t.Fatalf("Data not match, expect %s, but got %s", packet.Data, decoded.Data)
	}
}

func TestEncodeDecodeWithCrypto(t *testing.T) {
	secret := "the_secret"

	algorithms := map[string]uint8{
		"aes-128-cfb":       0x01,
		"aes-192-cfb":       0x02,
		"aes-256-cfb":       0x03,
		"aes-128-cbc":       0x04,
		"aes-192-cbc":       0x05,
		"aes-256-cbc":       0x06,
		"aes-128-gcm":       0x07,
		"aes-192-gcm":       0x08,
		"aes-256-gcm":       0x09,
		"chacha20-poly1305": 0x10,
	}

	for key, value := range algorithms {
		packet := &Forward{
			Crypto: value,
			Secret: secret,
			//
			ConnectionID: random.String(socksz.LengthConnectionID),
			Data:         []byte("hello world"),
		}

		encoded, err := packet.Encode()
		if err != nil {
			t.Fatalf("[crypto: %s] failed to encode %s", key, err)
		}

		decoded := &Forward{
			Crypto: value,
			Secret: secret,
			//
		}
		if err := decoded.Decode(encoded); err != nil {
			t.Fatalf("[crypto: %s] failed to decode %s", key, err)
		}

		if decoded.ConnectionID != packet.ConnectionID {
			t.Fatalf("[crypto: %s] ConnectionID not match, expect %s, but got %s", key, packet.ConnectionID, decoded.ConnectionID)
		}

		if string(decoded.Data) != string(packet.Data) {
			t.Fatalf("[crypto: %s] Data not match, expect %s, but got %s", key, packet.Data, decoded.Data)
		}
	}
}
