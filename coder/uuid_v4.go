package coder

import (
	uuid "github.com/satori/go.uuid"
)

func NewUuidV4Encoder() EncoderFunc {
	return func(src []byte) ([]byte, error) {
		return []byte(uuid.NewV4().String()), nil
	}
}
