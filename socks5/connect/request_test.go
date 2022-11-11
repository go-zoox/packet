package connect

import "testing"

func TestRequestEncodeDecode(t *testing.T) {
	packet := &Request{
		VER:     0x05,
		CMD:     0x01,
		RSV:     0x00,
		ATYP:    0x01,
		DSTADDR: "127.0.0.1",
		DSTPORT: 8080,
	}

	encoded, err := packet.Encode()
	if err != nil {
		t.Fatalf("failed to encode %s", err)
	}

	decoded := &Request{}
	if err := decoded.Decode(encoded); err != nil {
		t.Fatalf("failed to decode %s", err)
	}

	if decoded.VER != packet.VER {
		t.Fatalf("VER not match, expect %d, but got %d", packet.VER, decoded.VER)
	}

	if decoded.CMD != packet.CMD {
		t.Fatalf("CMD not match, expect %d, but got %d", packet.CMD, decoded.CMD)
	}

	if decoded.RSV != packet.RSV {
		t.Fatalf("RSV not match, expect %d, but got %d", packet.RSV, decoded.RSV)
	}

	if decoded.ATYP != packet.ATYP {
		t.Fatalf("ATYP not match, expect %d, but got %d", packet.ATYP, decoded.ATYP)
	}

	if decoded.DSTADDR != packet.DSTADDR {
		t.Fatalf("DST_ADDR not match, expect %s, but got %s", packet.DSTADDR, decoded.DSTADDR)
	}

	if decoded.DSTPORT != packet.DSTPORT {
		t.Fatalf("DST_PORT not match, expect %d, but got %d", packet.DSTPORT, decoded.DSTPORT)
	}
}
