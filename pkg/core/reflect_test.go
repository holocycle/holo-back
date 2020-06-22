package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AddSub(a, b int) (int, int) {
	return a + b, a - b
}

func Test_Call(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		res, err := Call(AddSub, 1, 2)
		assert.NoError(t, err)
		assert.Equal(t, 3, res[0])
		assert.Equal(t, -1, res[1])
	})
	t.Run("error - not func", func(t *testing.T) {
		res, err := Call(1, 1, 2)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
	t.Run("error - too many arg", func(t *testing.T) {
		res, err := Call(AddSub, 1, 2, 3)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
	t.Run("error - type error", func(t *testing.T) {
		res, err := Call(AddSub, 1.0, 2.0)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
