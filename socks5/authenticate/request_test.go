package authenticate

import "testing"

func TestRequestEncodeDecode(t *testing.T) {
	packet := &Request{
		Ver:      0x05,
		NMethods: 0x02,
		Methods:  []byte{0x01, 0x02},
	}

	encoded, err := packet.Encode()
	if err != nil {
		t.Fatalf("failed to encode %s", err)
	}

	decoded := &Request{}
	if err := decoded.Decode(encoded); err != nil {
		t.Fatalf("failed to decode %s", err)
	}

	if decoded.Ver != packet.Ver {
		t.Fatalf("VER not match, expect %d, but got %d", packet.Ver, decoded.Ver)
	}

	if decoded.NMethods != packet.NMethods {
		t.Fatalf("NMETHODS not match, expect %d, but got %d", packet.NMethods, decoded.NMethods)
	}

	if string(decoded.Methods) != string(packet.Methods) {
		t.Fatalf("METHODS not match, expect %s, but got %s", packet.Methods, decoded.Methods)
	}
}
