package joinAsAgent

import (
	"bytes"
	"fmt"
	"io"

	"github.com/go-zoox/packet/socksz"
)

// DATA Protocol:
//
// AUTHENTICATE DATA:
// response: STATUS | MESSAGE
//            1     |  -

// Response is the response for authenticate
type Response struct {
	Status  uint8
	Message string
}

// Encode encodes the response for authenticate
func (r *Response) Encode() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	err := buf.WriteByte(r.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to write Status:  %s", err)
	}

	_, err = buf.WriteString(r.Message)
	if err != nil {
		return nil, fmt.Errorf("failed to write Message:  %s", err)
	}

	return buf.Bytes(), nil
}

// Decode decodes the response for authenticate
func (r *Response) Decode(raw []byte) error {
	reader := bytes.NewReader(raw)

	// STATUS
	buf := make([]byte, socksz.LengthStatus)
	n, err := io.ReadFull(reader, buf)
	if n != socksz.LengthStatus || err != nil {
		return fmt.Errorf("failed to read status:  %s", err)
	}
	r.Status = uint8(buf[0])

	// Message
	buf, err = io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("failed to read message:  %s", err)
	}
	r.Message = string(buf)

	return nil
}
