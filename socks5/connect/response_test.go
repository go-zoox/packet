package connect

import "testing"

func TestResponseEncodeDecode(t *testing.T) {
	packet := &Response{
		Ver:      0x05,
		Rep:      0x00,
		Rsv:      0x00,
		ATyp:     0x01,
		BindAddr: "127.0.0.1",
		BindPort: 8080,
	}

	encoded, err := packet.Encode()
	if err != nil {
		t.Fatalf("failed to encode %s", err)
	}

	decoded := &Response{}
	if err := decoded.Decode(encoded); err != nil {
		t.Fatalf("failed to decode %s", err)
	}

	if decoded.Ver != packet.Ver {
		t.Fatalf("VER not match, expect %d, but got %d", packet.Ver, decoded.Ver)
	}

	if decoded.Rep != packet.Rep {
		t.Fatalf("REP not match, expect %d, but got %d", packet.Rep, decoded.Rep)
	}

	if decoded.Rsv != packet.Rsv {
		t.Fatalf("RSV not match, expect %d, but got %d", packet.Rsv, decoded.Rsv)
	}

	if decoded.ATyp != packet.ATyp {
		t.Fatalf("ATYP not match, expect %d, but got %d", packet.ATyp, decoded.ATyp)
	}

	if decoded.BindAddr != packet.BindAddr {
		t.Fatalf("BIND_ADDR not match, expect %s, but got %s", packet.BindAddr, decoded.BindAddr)
	}

	if decoded.BindPort != packet.BindPort {
		t.Fatalf("BIND_PORT not match, expect %d, but got %d", packet.BindPort, decoded.BindPort)
	}
}
