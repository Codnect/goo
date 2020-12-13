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
	*baseType
}

func newFunctionType(baseTyp *baseType) functionType {
	return functionType{
		baseTyp,
	}
}

func (fun functionType) GetFunctionParameterTypes() []Type {
	parameterTypes := make([]Type, 0)
	parameterCount := fun.GetFunctionParameterCount()
	for paramIndex := 0; paramIndex < parameterCount; paramIndex++ {
		paramTyp := fun.typ.In(paramIndex)
		parameterTypes = append(parameterTypes, getTypeFromGoType(paramTyp))
	}
	return parameterTypes
}

func (fun functionType) GetFunctionParameterCount() int {
	return fun.typ.NumIn()
}

func (fun functionType) GetFunctionReturnTypes() []Type {
	returnTypes := make([]Type, 0)
	returnTypeCount := fun.GetFunctionReturnTypeCount()
	for returnTypeIndex := 0; returnTypeIndex < returnTypeCount; returnTypeIndex++ {
		returnType := fun.typ.Out(returnTypeIndex)
		returnTypes = append(returnTypes, getTypeFromGoType(returnType))
	}
	return returnTypes
}

func (fun functionType) GetFunctionReturnTypeCount() int {
	return fun.typ.NumOut()
}

func (fun functionType) Call(args []interface{}) []interface{} {
	parameterCount := fun.GetFunctionParameterCount()
	if (args == nil && parameterCount != 0) || (args != nil && len(args) != parameterCount) {
		panic("Parameter counts don't match argument counts")
	}
	inputs := make([]reflect.Value, 0)
	for index, arg := range args {
		if arg == nil {
			paramType := fun.GetFunctionParameterTypes()[index]
			if paramType.IsPointer() {
				inputs = append(inputs, reflect.New(paramType.GetGoPointerType()).Elem())
			} else {
				inputs = append(inputs, reflect.New(paramType.GetGoType()).Elem())
			}
		} else {
			inputs = append(inputs, reflect.ValueOf(arg))
		}
	}
	outputs := make([]interface{}, 0)
	results := fun.val.Call(inputs)
	for _, outputParam := range results {
		outputs = append(outputs, outputParam.Interface())
	}
	return outputs
}
