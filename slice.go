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
		getTypeFromGoType(baseTyp.GetGoType().Elem()),
	}
}

func (slice sliceType) GetElementType() Type {
	return slice.elementType
}

func (slice sliceType) NewInstance() interface{} {
	instance := reflect.MakeSlice(slice.GetGoType(), slice.val.Len(), slice.val.Cap()).Interface()
	return instance
}
