package goo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceType_GetElementType(t *testing.T) {
	var slice []string
	typ := GetType(slice)
	assert.True(t, typ.IsSlice())

	sliceType := typ.ToSliceType()
	assert.Equal(t, "string", sliceType.GetElementType().GetFullName())
}

func TestSliceType_NewInstance(t *testing.T) {
	arr := [5]string{}
	typ := GetType(arr[2:3])
	assert.True(t, typ.IsSlice())

	sliceType := typ.ToSliceType()
	instance := sliceType.NewInstance().([]string)
	assert.NotNil(t, instance)
}
