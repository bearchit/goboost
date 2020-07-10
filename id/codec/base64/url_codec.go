package base64

import (
	"encoding/base64"
	"fmt"

	"github.com/bearchit/goboost/id"
)

type urlCodec struct{}

func NewURLCodec() id.Codec {
	return &urlCodec{}
}

func (c urlCodec) Encode(src id.ID) (id.ID, error) {
	return EncodeURL(src)
}

func (c urlCodec) Decode(src id.ID) (id.ID, id.Payload, error) {
	return DecodeURL(src)
}

func EncodeURL(src id.ID) (id.ID, error) {
	if src.Nil() {
		return id.NilID, fmt.Errorf("id: %s, no initial ID given", src)
	}
	return id.FromString(base64.URLEncoding.EncodeToString([]byte(src))), nil
}

func DecodeURL(src id.ID) (id.ID, id.Payload, error) {
	decoded, err := base64.URLEncoding.DecodeString(src.String())
	if err != nil {
		return id.NilID, nil, err
	}
	return id.FromString(string(decoded)), nil, nil
}
