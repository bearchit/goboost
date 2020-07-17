package coder

import (
	"encoding/base64"
)

func NewBase64URLEncoder() EncoderFunc {
	return func(src []byte) ([]byte, error) {
		dst := make([]byte, base64.URLEncoding.EncodedLen(len(src)))
		base64.URLEncoding.Encode(dst, src)
		return dst, nil
	}
}

func NewBase64URLDecoder() DecoderFunc {
	return func(src []byte) ([]byte, Payload, error) {
		decoded := make([]byte, base64.URLEncoding.DecodedLen(len(src)))
		n, err := base64.URLEncoding.Decode(decoded, src)
		if err != nil {
			return nil, nil, err
		}
		return decoded[:n], nil, nil
	}
}
