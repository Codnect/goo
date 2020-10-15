package goo

type Slice interface {
	Type
	GetElementType() Type
}

type sliceType struct {
	baseType
	elementType Type
}

func newSliceType(baseTyp baseType) Slice {
	return sliceType{
		baseTyp,
		GetTypeFromGoType(baseTyp.GetGoType().Elem()),
	}
}

func (slice sliceType) GetElementType() Type {
	return slice.elementType
}
