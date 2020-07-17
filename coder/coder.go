package coder

import (
	"errors"
	"fmt"
)

type Encoder interface {
	Encode(src []byte) ([]byte, error)
}

type EncoderFunc func(src []byte) ([]byte, error)

func (f EncoderFunc) Encode(src []byte) ([]byte, error) {
	return f(src)
}

type Decoder interface {
	Decode(src []byte) ([]byte, Payload, error)
}

type DecoderFunc func(src []byte) ([]byte, Payload, error)

func (f DecoderFunc) Decode(v []byte) ([]byte, Payload, error) {
	return f(v)
}

type Payload map[string]interface{}

var (
	ErrPayloadExists   = errors.New("payload already exists")
	ErrPayloadNotFound = errors.New("payload not found")
)

func (p Payload) Set(key string, value interface{}) error {
	if _, ok := p[key]; ok {
		return fmt.Errorf("%s, %w", key, ErrPayloadExists)
	}
	p[key] = value
	return nil
}

func (p Payload) Get(key string) (interface{}, error) {
	v, ok := p[key]
	if !ok {
		return nil, fmt.Errorf("%s, %w", key, ErrPayloadNotFound)
	}

	return v, nil
}

// Compose compose functions execute b after a
func composeEncoderFunc(a, b EncoderFunc) EncoderFunc {
	return func(src []byte) ([]byte, error) {
		res, err := a(src)
		if err != nil {
			return nil, err
		}
		return b(res)
	}
}

func ChainEncoder(fn ...EncoderFunc) EncoderFunc {
	return func(src []byte) ([]byte, error) {
		var res EncoderFunc
		res = fn[0]
		for i := 1; i < len(fn); i++ {
			res = composeEncoderFunc(res, fn[i])
		}
		return res(src)
	}
}

func composeDecoderFunc(a, b DecoderFunc) DecoderFunc {
	return func(src []byte) ([]byte, Payload, error) {
		var (
			err     error
			payload = make(Payload)
		)

		ra, pa, err := a(src)
		if err != nil {
			return nil, nil, err
		}
		for k, v := range pa {
			payload[k] = v
		}

		rb, pb, err := b(ra)
		if err != nil {
			return nil, nil, err
		}
		for k, v := range pb {
			payload[k] = v
		}

		return rb, payload, nil
	}
}

func ChainDecoder(fn ...DecoderFunc) DecoderFunc {
	return func(src []byte) ([]byte, Payload, error) {
		var res DecoderFunc
		res = fn[0]
		for i := 1; i < len(fn); i++ {
			res = composeDecoderFunc(res, fn[i])
		}
		return res(src)
	}
}
