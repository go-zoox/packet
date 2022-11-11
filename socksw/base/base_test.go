package close

import "testing"

func TestEncodeDecode(t *testing.T) {
	packet := &Base{
		VER:         0x01,
		CMD:         0x02,
		CRYPTO:      0x03,
		COMPRESSION: 0x04,
		DATA:        []byte("hello zero"),
	}

	encoded, err := packet.Encode()
	if err != nil {
		t.Fatalf("failed to encode %s", err)
	}

	decoded := &Base{}
	if err := decoded.Decode(encoded); err != nil {
		t.Fatalf("failed to decode %s", err)
	}

	if decoded.VER != packet.VER {
		t.Fatalf("VER not match, expect %d, but got %d", packet.VER, decoded.VER)
	}

	if decoded.CMD != packet.CMD {
		t.Fatalf("CMD not match, expect %d, but got %d", packet.CMD, decoded.CMD)
	}

	if decoded.CRYPTO != packet.CRYPTO {
		t.Fatalf("CRYPTO not match, expect %d, but got %d", packet.CRYPTO, decoded.CRYPTO)
	}

	if decoded.COMPRESSION != packet.COMPRESSION {
		t.Fatalf("COMPRESSION not match, expect %d, but got %d", packet.COMPRESSION, decoded.COMPRESSION)
	}

	if string(decoded.DATA) != string(packet.DATA) {
		t.Fatalf("COMPRESSION not match, expect %s, but got %s", packet.DATA, decoded.DATA)
	}
}
