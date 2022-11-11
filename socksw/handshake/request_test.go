package handshake

import "testing"

func TestRequestEncodeDecode(t *testing.T) {
	packet := &Request{
		CONNECTION_ID:              "20ed2884bde9d7565dbf1",
		TARGET_USER_CLIENT_ID:      "b0a501e947",
		TARGET_USER_PAIR_SIGNATURE: "64be94245dd12f7d6d2d5f95839ecd6c50a3887f58edce5a0cb03a85dba505bd",
		NETWORK:                    0x01,
		ATYP:                       0x01,
		DST_ADDR:                   "1.1.1.1",
		DST_PORT:                   80,
	}

	encoded, err := packet.Encode()
	if err != nil {
		t.Fatalf("failed to encode %s", err)
	}

	decoded := &Request{}
	if err := decoded.Decode(encoded); err != nil {
		t.Fatalf("failed to decode %s", err)
	}

	if decoded.CONNECTION_ID != packet.CONNECTION_ID {
		t.Fatalf("ConnectionID not match, expect %s, but got %s", packet.CONNECTION_ID, decoded.CONNECTION_ID)
	}

	if decoded.TARGET_USER_CLIENT_ID != packet.TARGET_USER_CLIENT_ID {
		t.Fatalf("TargetUserClientID not match, expect %s, but got %s", packet.TARGET_USER_CLIENT_ID, decoded.TARGET_USER_CLIENT_ID)
	}

	if decoded.TARGET_USER_PAIR_SIGNATURE != packet.TARGET_USER_PAIR_SIGNATURE {
		t.Fatalf("TargetUserPairKey not match, expect %s, but got %s", packet.TARGET_USER_PAIR_SIGNATURE, decoded.TARGET_USER_PAIR_SIGNATURE)
	}

	if decoded.NETWORK != packet.NETWORK {
		t.Fatalf("Network not match, expect %d, but got %d", packet.NETWORK, decoded.NETWORK)
	}

	if decoded.ATYP != packet.ATYP {
		t.Fatalf("ATyp not match, expect %d, but got %d", packet.ATYP, decoded.ATYP)
	}

	if decoded.DST_ADDR != packet.DST_ADDR {
		t.Fatalf("DSTAddr not match, expect %s, but got %s", packet.DST_ADDR, decoded.DST_ADDR)
	}

	if decoded.DST_PORT != packet.DST_PORT {
		t.Fatalf("DSTPort not match, expect %d, but got %d", packet.DST_PORT, decoded.DST_PORT)
	}
}
