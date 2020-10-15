package goo

type Map interface {
	Type
	GetKeyType() Type
	GetValueType() Type
}

type baseMap struct {
	baseType
	keyType   Type
	valueType Type
}

func newMap(baseTyp baseType) Map {
	return baseMap{
		baseTyp,
		GetTypeFromGoType(baseTyp.GetGoType().Key()),
		GetTypeFromGoType(baseTyp.GetGoType().Elem()),
	}
}

func (m baseMap) GetKeyType() Type {
	return m.keyType
}

func (m baseMap) GetValueType() Type {
	return m.valueType
}
