package goo

import "reflect"

type Slice interface {
	Type
	Instantiable
	GetElementType() Type
}

type sliceType struct {
	*baseType
	elementType Type
}

func newSliceType(baseTyp *baseType) Slice {
	return sliceType{
		baseTyp,
		GetTypeFromGoType(baseTyp.GetGoType().Elem()),
	}
}

func (slice sliceType) GetElementType() Type {
	return slice.elementType
}

func (slice sliceType) NewInstance() interface{} {
	return reflect.SliceOf(slice.GetGoType()).Elem()
}
