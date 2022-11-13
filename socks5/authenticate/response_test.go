package authenticate

import "testing"

func TestResponseEncodeDecode(t *testing.T) {
	packet := &Response{
		Ver:    0x05,
		Method: 0x01,
	}

	encoded, err := packet.Encode()
	if err != nil {
		t.Fatalf("failed to encode %s", err)
	}

	decoded := &Response{}
	if err := decoded.Decode(encoded); err != nil {
		t.Fatalf("failed to decode %s", err)
	}

	if decoded.Ver != packet.Ver {
		t.Fatalf("VER not match, expect %d, but got %d", packet.Ver, decoded.Ver)
	}

	if decoded.Method != packet.Method {
		t.Fatalf("METHOD not match, expect %d, but got %d", packet.Method, decoded.Method)
	}
}
