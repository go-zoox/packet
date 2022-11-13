package connect

import "testing"

func TestRequestEncodeDecode(t *testing.T) {
	packet := &Request{
		Ver:     0x05,
		Cmd:     0x01,
		Rsv:     0x00,
		ATyp:    0x01,
		DSTAddr: "127.0.0.1",
		DSTPort: 8080,
	}

	encoded, err := packet.Encode()
	if err != nil {
		t.Fatalf("failed to encode %s", err)
	}

	decoded := &Request{}
	if err := decoded.Decode(encoded); err != nil {
		t.Fatalf("failed to decode %s", err)
	}

	if decoded.Ver != packet.Ver {
		t.Fatalf("VER not match, expect %d, but got %d", packet.Ver, decoded.Ver)
	}

	if decoded.Cmd != packet.Cmd {
		t.Fatalf("CMD not match, expect %d, but got %d", packet.Cmd, decoded.Cmd)
	}

	if decoded.Rsv != packet.Rsv {
		t.Fatalf("RSV not match, expect %d, but got %d", packet.Rsv, decoded.Rsv)
	}

	if decoded.ATyp != packet.ATyp {
		t.Fatalf("ATYP not match, expect %d, but got %d", packet.ATyp, decoded.ATyp)
	}

	if decoded.DSTAddr != packet.DSTAddr {
		t.Fatalf("DST_ADDR not match, expect %s, but got %s", packet.DSTAddr, decoded.DSTAddr)
	}

	if decoded.DSTPort != packet.DSTPort {
		t.Fatalf("DST_PORT not match, expect %d, but got %d", packet.DSTPort, decoded.DSTPort)
	}
}
