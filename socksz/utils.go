package socksz

import (
	nanoid "github.com/matoous/go-nanoid/v2"
)

// GenerateID generates a connection id
func GenerateID() string {
	id, _ := nanoid.New(LengthConnectionID)
	return id
}
