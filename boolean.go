package goo

type Boolean interface {
	Type
	ToBoolean(value string) bool
	ToString(value bool) string
}

type BooleanType struct {
	baseType
}

func newBooleanType(baseTyp baseType) Boolean {
	return BooleanType{
		baseTyp,
	}
}

func (b BooleanType) ToBoolean(value string) bool {
	if value == "true" {
		return true
	} else if value == "false" {
		return false
	}
	panic("Given value is not true or false")
}

func (b BooleanType) ToString(value bool) string {
	if value {
		return "true"
	} else {
		return "false"
	}
}
