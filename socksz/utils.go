package socksz

import (
	"fmt"

	nanoid "github.com/matoous/go-nanoid/v2"
)

// GenerateID generates a connection id
func GenerateID() string {
	id, _ := nanoid.New(LengthConnectionID)
	return id
}

//
func GetCrypto(algorithm string) (uint8, error) {
	CryptoAlgorithms := map[string]uint8{
		"":                  0x00,
		"aes-128-cfb":       0x01,
		"aes-192-cfb":       0x02,
		"aes-256-cfb":       0x03,
		"aes-128-cbc":       0x04,
		"aes-192-cbc":       0x05,
		"aes-256-cbc":       0x06,
		"aes-128-gcm":       0x07,
		"aes-192-gcm":       0x08,
		"aes-256-gcm":       0x09,
		"chacha20-poly1305": 0x10,
	}

	if v, ok := CryptoAlgorithms[algorithm]; ok {
		return v, nil
	}

	return 0, fmt.Errorf("unknown algorithm: %s", algorithm)

}
