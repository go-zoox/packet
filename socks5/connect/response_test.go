package connect

import "testing"

func TestResponseEncodeDecode(t *testing.T) {
	packet := &Response{
		VER:      0x05,
		REP:      0x00,
		RSV:      0x00,
		ATYP:     0x01,
		BINDAddr: "127.0.0.1",
		BINDPort: 8080,
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

	if decoded.BINDAddr != packet.BINDAddr {
		t.Fatalf("BIND_ADDR not match, expect %s, but got %s", packet.BINDAddr, decoded.BINDAddr)
	}

	if decoded.BINDPort != packet.BINDPort {
		t.Fatalf("BIND_PORT not match, expect %d, but got %d", packet.BINDPort, decoded.BINDPort)
	}
}
