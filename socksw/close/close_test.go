package close

import "testing"

func TestEncodeDecode(t *testing.T) {
	packet := &Close{
		ConnectionID: "20ed2884bde9d7565dbf1",
	}

	encoded, err := packet.Encode()
	if err != nil {
		t.Fatalf("failed to encode %s", err)
	}

	decoded := &Close{}
	if err := decoded.Decode(encoded); err != nil {
		t.Fatalf("failed to decode %s", err)
	}

	if decoded.ConnectionID != packet.ConnectionID {
		t.Fatalf("CONNECTION_ID not match, expect %s, but got %s", packet.ConnectionID, decoded.ConnectionID)
	}
}
