package core

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func addSub(a, b int) (int, int) {
	return a + b, a - b
}

func errStr(err error) string {
	return fmt.Sprintf("%+v", err)
}

func Test_Call(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		res, err := Call(addSub, 1, 2)
		assert.NoError(t, err)
		assert.Equal(t, 3, res[0])
		assert.Equal(t, -1, res[1])
	})
	t.Run("normal - interface", func(t *testing.T) {
		res, err := Call(errStr, errors.New("hoge"))
		assert.NoError(t, err)
		assert.Equal(t, "hoge", res[0])
	})
	t.Run("error - not func", func(t *testing.T) {
		res, err := Call(1, 1, 2)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
	t.Run("error - too many arg", func(t *testing.T) {
		res, err := Call(addSub, 1, 2, 3)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
	t.Run("error - type error", func(t *testing.T) {
		res, err := Call(addSub, 1.0, 2.0)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}
