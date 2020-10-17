package goo

import "reflect"

type Struct interface {
	Type
	Instantiable
	GetFields() []Field
	GetFieldCount() int
	GetMethods() []Method
	GetMethodCount() int
	Implements(i Interface) bool
	EmbeddedStruct(candidate Struct) bool
}

type structType struct {
	baseType
}

func newStructType(baseTyp baseType) structType {
	return structType{
		baseTyp,
	}
}

func (typ structType) GetFields() []Field {
	fields := getFieldsFromCache(typ.GetFullName())
	if fields != nil {
		return fields
	}
	fields = make([]Field, 0)
	fieldCount := typ.GetFieldCount()
	for fieldIndex := 0; fieldIndex < fieldCount; fieldIndex++ {
		field := typ.typ.Field(fieldIndex)
		fields = append(fields, convertGoFieldToMemberField(field))
	}
	return putFieldsIntoCache(typ.GetFullName(), fields)
}

func (typ structType) GetFieldCount() int {
	return typ.typ.NumField()
}

func (typ structType) GetMethods() []Method {
	methods := getMethodsFromCache(typ.GetFullName())
	if methods != nil {
		return methods
	}
	methods = make([]Method, 0)
	methodCount := typ.GetMethodCount()
	for methodIndex := 0; methodIndex < methodCount; methodIndex++ {
		method := typ.typ.Method(methodIndex)
		methods = append(methods, convertGoMethodToMemberMethod(method))
	}
	return putMethodsIntoCache(typ.GetFullName(), methods)
}

func (typ structType) GetMethodCount() int {
	return typ.typ.NumMethod()
}

func (typ structType) Implements(i Interface) bool {
	return typ.GetGoType().Implements(i.GetGoType())
}

func (typ structType) NewInstance() interface{} {
	if typ.isPointer {
		return reflect.New(typ.GetGoType()).Interface()
	}
	return reflect.New(typ.GetGoType()).Elem().Interface()
}

func (typ structType) EmbeddedStruct(candidate Struct) bool {
	if candidate == nil {
		panic("candidate must not be null")
	}
	return typ.embeddedStruct(typ, candidate)
}

func (typ structType) embeddedStruct(parent Struct, candidate Struct) bool {
	fields := parent.GetFields()
	for _, field := range fields {
		if field.IsAnonymous() && field.GetType().IsStruct() {
			if candidate.Equals(field.GetType()) {
				return true
			}
			if field.GetType().(Struct).GetFieldCount() > 0 {
				return typ.embeddedStruct(field.GetType().(Struct), candidate)
			}
		}
	}
	return false
}

func (typ structType) visitEmbeddedStruct(parent Struct, embeddeds []Struct) []Struct {
	fields := parent.GetFields()
	for _, field := range fields {
		if field.IsAnonymous() && field.GetType().IsStruct() {
			embeddeds = append(embeddeds, field.GetType().(Struct))
			if field.GetType().(Struct).GetFieldCount() > 0 {
				embeddeds = typ.visitEmbeddedStruct(field.GetType().(Struct), embeddeds)
			}
		}
	}
	return embeddeds
}
