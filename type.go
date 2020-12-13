package goo

import (
	"reflect"
)

type TypeConverter interface {
	IsBoolean() bool
	IsNumber() bool
	IsFunction() bool
	IsStruct() bool
	IsInterface() bool
	IsString() bool
	IsMap() bool
	IsArray() bool
	IsSlice() bool
	ToBooleanType() Boolean
	ToNumberType() Number
	ToFunctionType() Function
	ToStructType() Struct
	ToInterfaceType() Interface
	ToStringType() String
	ToMapType() Map
	ToArrayType() Array
	ToSliceType() Slice
}

type Type interface {
	TypeConverter
	GetName() string
	GetFullName() string
	GetPackageName() string
	GetPackageFullName() string
	GetGoPointerType() reflect.Type
	GetGoPointerValue() reflect.Value
	GetGoType() reflect.Type
	GetGoValue() reflect.Value
	IsPointer() bool
	IsInstantiable() bool
	String() string
	Equals(anotherType Type) bool
}

type baseType struct {
	parentType      interface{}
	name            string
	packageName     string
	packageFullName string
	typ             reflect.Type
	val             reflect.Value
	ptrType         reflect.Type
	ptrVal          reflect.Value
	kind            reflect.Kind
	isNumber        bool
	isPointer       bool
}

func newBaseType(typ reflect.Type, val reflect.Value) *baseType {
	return &baseType{
		nil,
		getTypeName(typ, val),
		getPackageName(typ, val),
		getPackageFullName(typ, val),
		typ,
		val,
		nil,
		reflect.Value{},
		typ.Kind(),
		isNumber(typ),
		false,
	}
}

func (typ baseType) GetName() string {
	return typ.name
}

func (typ baseType) GetFullName() string {
	if typ.packageFullName == "" {
		return typ.name
	}
	return typ.packageFullName + "." + typ.name
}

func (typ baseType) GetPackageName() string {
	return typ.packageName
}

func (typ baseType) GetPackageFullName() string {
	return typ.packageFullName
}

func (typ baseType) GetGoPointerType() reflect.Type {
	return typ.ptrType
}

func (typ baseType) GetGoPointerValue() reflect.Value {
	return typ.ptrVal
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

func (typ baseType) IsArray() bool {
	return reflect.Array == typ.kind
}

func (typ baseType) IsSlice() bool {
	return reflect.Slice == typ.kind
}

func (typ baseType) IsPointer() bool {
	return typ.isPointer
}

func (typ baseType) IsInstantiable() bool {
	if typ.IsInterface() || typ.IsFunction() {
		return false
	}
	return true
}

func (typ baseType) ToBooleanType() Boolean {
	return typ.parentType.(Boolean)
}

func (typ baseType) ToNumberType() Number {
	return typ.parentType.(Number)
}

func (typ baseType) ToFunctionType() Function {
	return typ.parentType.(Function)
}

func (typ baseType) ToInterfaceType() Interface {
	return typ.parentType.(Interface)
}

func (typ baseType) ToStringType() String {
	return typ.parentType.(String)
}

func (typ baseType) ToMapType() Map {
	return typ.parentType.(Map)
}

func (typ baseType) ToArrayType() Array {
	return typ.parentType.(Array)
}

func (typ baseType) ToSliceType() Slice {
	return typ.parentType.(Slice)
}

func (typ baseType) ToStructType() Struct {
	return typ.parentType.(Struct)
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
	typ, val, isPointer := getGoTypeAndValue(obj)
	baseTyp := createBaseType(typ, val)
	populatePointerInfo(obj, baseTyp, isPointer)
	actualType := getActualTypeFromBaseType(baseTyp)
	baseTyp.parentType = actualType
	return actualType
}

func populatePointerInfo(obj interface{}, baseType *baseType, isPointer bool) {
	if isPointer {
		typ, val := getGoPointerTypeAndValue(obj)
		baseType.isPointer = true
		baseType.ptrType = typ
		baseType.ptrVal = val
	}
}

func getTypeFromGoType(typ reflect.Type) Type {
	if typ == nil {
		panic("Type cannot be nil")
	}
	isPointer := false
	var ptrType reflect.Type
	if typ.Kind() == reflect.Ptr {
		ptrType = typ
		typ = typ.Elem()
		isPointer = true
	}
	baseTyp := newBaseType(typ, reflect.Value{})
	actualType := getActualTypeFromBaseType(baseTyp)
	baseTyp.parentType = actualType
	if isPointer {
		baseTyp.ptrType = ptrType
		baseTyp.isPointer = true
	}
	return actualType
}
