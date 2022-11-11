package packet

// Packet is a packet
type Packet interface {
	Encode() ([]byte, error)
	Decode(raw []byte) error
}
