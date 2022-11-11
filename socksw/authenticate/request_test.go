package authenticate

import "testing"

func TestRequestEncodeDecode(t *testing.T) {
	packet := &Request{
		USER_CLIENT_ID: "0123456789",
		TIMESTAMP:      "1667982806000",
		NONCE:          "123456",
		SIGNATURE:      "8665ebcb30adc07590ae3209e8cb0c2b9b43762cf6656d95ddb52fbc2a45e39c",
	}

	encoded, err := packet.Encode()
	if err != nil {
		t.Fatalf("failed to encode %s", err)
	}

	decoded := &Request{}
	if err := decoded.Decode(encoded); err != nil {
		t.Fatalf("failed to decode %s", err)
	}

	if decoded.USER_CLIENT_ID != packet.USER_CLIENT_ID {
		t.Fatalf("UserClientID not match, expect %s, but got %s", packet.USER_CLIENT_ID, decoded.USER_CLIENT_ID)
	}

	if decoded.TIMESTAMP != packet.TIMESTAMP {
		t.Fatalf("Timestamp not match, expect %s, but got %s", packet.TIMESTAMP, decoded.TIMESTAMP)
	}

	if decoded.NONCE != packet.NONCE {
		t.Fatalf("Nonce not match, expect %s, but got %s", packet.NONCE, decoded.NONCE)
	}

	if decoded.SIGNATURE != packet.SIGNATURE {
		t.Fatalf("Signature not match, expect %s, but got %s", packet.SIGNATURE, decoded.SIGNATURE)
	}
}
