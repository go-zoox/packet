package base

import "testing"

func TestEncodeDecode(t *testing.T) {
	packet := &Base{
		Ver:         0x01,
		Cmd:         0x02,
		Crypto:      0x00,
		Compression: 0x04,
		Data:        []byte("hello zero"),
	}

	encoded, err := packet.Encode()
	if err != nil {
		t.Fatalf("failed to encode %s", err)
	}

	decoded := &Base{}
	if err := decoded.Decode(encoded); err != nil {
		t.Fatalf("failed to decode %s", err)
	}

	if decoded.Ver != packet.Ver {
		t.Fatalf("VER not match, expect %d, but got %d", packet.Ver, decoded.Ver)
	}

	if decoded.Cmd != packet.Cmd {
		t.Fatalf("CMD not match, expect %d, but got %d", packet.Cmd, decoded.Cmd)
	}

	if decoded.Crypto != packet.Crypto {
		t.Fatalf("CRYPTO not match, expect %d, but got %d", packet.Crypto, decoded.Crypto)
	}

	if decoded.Compression != packet.Compression {
		t.Fatalf("COMPRESSION not match, expect %d, but got %d", packet.Compression, decoded.Compression)
	}

	if string(decoded.Data) != string(packet.Data) {
		t.Fatalf("COMPRESSION not match, expect %s, but got %s", packet.Data, decoded.Data)
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
		packet := &Base{
			Ver:         0x01,
			Cmd:         0x02,
			Crypto:      value, // aes-128-cfb
			Compression: 0x04,
			Data:        []byte("hello zero"),
			Secret:      secret,
		}

		encoded, err := packet.Encode()
		if err != nil {
			t.Fatalf("[crypto][%s] failed to encode %s", key, err)
		}

		decoded := &Base{
			Secret: secret,
		}
		if err := decoded.Decode(encoded); err != nil {
			t.Fatalf("[crypto][%s] failed to decode %s", key, err)
		}

		if decoded.Ver != packet.Ver {
			t.Fatalf("[crypto][%s] VER not match, expect %d, but got %d", key, packet.Ver, decoded.Ver)
		}

		if decoded.Cmd != packet.Cmd {
			t.Fatalf("[crypto][%s] CMD not match, expect %d, but got %d", key, packet.Cmd, decoded.Cmd)
		}

		if decoded.Crypto != packet.Crypto {
			t.Fatalf("[crypto][%s] CRYPTO not match, expect %d, but got %d", key, packet.Crypto, decoded.Crypto)
		}

		if decoded.Compression != packet.Compression {
			t.Fatalf("[crypto][%s] COMPRESSION not match, expect %d, but got %d", key, packet.Compression, decoded.Compression)
		}

		if string(decoded.Data) != string(packet.Data) {
			t.Fatalf("[crypto][%s] COMPRESSION not match, expect %s, but got %s", key, packet.Data, decoded.Data)
		}
	}
}
