package onetwo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputCovert(t *testing.T) {
	t.Run("input should be 4-digit", func(t *testing.T) {
		input := "1234"
		ans, err := InputCovert(input)
		assert.NoError(t, err, "input should be 4-digit")
		assert.Len(t, ans, 4, "input should be 4-digit")
	})

	t.Run("input should not be 3-digit", func(t *testing.T) {
		input := "123"
		_, err := InputCovert(input)
		assert.Error(t, err, "input should not be 3-digit")
	})

	t.Run("input should not be 5-digit", func(t *testing.T) {
		input := "12345"
		_, err := InputCovert(input)
		assert.Error(t, err, "input should not be 5-digit")
	})

	t.Run("input should not be 4-letter", func(t *testing.T) {
		input := "abcd"
		_, err := InputCovert(input)
		assert.Error(t, err, "input should not be 4-letter")
	})

	t.Run("input should not be 4-symbol", func(t *testing.T) {
		input := "!@#$"
		_, err := InputCovert(input)
		assert.Error(t, err, "input should not be 4-symbol")
	})

	t.Run("input should not be 4-mixed", func(t *testing.T) {
		input := "a1b2"
		_, err := InputCovert(input)
		assert.Error(t, err, "input should not be 4-mixed")
	})
}
