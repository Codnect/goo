package goo

import "reflect"

type Boolean interface {
	Type
	Instantiable
	ToBoolean(value string) bool
	ToString(value bool) string
}

type booleanType struct {
	*baseType
}

func newBooleanType(baseTyp *baseType) Boolean {
	return booleanType{
		baseTyp,
	}
}

func (b booleanType) ToBoolean(value string) bool {
	if value == "true" {
		return true
	} else if value == "false" {
		return false
	}
	panic("Given value is not true or false")
}

func (b booleanType) ToString(value bool) string {
	if value {
		return "true"
	} else {
		return "false"
	}
}

func (b booleanType) NewInstance() interface{} {
	return reflect.New(b.GetGoType()).Interface()
}
