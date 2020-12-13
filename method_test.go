package goo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemberMethod_GetName(t *testing.T) {
	typ := GetType((*testInterface)(nil))
	assert.True(t, typ.IsInterface())

	interfaceType := typ.ToInterfaceType()
	assert.Len(t, interfaceType.GetMethods(), 3)

	assert.Equal(t, "TestMethod3", interfaceType.GetMethods()[0].GetName())
	assert.Equal(t, "testMethod", interfaceType.GetMethods()[1].GetName())
	assert.Equal(t, "testMethod2", interfaceType.GetMethods()[2].GetName())
}

func TestMemberMethod_GetMethodParameterCount(t *testing.T) {
	typ := GetType((*testInterface)(nil))
	assert.True(t, typ.IsInterface())

	interfaceType := typ.ToInterfaceType()
	assert.Len(t, interfaceType.GetMethods(), 3)

	assert.Equal(t, 0, interfaceType.GetMethods()[0].GetMethodParameterCount())
	assert.Equal(t, 3, interfaceType.GetMethods()[1].GetMethodParameterCount())
	assert.Equal(t, 0, interfaceType.GetMethods()[2].GetMethodParameterCount())
}

func TestMemberMethod_GetMethodParameterTypes(t *testing.T) {
	typ := GetType((*testInterface)(nil))
	assert.True(t, typ.IsInterface())

	interfaceType := typ.ToInterfaceType()
	assert.Len(t, interfaceType.GetMethods(), 3)

	assert.Equal(t, 0, len(interfaceType.GetMethods()[0].GetMethodParameterTypes()))
	assert.Equal(t, 3, len(interfaceType.GetMethods()[1].GetMethodParameterTypes()))
	assert.Equal(t, 0, len(interfaceType.GetMethods()[2].GetMethodParameterTypes()))

	types := interfaceType.GetMethods()[1].GetMethodParameterTypes()
	assert.Equal(t, "string", types[0].GetFullName())
	assert.Equal(t, "int", types[1].GetFullName())
	assert.Equal(t, "bool", types[2].GetFullName())
}

func TestMemberMethod_GetMethodReturnTypeCount(t *testing.T) {
	typ := GetType((*testInterface)(nil))
	assert.True(t, typ.IsInterface())

	interfaceType := typ.ToInterfaceType()
	assert.Len(t, interfaceType.GetMethods(), 3)

	assert.Equal(t, 0, interfaceType.GetMethods()[0].GetMethodReturnTypeCount())
	assert.Equal(t, 2, interfaceType.GetMethods()[1].GetMethodReturnTypeCount())
	assert.Equal(t, 0, interfaceType.GetMethods()[2].GetMethodReturnTypeCount())
}

func TestMemberMethod_GetMethodReturnTypes(t *testing.T) {
	typ := GetType((*testInterface)(nil))
	assert.True(t, typ.IsInterface())

	interfaceType := typ.ToInterfaceType()
	assert.Len(t, interfaceType.GetMethods(), 3)

	assert.Equal(t, 0, len(interfaceType.GetMethods()[0].GetMethodReturnTypes()))
	assert.Equal(t, 2, len(interfaceType.GetMethods()[1].GetMethodReturnTypes()))
	assert.Equal(t, 0, len(interfaceType.GetMethods()[2].GetMethodReturnTypes()))

	types := interfaceType.GetMethods()[1].GetMethodReturnTypes()
	assert.Equal(t, "string", types[0].GetFullName())
	assert.Equal(t, "error", types[1].GetFullName())
}

func TestMemberMethod_Invoke(t *testing.T) {
	typ := GetType(Dog{})
	structType := typ.ToStructType()
	methods := structType.GetStructMethods()

	assert.NotPanics(t, func() {
		methods[0].Invoke(Dog{})
		methods[2].Invoke(Dog{}, nil, nil)
		methods[2].Invoke(Dog{}, "test", nil)
		outputs := methods[3].Invoke(Dog{})
		assert.Len(t, outputs, 1)
	})

	assert.Panics(t, func() {
		methods[0].Invoke(2)
	})

	assert.Panics(t, func() {
		methods[0].Invoke(Dog{}, "arg1", "arg2")
	})

	assert.Panics(t, func() {
		methods[0].Invoke(Product{})
	})
}

func TestMemberMethod_IsExported(t *testing.T) {
	typ := GetType((*testInterface)(nil))
	assert.True(t, typ.IsInterface())

	interfaceType := typ.ToInterfaceType()
	assert.Len(t, interfaceType.GetMethods(), 3)

	assert.True(t, interfaceType.GetMethods()[0].IsExported())
	assert.False(t, interfaceType.GetMethods()[1].IsExported())
	assert.False(t, interfaceType.GetMethods()[2].IsExported())
}

func TestMemberMethod_String(t *testing.T) {
	typ := GetType((*testInterface)(nil))
	assert.True(t, typ.IsInterface())

	interfaceType := typ.ToInterfaceType()
	assert.Len(t, interfaceType.GetMethods(), 3)

	assert.Equal(t, "TestMethod3", interfaceType.GetMethods()[0].String())
	assert.Equal(t, "testMethod", interfaceType.GetMethods()[1].String())
	assert.Equal(t, "testMethod2", interfaceType.GetMethods()[2].String())
}
