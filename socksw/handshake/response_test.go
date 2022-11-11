package handshake

import "testing"

func TestResponseEncodeDecode(t *testing.T) {
	packet := &Response{
		CONNECTION_ID: "012345678901234567890",
		STATUS:        0x01,
		MESSAGE:       "hello world",
	}

	encoded, err := packet.Encode()
	if err != nil {
		t.Fatalf("failed to encode %s", err)
	}

	decoded := &Response{}
	if err := decoded.Decode(encoded); err != nil {
		t.Fatalf("failed to decode %s", err)
	}

	if decoded.STATUS != packet.STATUS {
		t.Fatalf("Status not match, expect %d, but got %d", packet.STATUS, decoded.STATUS)
	}

	if decoded.MESSAGE != packet.MESSAGE {
		t.Fatalf("Message not match, expect %s, but got %s", packet.MESSAGE, decoded.MESSAGE)
	}
}
