package config

//go:generate go run github.com/golang/mock/mockgen -source=config.go -package=mocks -destination=./mocks/config.go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/creasty/defaults"
	"github.com/go-playground/validator"
	"github.com/kelseyhightower/envconfig"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

type Engine struct {
	scanners []Scanner
}

func (e *Engine) AddScanner(scanners ...Scanner) {
	e.scanners = append(e.scanners, scanners...)
}

func New(options ...func(*Engine)) *Engine {
	e := new(Engine)
	for _, option := range options {
		option(e)
	}

	return e
}

func WithScanners(scanners ...Scanner) func(*Engine) {
	return func(engine *Engine) {
		engine.AddScanner(scanners...)
	}
}

func (e Engine) Unmarshal(v interface{}) error {
	for _, scanner := range e.scanners {
		if err := scanner.Struct(v); err != nil {
			if scanner.BreakOnError() {
				return err
			}
		}
	}
	return nil
}

type Scanner interface {
	Struct(v interface{}) error
	BreakOnError() bool
}

type ymlScanner struct {
	filePath     string
	breakOnError bool
}

func (s ymlScanner) Struct(v interface{}) error {
	fb, err := ioutil.ReadFile(s.filePath)
	if err != nil {
		return err
	}
	m := make(map[string]interface{})
	if err := yaml.Unmarshal(fb, m); err != nil {
		return err
	}
	return mapstructure.Decode(m, v)
}

func (s ymlScanner) BreakOnError() bool {
	return s.breakOnError
}

func NewYMLScanner(
	filePath string,
	breakOnError bool,
) Scanner {
	return &ymlScanner{
		filePath:     filePath,
		breakOnError: breakOnError,
	}
}

type envScanner struct {
	prefix       string
	breakOnError bool
}

func (s envScanner) Struct(v interface{}) error {
	return envconfig.Process(s.prefix, v)
}

func (s envScanner) BreakOnError() bool {
	return s.breakOnError
}

func NewEnvScanner(
	prefix string,
	breakOnError bool,
) Scanner {
	return &envScanner{
		prefix:       prefix,
		breakOnError: breakOnError,
	}
}

func NewDefaultEnvDrivenScanners(envPrefix string, yml string) []Scanner {
	return []Scanner{
		NewYMLScanner(yml, false),
		NewEnvScanner(envPrefix, true),
	}
}

func Load(v interface{}, scanners ...Scanner) error {
	if err := defaults.Set(v); err != nil {
		return fmt.Errorf("failed to set default values: %w", err)
	}

	c := New(WithScanners(scanners...))
	if err := c.Unmarshal(v); err != nil {
		return fmt.Errorf("failed to parse configurations: %w", err)
	}

	if err := validator.New().Struct(v); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	return nil
}

func JSON(v interface{}) string {
	j, _ := json.MarshalIndent(v, "", "  ")
	return string(j)
}
