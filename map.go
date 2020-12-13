package goo

import "reflect"

type Map interface {
	Type
	Instantiable
	GetKeyType() Type
	GetValueType() Type
}

type mapType struct {
	*baseType
	keyType   Type
	valueType Type
}

func newMapType(baseTyp *baseType) Map {
	return mapType{
		baseTyp,
		getTypeFromGoType(baseTyp.GetGoType().Key()),
		getTypeFromGoType(baseTyp.GetGoType().Elem()),
	}
}

func (m mapType) GetKeyType() Type {
	return m.keyType
}

func (m mapType) GetValueType() Type {
	return m.valueType
}

func (m mapType) NewInstance() interface{} {
	mapType := reflect.MapOf(m.keyType.GetGoType(), m.valueType.GetGoType())
	instance := reflect.MakeMapWithSize(mapType, 0)
	return instance.Interface()
}
