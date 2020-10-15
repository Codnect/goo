package goo

type Function interface {
	Type
	GetFunctionParameterTypes() []Type
	GetFunctionParameterCount() int
	GetFunctionReturnTypes() []Type
	GetFunctionReturnTypeCount() int
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
