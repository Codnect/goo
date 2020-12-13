package goo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapType_NewInstance(t *testing.T) {
	m := make(map[string]bool)
	typ := GetType(m)
	assert.True(t, typ.IsMap())

	mapType := typ.ToMapType()
	newMapInstance := mapType.NewInstance().(map[string]bool)
	assert.NotNil(t, newMapInstance)
}

func TestMapType_GetKeyType(t *testing.T) {
	m := make(map[string]bool)
	typ := GetType(m)
	assert.True(t, typ.IsMap())

	mapType := typ.ToMapType()
	assert.Equal(t, "string", mapType.GetKeyType().GetFullName())
}

func TestMapType_GetValueType(t *testing.T) {
	m := make(map[string]bool)
	typ := GetType(m)
	assert.True(t, typ.IsMap())

	mapType := typ.ToMapType()
	assert.Equal(t, "bool", mapType.GetValueType().GetFullName())
}
