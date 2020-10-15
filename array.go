package goo

import "reflect"

type Array interface {
	Type
	Instantiable
	GetElementType() Type
	GetLength() int
}

type arrayType struct {
	baseType
	elementType Type
	length      int
}

func newArrayType(baseTyp baseType) Array {
	return arrayType{
		baseTyp,
		GetTypeFromGoType(baseTyp.GetGoType().Elem()),
		baseTyp.GetGoType().Len(),
	}
}

func (array arrayType) GetElementType() Type {
	return array.elementType
}

func (array arrayType) GetLength() int {
	return array.length
}

func (array arrayType) NewInstance() interface{} {
	return reflect.ArrayOf(array.length, array.GetGoType()).Elem()
}
