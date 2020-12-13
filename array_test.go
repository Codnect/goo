package goo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayType_GetElementType(t *testing.T) {
	arr := [5]string{}
	typ := GetType(arr)
	assert.True(t, typ.IsArray())

	arrayType := typ.ToArrayType()
	assert.Equal(t, "string", arrayType.GetElementType().GetFullName())
}

func TestArrayType_GetLength(t *testing.T) {
	arr := [5]string{}
	typ := GetType(arr)
	assert.True(t, typ.IsArray())

	arrayType := typ.ToArrayType()
	assert.Equal(t, 5, arrayType.GetLength())
}

func TestArrayType_NewInstance(t *testing.T) {
	arr := [5]string{}
	typ := GetType(arr)
	assert.True(t, typ.IsArray())

	arrayType := typ.ToArrayType()
	instance := arrayType.NewInstance().(*[5]string)
	assert.NotNil(t, instance)
	assert.Len(t, *instance, 5)
}
