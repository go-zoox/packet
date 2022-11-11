package forward

import "testing"

func TestEncodeDecode(t *testing.T) {
	packet := &Forward{
		ConnectionID: "20ed2884bde9d7565dbf1",
		Data:         []byte("hello world"),
	}

	encoded, err := packet.Encode()
	if err != nil {
		t.Fatalf("failed to encode %s", err)
	}

	decoded := &Forward{}
	if err := decoded.Decode(encoded); err != nil {
		t.Fatalf("failed to decode %s", err)
	}

	if decoded.ConnectionID != packet.ConnectionID {
		t.Fatalf("ConnectionID not match, expect %s, but got %s", packet.ConnectionID, decoded.ConnectionID)
	}

	if string(decoded.Data) != string(packet.Data) {
		t.Fatalf("Data not match, expect %s, but got %s", packet.Data, decoded.Data)
	}
}
