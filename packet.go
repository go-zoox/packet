package packet

// packet is a packet
type Packet interface {
	Encode() ([]byte, error)
	Decode(raw []byte) error
}
