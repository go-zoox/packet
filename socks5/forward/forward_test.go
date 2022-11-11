package forward

import "testing"

func TestEncodeDecode(t *testing.T) {
	packet := &Forward{
		DATA: []byte("hello, Zero"),
	}

	encoded, err := packet.Encode()
	if err != nil {
		t.Fatalf("failed to encode %s", err)
	}

	decoded := &Forward{}
	if err := decoded.Decode(encoded); err != nil {
		t.Fatalf("failed to decode %s", err)
	}

	if string(decoded.DATA) != string(packet.DATA) {
		t.Fatalf("DATA not match, expect %s, but got %s", packet.DATA, decoded.DATA)
	}

}
