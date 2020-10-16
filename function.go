package goo

import "reflect"

type Function interface {
	Type
	GetFunctionParameterTypes() []Type
	GetFunctionParameterCount() int
	GetFunctionReturnTypes() []Type
	GetFunctionReturnTypeCount() int
	Call(args []interface{}) []interface{}
}

type functionType struct {
	baseType
}

func newFunctionType(baseTyp baseType) functionType {
	return functionType{
		baseTyp,
	}
}

func (fun functionType) GetFunctionParameterTypes() []Type {
	parameterTypes := make([]Type, 0)
	parameterCount := fun.GetFunctionParameterCount()
	for paramIndex := 0; paramIndex < parameterCount; paramIndex++ {
		paramTyp := fun.typ.In(paramIndex)
		parameterTypes = append(parameterTypes, GetTypeFromGoType(paramTyp))
	}
	return parameterTypes
}

func (fun functionType) GetFunctionParameterCount() int {
	return fun.typ.NumIn()
}

func (fun functionType) GetFunctionReturnTypes() []Type {
	returnTypes := make([]Type, 0)
	returnTypeCount := fun.GetFunctionParameterCount()
	for returnTypeIndex := 0; returnTypeIndex < returnTypeCount; returnTypeIndex++ {
		returnType := fun.typ.Out(returnTypeIndex)
		returnTypes = append(returnTypes, GetTypeFromGoType(returnType))
	}
	return returnTypes
}

func (fun functionType) GetFunctionReturnTypeCount() int {
	return fun.typ.NumOut()
}

func (fun functionType) Call(args []interface{}) []interface{} {
	inputs := make([]reflect.Value, 0)
	for _, arg := range args {
		inputs = append(inputs, reflect.ValueOf(arg))
	}
	outputs := make([]interface{}, 0)
	results := fun.val.Call(inputs)
	for _, outputParam := range results {
		outputs = append(outputs, outputParam.Interface())
	}
	return outputs
}
