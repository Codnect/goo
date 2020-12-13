package goo

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func testFunction(name string, i int, val bool, test *string) (string, error) {
	return "test", errors.New("test error")
}

func TestFunctionType_GetFunctionParameterCount(t *testing.T) {
	typ := GetType(testFunction)
	assert.True(t, typ.IsFunction())

	functionType := typ.ToFunctionType()
	assert.Equal(t, 4, functionType.GetFunctionParameterCount())
}

func TestFunctionType_GetFunctionParameterTypes(t *testing.T) {
	typ := GetType(testFunction)
	assert.True(t, typ.IsFunction())

	functionType := typ.ToFunctionType()
	parameterTypes := functionType.GetFunctionParameterTypes()
	assert.Equal(t, GetType("").GetGoType(), parameterTypes[0].GetGoType())
	assert.Equal(t, GetType(0).GetGoType(), parameterTypes[1].GetGoType())
	assert.Equal(t, GetType(true).GetGoType(), parameterTypes[2].GetGoType())
}

func TestFunctionType_GetFunctionReturnTypeCount(t *testing.T) {
	typ := GetType(testFunction)
	assert.True(t, typ.IsFunction())

	functionType := typ.ToFunctionType()
	assert.Equal(t, 2, functionType.GetFunctionReturnTypeCount())
}

func TestFunctionType_GetFunctionReturnTypes(t *testing.T) {
	typ := GetType(testFunction)
	assert.True(t, typ.IsFunction())

	functionType := typ.ToFunctionType()
	parameterTypes := functionType.GetFunctionReturnTypes()
	assert.Equal(t, GetType("").GetGoType(), parameterTypes[0].GetGoType())
	assert.Equal(t, "error", parameterTypes[1].GetFullName())
}

func TestFunctionType_Call(t *testing.T) {
	typ := GetType(testFunction)
	assert.True(t, typ.IsFunction())
	functionType := typ.ToFunctionType()

	args := make([]interface{}, 0)
	args = append(args, "test")
	args = append(args, 1)
	args = append(args, nil)
	args = append(args, nil)

	outputs := functionType.Call(args)
	assert.Len(t, outputs, 2)

	assert.Equal(t, "test", outputs[0])
	assert.Equal(t, "test error", outputs[1].(error).Error())
}

func TestFunctionType_Call_WithMissingParameters(t *testing.T) {
	typ := GetType(testFunction)
	assert.True(t, typ.IsFunction())
	functionType := typ.ToFunctionType()

	args := make([]interface{}, 0)
	args = append(args, "test")
	args = append(args, 1)

	assert.Panics(t, func() {
		functionType.Call(args)
	})
}
