package coder

import (
	"bytes"

	uuid "github.com/satori/go.uuid"
)

func NewUuidV4Encoder() EncoderFunc {
	return func(src []byte) ([]byte, error) {
		b := bytes.NewBufferString(uuid.NewV4().String())
		return b.Bytes(), nil
	}
}
