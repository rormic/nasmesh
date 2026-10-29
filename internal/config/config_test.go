package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigParse(t *testing.T) {
	t.Run("missing required field returns error", func(t *testing.T) {
		_, err := Load()
		assert.Error(t, err)
	})

	t.Run("all fields set returns correct values", func(t *testing.T) {
		t.Setenv("EXAMPLE", "test-example")

		c, err := Load()
		require.NoError(t, err)

		assert.Equal(t, "test-example", c.Example)
	})
}
