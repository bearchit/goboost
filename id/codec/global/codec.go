package global

import (
	"github.com/bearchit/goboost/coder"
	"github.com/bearchit/goboost/id"
)

func NewGenerator(
	namespace string,
) id.Generator {
	return id.NewGenerator(
		coder.ChainEncoder(
			coder.NewUuidV4Encoder(),
			coder.NewPrefixEncoder(namespace),
			coder.NewBase64URLEncoder(),
		),
	)
}

func NewParser() id.Parser {
	return id.NewParser(
		coder.ChainDecoder(
			coder.NewBase64URLDecoder(),
			coder.NewPrefixDecoder(),
		),
	)
}

func Namespace(payload coder.Payload) string {
	p, err := payload.Get(coder.PayloadPrefix)
	if err != nil {
		panic(err)
	}
	return string(p.([]byte))
}
