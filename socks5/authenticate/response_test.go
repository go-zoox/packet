package authenticate

import "testing"

func TestResponseEncodeDecode(t *testing.T) {
	packet := &Response{
		VER:    0x05,
		METHOD: 0x01,
	}

	encoded, err := packet.Encode()
	if err != nil {
		t.Fatalf("failed to encode %s", err)
	}

	decoded := &Response{}
	if err := decoded.Decode(encoded); err != nil {
		t.Fatalf("failed to decode %s", err)
	}

	if decoded.VER != packet.VER {
		t.Fatalf("VER not match, expect %d, but got %d", packet.VER, decoded.VER)
	}

	if decoded.METHOD != packet.METHOD {
		t.Fatalf("METHOD not match, expect %d, but got %d", packet.METHOD, decoded.METHOD)
	}
}
