package id

import (
	"github.com/bearchit/goboost/coder"
)

func NewGlobalGenerator(
	namespace string,
) Generator {
	return NewGenerator(
		coder.ChainEncoder(
			coder.NewUuidV4Encoder(),
			coder.NewPrefixEncoder([]byte(namespace)),
			coder.NewBase64URLEncoder(),
		),
	)
}

func NewGlobalParser() Parser {
	return NewParser(
		coder.ChainDecoder(
			coder.NewBase64URLDecoder(),
			coder.NewPrefixDecoder(),
		),
	)
}

func GlobalNamespace(payload coder.Payload) string {
	p, err := payload.Get(coder.PayloadPrefix)
	if err != nil {
		panic(err)
	}
	return string(p.([]byte))
}
