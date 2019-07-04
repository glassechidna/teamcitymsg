package teamcitymsg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEscapeField(t *testing.T) {
	t.Run("apostrophe", func(t *testing.T) {
		assert.Equal(t, "|'", EscapeField("'"))
	})

	t.Run("line feed", func(t *testing.T) {
		assert.Equal(t, "|n", EscapeField("\n"))
	})

	t.Run("carriage return", func(t *testing.T) {
		assert.Equal(t, "|r", EscapeField("\r"))
	})

	t.Run("vertical bar", func(t *testing.T) {
		assert.Equal(t, "||", EscapeField("|"))
	})

	t.Run("opening bracket", func(t *testing.T) {
		assert.Equal(t, "|[", EscapeField("["))
	})

	t.Run("closing bracket", func(t *testing.T) {
		assert.Equal(t, "|]", EscapeField("]"))
	})

	t.Run("nothing", func(t *testing.T) {
		assert.Equal(t, "hello world", EscapeField("hello world"))
	})

	t.Run("middle", func(t *testing.T) {
		assert.Equal(t, "hello|nworld", EscapeField("hello\nworld"))
	})

	t.Run("emoji", func(t *testing.T) {
		t.SkipNow()
		assert.Equal(t, "|0x1f60a", EscapeField("\u1f60a"))
	})
}
