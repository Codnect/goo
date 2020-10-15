package goo

type Array interface {
	Type
	GetElementType() Type
}

type arrayType struct {
	baseType
	elementType Type
}

func newArrayType(baseTyp baseType) Array {
	return arrayType{
		baseTyp,
		GetTypeFromGoType(baseTyp.GetGoType().Elem()),
	}
}

func (array arrayType) GetElementType() Type {
	return array.elementType
}
