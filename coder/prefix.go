package coder

import (
	"fmt"
	"regexp"
)

func NewPrefixEncoder(prefix string) EncoderFunc {
	return func(src []byte) ([]byte, error) {
		return []byte(fmt.Sprintf("%s:%s", prefix, src)), nil
	}
}

const (
	PayloadPrefix = "prefix.prefix"
)

func NewPrefixDecoder() DecoderFunc {
	return func(src []byte) ([]byte, Payload, error) {
		pattern := regexp.MustCompile(`(.+):(.+)`)
		matches := pattern.FindSubmatch(src)
		payload := make(Payload)

		if err := payload.Set(PayloadPrefix, matches[1]); err != nil {
			return nil, nil, err
		}

		return matches[2], payload, nil
	}
}
