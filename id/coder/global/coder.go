package global

import (
	"github.com/bearchit/goboost/id"
	"github.com/bearchit/goboost/id/codec/base64"
	"github.com/bearchit/goboost/id/codec/namespace"
	"github.com/bearchit/goboost/id/codec/uuid"
)

type option struct {
	delim string
}

type OptionFunc func(option *option)

func WithDelimiter(delim string) OptionFunc {
	return func(option *option) {
		option.delim = delim
	}
}

func NewCoder(
	ns string,
	optionFunc ...OptionFunc,
) id.Coder {
	option := new(option)
	for _, f := range optionFunc {
		f(option)
	}

	namespaceOption := make([]namespace.OptionFunc, 0)
	if option.delim != "" {
		namespaceOption = append(namespaceOption, namespace.WithDelimiter(option.delim))
	}

	return id.NewCoder(
		uuid.NewV4Codec(),
		namespace.NewPrefixCodec(ns, namespaceOption...),
		base64.NewURLCodec(),
	)
}
