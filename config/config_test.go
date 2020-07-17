package config_test

import (
	"os"
	"testing"

	"github.com/bearchit/goboost/config"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestNewEngine(t *testing.T) {
	e := config.New()
	assert.NotNil(t, e)
}

func TestEngine_Env(t *testing.T) {
	c := new(struct {
		Name string
	})

	t.Run("env", func(t *testing.T) {
		os.Setenv("NAME", "boost_config")
		e := config.New(
			config.WithScanners(
				config.NewEnvScanner("", false),
			),
		)

		require.NoError(t, e.Unmarshal(c))
		assert.Equal(t, "boost_config", c.Name)
	})
}

func TestEnvScanner_ErrorOnFail(t *testing.T) {
	newConfig := func(breakOnError bool) *config.Engine {
		return config.New(
			config.WithScanners(
				config.NewYMLScanner("config/test.env", breakOnError)),
		)
	}

	t.Run("skip on error", func(t *testing.T) {
		cfg := newConfig(false)
		assert.NoError(t, cfg.Unmarshal(&struct{}{}))
	})

	t.Run("break on error", func(t *testing.T) {
		cfg := newConfig(true)
		assert.Error(t, cfg.Unmarshal(&struct{}{}))
	})
}
