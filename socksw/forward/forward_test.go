package forward

import "testing"

func TestEncodeDecode(t *testing.T) {
	packet := &Forward{
		CONNECTION_ID: "20ed2884bde9d7565dbf1",
		DATA:          []byte("hello world"),
	}

	encoded, err := packet.Encode()
	if err != nil {
		t.Fatalf("failed to encode %s", err)
	}

	decoded := &Forward{}
	if err := decoded.Decode(encoded); err != nil {
		t.Fatalf("failed to decode %s", err)
	}

	if decoded.CONNECTION_ID != packet.CONNECTION_ID {
		t.Fatalf("ConnectionID not match, expect %s, but got %s", packet.CONNECTION_ID, decoded.CONNECTION_ID)
	}

	if string(decoded.DATA) != string(packet.DATA) {
		t.Fatalf("Data not match, expect %s, but got %s", packet.DATA, decoded.DATA)
	}
}
