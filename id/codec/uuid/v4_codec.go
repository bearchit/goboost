package uuid

import (
	"fmt"

	"github.com/bearchit/goboost/id"

	uuid "github.com/satori/go.uuid"
)

func EncodeV4(src id.ID) (id.ID, error) {
	if !src.Nil() {
		return id.NilID, fmt.Errorf("initialID given")
	}
	return id.FromString(uuid.NewV4().String()), nil
}

func DecodeV4(src id.ID) (id.ID, id.Payload, error) {
	return src, nil, nil
}

type v4Codec struct{}

func NewV4Codec() id.Codec {
	return new(v4Codec)
}

func (c v4Codec) Encode(src id.ID) (id.ID, error) {
	return EncodeV4(src)
}

func (c v4Codec) Decode(src id.ID) (id.ID, id.Payload, error) {
	return DecodeV4(src)
}
