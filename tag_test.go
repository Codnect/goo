package goo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTag_String(t *testing.T) {
	tag := &Tag{
		Name:  "test-tag",
		Value: "test-value",
	}
	assert.Equal(t, tag.Name+"->"+tag.Value, tag.String())
}
