package id

import (
	"fmt"
)

type Encoder interface {
	Encode(src ID) (ID, error)
}

type EncoderFunc func(ID) (ID, error)

func (f EncoderFunc) Encode(src ID) (ID, error) {
	return f(src)
}

type Payload map[string]interface{}

func (p Payload) IsZero() bool {
	return len(p) == 0
}

type Decoder interface {
	Decode(src ID) (ID, Payload, error)
}

type DecoderFunc func(ID) (ID, Payload, error)

func (f DecoderFunc) Decode(src ID) (ID, Payload, error) {
	return f(src)
}

type Codec interface {
	Encoder
	Decoder
}

type Coder interface {
	Encode() (ID, error)
	MustEncode() ID

	EncodeWith(initialID ID) (ID, error)
	MustEncodeWith(initialID ID) ID

	Decode(id ID) (ID, Payload, error)
	MustDecode(id ID) (ID, Payload)
}

type coder struct {
	initialID ID
	codecs    []Codec
}

func NewCoder(
	codecs ...Codec,
) Coder {
	return &coder{
		codecs: codecs,
	}
}

func (e coder) Encode() (ID, error) {
	var (
		id  ID
		err error
	)

	for _, codec := range e.codecs {
		id, err = codec.Encode(id)
		if err != nil {
			return NilID, err
		}
	}

	return id, nil
}

func (e coder) MustEncode() ID {
	id, err := e.Encode()
	if err != nil {
		panic(err)
	}
	return id
}

func (e coder) EncodeWith(initialID ID) (ID, error) {
	var err error

	for _, codec := range e.codecs {
		initialID, err = codec.Encode(initialID)
		if err != nil {
			return NilID, err
		}
	}
	return initialID, nil
}

func (e coder) MustEncodeWith(initialID ID) ID {
	id, err := e.EncodeWith(initialID)
	if err != nil {
		panic(err)
	}
	return id
}

func (e coder) Decode(id ID) (ID, Payload, error) {
	var (
		err           error
		entirePayload = make(Payload)
	)

	for i := len(e.codecs) - 1; i >= 0; i-- {
		var payload Payload
		id, payload, err = e.codecs[i].Decode(id)
		if err != nil {
			return NilID, nil, err
		}

		for k, v := range payload {
			if _, ok := entirePayload[k]; ok {
				return NilID, nil, fmt.Errorf("%s already exists", k)
			}
			entirePayload[k] = v
		}
	}

	return id, entirePayload, nil
}

func (e coder) MustDecode(id ID) (ID, Payload) {
	decoded, payload, err := e.Decode(id)
	if err != nil {
		panic(err)
	}
	return decoded, payload
}
