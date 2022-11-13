package forward

import "testing"

func TestEncodeDecode(t *testing.T) {
	packet := &Forward{
		Data: []byte("hello, Zero"),
	}

	encoded, err := packet.Encode()
	if err != nil {
		t.Fatalf("failed to encode %s", err)
	}

	decoded := &Forward{}
	if err := decoded.Decode(encoded); err != nil {
		t.Fatalf("failed to decode %s", err)
	}

	if string(decoded.Data) != string(packet.Data) {
		t.Fatalf("DATA not match, expect %s, but got %s", packet.Data, decoded.Data)
	}

}
