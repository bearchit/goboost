package config

import (
	"github.com/go-errors/errors"
	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
)

type DotEnvScanner struct {
	breakOnError bool
}

func NewDotEnvScanner(
	breakOnError bool,
) *DotEnvScanner {
	return &DotEnvScanner{
		breakOnError: breakOnError,
	}
}

func (s DotEnvScanner) Struct(v interface{}) error {
	envs, err := godotenv.Read()
	if err != nil {
		return errors.Wrapf("failed to scan from env file", err)
	}

	if err := mapstructure.Decode(&envs, v); err != nil {
		return errors.Wrapf("failed to scan from env file", err)
	}

	return nil
}

func (s DotEnvScanner) BreakOnError() bool {
	return s.breakOnError
}
