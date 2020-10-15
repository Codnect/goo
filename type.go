package goo

import (
	"reflect"
)

type Type interface {
	GetName() string
	GetFullName() string
	GetPackageName() string
	GetPackageFullName() string
	GetGoType() reflect.Type
	GetGoValue() reflect.Value
	IsBoolean() bool
	IsNumber() bool
	IsFunction() bool
	IsStruct() bool
	IsInterface() bool
	IsString() bool
	IsMap() bool
	IsPointer() bool
	String() string
	Equals(anotherType Type) bool
}

type baseType struct {
	name            string
	packageName     string
	packageFullName string
	typ             reflect.Type
	val             reflect.Value
	kind            reflect.Kind
	isNumber        bool
	isPointer       bool
}

func newBaseType(typ reflect.Type, val reflect.Value, isPointer bool) baseType {
	return baseType{
		getTypeName(typ, val),
		getPackageName(typ, val),
		getPackageFullName(typ, val),
		typ,
		val,
		typ.Kind(),
		isNumber(typ),
		isPointer,
	}
}

func (typ baseType) GetName() string {
	return typ.name
}

func (typ baseType) GetFullName() string {
	return typ.packageFullName + "." + typ.name
}

func (typ baseType) GetPackageName() string {
	return typ.packageName
}

func (typ baseType) GetPackageFullName() string {
	return typ.packageFullName
}

func (typ baseType) GetGoType() reflect.Type {
	return typ.typ
}

func (typ baseType) GetGoValue() reflect.Value {
	return typ.val
}

func (typ baseType) IsBoolean() bool {
	return reflect.Bool == typ.kind
}

func (typ baseType) IsNumber() bool {
	return typ.isNumber
}

func (typ baseType) IsFunction() bool {
	return reflect.Func == typ.kind
}

func (typ baseType) IsStruct() bool {
	return reflect.Struct == typ.kind
}

func (typ baseType) IsInterface() bool {
	return reflect.Interface == typ.kind
}

func (typ baseType) IsString() bool {
	return reflect.String == typ.kind
}

func (typ baseType) IsMap() bool {
	return reflect.Map == typ.kind
}

func (typ baseType) IsPointer() bool {
	return typ.isPointer
}

func (typ baseType) String() string {
	return typ.name
}

func (typ baseType) Equals(anotherType Type) bool {
	if anotherType == nil {
		return false
	}
	return typ.typ == anotherType.GetGoType()
}

func GetType(obj interface{}) Type {
	typ, val, isPointer := GetGoTypeAndValue(obj)
	typeFromCache := getTypeFromCache(typ)
	if typeFromCache != nil {
		return typeFromCache
	}
	baseTyp := createBaseType(typ, val, isPointer)
	return putTypeIntoCache(getActualTypeFromBaseType(baseTyp))
}

func GetTypeFromGoType(typ reflect.Type) Type {
	if typ == nil {
		panic("Type cannot be nil")
	}
	isPointer := false
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		isPointer = true
	}
	typeFromCache := getTypeFromCache(typ)
	if typeFromCache != nil {
		return typeFromCache
	}
	baseTyp := newBaseType(typ, reflect.Value{}, isPointer)
	return getActualTypeFromBaseType(baseTyp)
}
