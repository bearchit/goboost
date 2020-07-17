package coder

import (
	"bytes"
	"fmt"
	"regexp"
)

type prefixOption struct {
	delim []byte
}

type PrefixOptionFunc func(*prefixOption)

func PrefixDelim(delim []byte) PrefixOptionFunc {
	return func(option *prefixOption) {
		option.delim = delim
	}
}

func NewPrefixEncoder(prefix []byte, optionFn ...PrefixOptionFunc) EncoderFunc {
	var option prefixOption
	for _, fn := range optionFn {
		fn(&option)
	}
	if len(option.delim) <= 0 {
		option.delim = []byte(":")
	}

	return func(src []byte) ([]byte, error) {
		buf := new(bytes.Buffer)
		buf.Write(prefix)
		buf.Write(option.delim)
		buf.Write(src)

		return buf.Bytes(), nil
	}
}

const (
	PayloadPrefix = "prefix.prefix"
)

func NewPrefixDecoder(optionFn ...PrefixOptionFunc) DecoderFunc {
	var option prefixOption
	for _, fn := range optionFn {
		fn(&option)
	}
	if len(option.delim) <= 0 {
		option.delim = []byte(":")
	}

	return func(src []byte) ([]byte, Payload, error) {
		pattern := regexp.MustCompile(fmt.Sprintf(`(.+)%s(.+)`, option.delim))
		matches := pattern.FindSubmatch(src)
		payload := NewPayload()

		if err := payload.Set(PayloadPrefix, matches[1]); err != nil {
			return nil, nil, err
		}

		return matches[2], payload, nil
	}
}
