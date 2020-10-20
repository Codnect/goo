package goo

type Interface interface {
	Type
	GetMethods() []Method
	GetMethodCount() int
}

type interfaceType struct {
	*baseType
}

func newInterfaceType(baseTyp *baseType) interfaceType {
	return interfaceType{
		baseTyp,
	}
}

func (typ interfaceType) GetMethods() []Method {
	methods := make([]Method, 0)
	methodCount := typ.GetMethodCount()
	for methodIndex := 0; methodIndex < methodCount; methodIndex++ {
		method := typ.typ.Method(methodIndex)
		methods = append(methods, convertGoMethodToMemberMethod(method))
	}
	return methods
}

func (typ interfaceType) GetMethodCount() int {
	return typ.typ.NumMethod()
}
