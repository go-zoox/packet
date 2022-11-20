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
