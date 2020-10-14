package goo

type Interface interface {
	Type
	GetMethods() []Method
	GetMethodCount() int
}

type InterfaceType struct {
	baseType
}

func newInterfaceType(baseTyp baseType) InterfaceType {
	return InterfaceType{
		baseTyp,
	}
}

func (typ InterfaceType) GetMethods() []Method {
	methods := make([]Method, 0)
	methodCount := typ.GetMethodCount()
	for methodIndex := 0; methodIndex < methodCount; methodIndex++ {
		method := typ.typ.Method(methodIndex)
		methods = append(methods, convertGoMethodToMemberMethod(method))
	}
	return methods
}

func (typ InterfaceType) GetMethodCount() int {
	return typ.typ.NumMethod()
}
