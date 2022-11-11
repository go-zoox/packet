package connect

import "testing"

func TestResponseEncodeDecode(t *testing.T) {
	packet := &Response{
		VER:       0x05,
		REP:       0x00,
		RSV:       0x00,
		ATYP:      0x01,
		BIND_ADDR: "127.0.0.1",
		BIND_PORT: 8080,
	}

	encoded, err := packet.Encode()
	if err != nil {
		t.Fatalf("failed to encode %s", err)
	}

	decoded := &Response{}
	if err := decoded.Decode(encoded); err != nil {
		t.Fatalf("failed to decode %s", err)
	}

	if decoded.VER != packet.VER {
		t.Fatalf("VER not match, expect %d, but got %d", packet.VER, decoded.VER)
	}

	if decoded.REP != packet.REP {
		t.Fatalf("REP not match, expect %d, but got %d", packet.REP, decoded.REP)
	}

	if decoded.RSV != packet.RSV {
		t.Fatalf("RSV not match, expect %d, but got %d", packet.RSV, decoded.RSV)
	}

	if decoded.ATYP != packet.ATYP {
		t.Fatalf("ATYP not match, expect %d, but got %d", packet.ATYP, decoded.ATYP)
	}

	if decoded.BIND_ADDR != packet.BIND_ADDR {
		t.Fatalf("BIND_ADDR not match, expect %s, but got %s", packet.BIND_ADDR, decoded.BIND_ADDR)
	}

	if decoded.BIND_PORT != packet.BIND_PORT {
		t.Fatalf("BIND_PORT not match, expect %d, but got %d", packet.BIND_PORT, decoded.BIND_PORT)
	}
}
