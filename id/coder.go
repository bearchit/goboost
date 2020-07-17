package id

import "github.com/bearchit/goboost/coder"

func must(err error) {
	if err != nil {
		panic(err)
	}
}

type Generator interface {
	Generate() (ID, error)
	MustGenerate() ID

	GenerateWith(initialID ID) (ID, error)
	MustGenerateWith(initialID ID) ID
}

type generator struct {
	coder.Encoder
}

func NewGenerator(enc coder.Encoder) Generator {
	return &generator{Encoder: enc}
}

func (encoder generator) Generate() (ID, error) {
	return encoder.GenerateWith(NilID)
}

func (encoder generator) GenerateWith(initialID ID) (ID, error) {
	encoded, err := encoder.Encoder.Encode([]byte(initialID))
	if err != nil {
		return NilID, err
	}
	return FromBytes(encoded), nil
}

func (encoder generator) MustGenerate() ID {
	encoded, err := encoder.Generate()
	must(err)
	return encoded
}

func (encoder generator) MustGenerateWith(initialID ID) ID {
	encoded, err := encoder.GenerateWith(initialID)
	must(err)
	return encoded
}

type Parser interface {
	Parse(id ID) (ID, coder.Payload, error)
	MustParse(id ID) (ID, coder.Payload)
}

type parser struct {
	coder.Decoder
}

func NewParser(dec coder.Decoder) Parser {
	return &parser{Decoder: dec}
}

func (decoder parser) Parse(id ID) (ID, coder.Payload, error) {
	decoded, payload, err := decoder.Decoder.Decode([]byte(id))
	if err != nil {
		return NilID, nil, err
	}

	return FromBytes(decoded), payload, nil
}

func (decoder parser) MustParse(id ID) (ID, coder.Payload) {
	decoded, payload, err := decoder.Parse(id)
	must(err)
	return decoded, payload
}

func NewStubGenerator(mockedID ID) Generator {
	return NewGenerator(coder.EncoderFunc(func(src []byte) ([]byte, error) {
		return []byte(mockedID), nil
	}))
}
