package goo

type Function interface {
	Type
	GetFunctionParameterTypes() []Type
	GetFunctionParameterCount() int
	GetFunctionReturnTypes() []Type
	GetFunctionReturnTypeCount() int
}

type FunctionType struct {
	baseType
}

func newFunctionType(baseTyp baseType) FunctionType {
	return FunctionType{
		baseTyp,
	}
}

func (fun FunctionType) GetFunctionParameterTypes() []Type {
	parameterTypes := make([]Type, 0)
	parameterCount := fun.GetFunctionParameterCount()
	for paramIndex := 0; paramIndex < parameterCount; paramIndex++ {
		paramTyp := fun.typ.In(paramIndex)
		parameterTypes = append(parameterTypes, GetTypeFromGoType(paramTyp))
	}
	return parameterTypes
}

func (fun FunctionType) GetFunctionParameterCount() int {
	return fun.typ.NumIn()
}

func (fun FunctionType) GetFunctionReturnTypes() []Type {
	returnTypes := make([]Type, 0)
	returnTypeCount := fun.GetFunctionParameterCount()
	for returnTypeIndex := 0; returnTypeIndex < returnTypeCount; returnTypeIndex++ {
		returnType := fun.typ.Out(returnTypeIndex)
		returnTypes = append(returnTypes, GetTypeFromGoType(returnType))
	}
	return returnTypes
}

func (fun FunctionType) GetFunctionReturnTypeCount() int {
	return fun.typ.NumOut()
}
