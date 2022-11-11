package authenticate

import "testing"

func TestRequestEncodeDecode(t *testing.T) {
	packet := &Request{
		VER:      0x05,
		NMETHODS: 0x02,
		METHODS:  []byte{0x01, 0x02},
	}

	encoded, err := packet.Encode()
	if err != nil {
		t.Fatalf("failed to encode %s", err)
	}

	decoded := &Request{}
	if err := decoded.Decode(encoded); err != nil {
		t.Fatalf("failed to decode %s", err)
	}

	if decoded.VER != packet.VER {
		t.Fatalf("VER not match, expect %d, but got %d", packet.VER, decoded.VER)
	}

	if decoded.NMETHODS != packet.NMETHODS {
		t.Fatalf("NMETHODS not match, expect %d, but got %d", packet.NMETHODS, decoded.NMETHODS)
	}

	if string(decoded.METHODS) != string(packet.METHODS) {
		t.Fatalf("METHODS not match, expect %s, but got %s", packet.METHODS, decoded.METHODS)
	}
}
