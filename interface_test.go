package goo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testInterface interface {
	testMethod()
	testMethod2() int
}

func TestInterfaceType_GetMethodCount(t *testing.T) {
	typ := GetType((*testInterface)(nil))
	assert.True(t, typ.IsInterface())

	interfaceType := typ.ToInterfaceType()
	assert.Equal(t, 2, interfaceType.GetMethodCount())
}

func TestInterfaceType_GetMethods(t *testing.T) {
	typ := GetType((*testInterface)(nil))
	assert.True(t, typ.IsInterface())

	interfaceType := typ.ToInterfaceType()
	assert.Len(t, interfaceType.GetMethods(), 2)
}
