package namespace

import (
	"fmt"
	"regexp"

	"github.com/bearchit/goboost/id"
)

const (
	defaultDelimiter = ":"
)

const (
	PayloadKeyNamespace = "namespace.namespace"
)

type prefixCodec struct {
	namespace string
	delim     string
}

type OptionFunc func(codec *prefixCodec)

func WithDelimiter(delim string) OptionFunc {
	return func(codec *prefixCodec) {
		codec.delim = delim
	}
}

func NewPrefixCodec(namespace string, option ...OptionFunc) id.Codec {
	codec := &prefixCodec{namespace: namespace}
	for _, x := range option {
		x(codec)
	}
	if codec.delim == "" {
		codec.delim = defaultDelimiter
	}

	return codec
}

func (c prefixCodec) Encode(src id.ID) (id.ID, error) {
	return EncodePrefix(src, c.namespace, c.delim)
}

func (c prefixCodec) Decode(src id.ID) (id.ID, id.Payload, error) {
	return DecodePrefix(src, c.delim)
}

func EncodePrefix(src id.ID, namespace, delim string) (id.ID, error) {
	return id.FromString(fmt.Sprintf("%s%s%s", namespace, delim, src)), nil
}

func DecodePrefix(src id.ID, delim string) (id.ID, id.Payload, error) {
	pattern := regexp.MustCompile(fmt.Sprintf(`(.+)%s(.+)`, delim))
	matches := pattern.FindStringSubmatch(src.String())
	if len(matches) != 3 {
		return id.NilID, nil, fmt.Errorf("%s, invalid namespace codec", src)
	}

	decoded := id.FromString(matches[2])
	payload := make(id.Payload)
	payload[PayloadKeyNamespace] = matches[1]

	return decoded, payload, nil
}
