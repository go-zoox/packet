package handshake

import "testing"

func TestRequestEncodeDecode(t *testing.T) {
	packet := &Request{
		ConnectionID:            "20ed2884bde9d7565dbf1",
		TargetUserClientID:      "b0a501e947",
		TargetUserPairSignature: "64be94245dd12f7d6d2d5f95839ecd6c50a3887f58edce5a0cb03a85dba505bd",
		Network:                 0x01,
		ATyp:                    0x01,
		DSTAddr:                 "1.1.1.1",
		DSTPort:                 80,
	}

	encoded, err := packet.Encode()
	if err != nil {
		t.Fatalf("failed to encode %s", err)
	}

	decoded := &Request{}
	if err := decoded.Decode(encoded); err != nil {
		t.Fatalf("failed to decode %s", err)
	}

	if decoded.ConnectionID != packet.ConnectionID {
		t.Fatalf("ConnectionID not match, expect %s, but got %s", packet.ConnectionID, decoded.ConnectionID)
	}

	if decoded.TargetUserClientID != packet.TargetUserClientID {
		t.Fatalf("TargetUserClientID not match, expect %s, but got %s", packet.TargetUserClientID, decoded.TargetUserClientID)
	}

	if decoded.TargetUserPairSignature != packet.TargetUserPairSignature {
		t.Fatalf("TargetUserPairKey not match, expect %s, but got %s", packet.TargetUserPairSignature, decoded.TargetUserPairSignature)
	}

	if decoded.Network != packet.Network {
		t.Fatalf("Network not match, expect %d, but got %d", packet.Network, decoded.Network)
	}

	if decoded.ATyp != packet.ATyp {
		t.Fatalf("ATyp not match, expect %d, but got %d", packet.ATyp, decoded.ATyp)
	}

	if decoded.DSTAddr != packet.DSTAddr {
		t.Fatalf("DSTAddr not match, expect %s, but got %s", packet.DSTAddr, decoded.DSTAddr)
	}

	if decoded.DSTPort != packet.DSTPort {
		t.Fatalf("DSTPort not match, expect %d, but got %d", packet.DSTPort, decoded.DSTPort)
	}
}
