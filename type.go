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
	IsNumber() bool
	IsFunction() bool
	IsStruct() bool
	IsInterface() bool
	IsString() bool
	String() string
}

type baseType struct {
	name            string
	packageName     string
	packageFullName string
	typ             reflect.Type
	val             reflect.Value
	kind            reflect.Kind
	isNumber        bool
}

func newBaseType(typ reflect.Type, val reflect.Value) baseType {
	return baseType{
		getTypeName(typ, val),
		getPackageName(typ, val),
		getPackageFullName(typ, val),
		typ,
		val,
		typ.Kind(),
		isNumber(typ),
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

func (typ baseType) String() string {
	return typ.name
}

func GetType(obj interface{}) Type {
	typ, val := GetGoTypeAndValue(obj)
	typeFromCache := getTypeFromCache(typ)
	if typeFromCache != nil {
		return typeFromCache
	}
	baseTyp := createBaseType(typ, val)
	return putTypeIntoCache(getActualTypeFromBaseType(baseTyp))
}

func GetTypeFromGoType(typ reflect.Type) Type {
	if typ == nil {
		panic("Type cannot be nil")
	}
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	typeFromCache := getTypeFromCache(typ)
	if typeFromCache != nil {
		return typeFromCache
	}
	baseTyp := newBaseType(typ, reflect.Value{})
	return getActualTypeFromBaseType(baseTyp)
}
