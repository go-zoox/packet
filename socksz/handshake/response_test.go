package handshake

import (
	"testing"

	"github.com/go-zoox/packet/socksz"
	"github.com/go-zoox/random"
)

func TestResponseEncodeDecode(t *testing.T) {
	packet := &Response{
		ConnectionID: random.String(socksz.LengthConnectionID),
		Status:       0x01,
		Message:      "hello world",
	}

	encoded, err := packet.Encode()
	if err != nil {
		t.Fatalf("failed to encode %s", err)
	}

	decoded := &Response{}
	if err := decoded.Decode(encoded); err != nil {
		t.Fatalf("failed to decode %s", err)
	}

	if decoded.Status != packet.Status {
		t.Fatalf("Status not match, expect %d, but got %d", packet.Status, decoded.Status)
	}

	if decoded.Message != packet.Message {
		t.Fatalf("Message not match, expect %s, but got %s", packet.Message, decoded.Message)
	}
}
