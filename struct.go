package goo

type Struct interface {
	Type
	GetFields() []Field
	GetFieldCount() int
	GetMethods() []Method
	GetMethodCount() int
	GetInterfaces() []Interface
	GetInterfaceCount() int
	GetEmbeddeds() []Type
	GetEmbeddedCount() int
	GetEmbeddedInterfaces() []Interface
	GetEmbeddedInterfaceCount() int
	GetEmbeddedStructs() []Struct
	GetEmbeddedStructCount() int
	Implements(i Interface) bool
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

func (typ structType) GetInterfaces() []Interface {
	return nil
}

func (typ structType) GetInterfaceCount() int {
	return 0
}

func (typ structType) GetEmbeddeds() []Type {
	return nil
}

func (typ structType) GetEmbeddedCount() int {
	return 0
}

func (typ structType) GetEmbeddedInterfaces() []Interface {
	return nil
}

func (typ structType) GetEmbeddedInterfaceCount() int {
	return 0
}

func (typ structType) GetEmbeddedStructs() []Struct {
	return nil
}

func (typ StructType) GetEmbeddedStructCount() int {
	return 0
}

func (typ StructType) Implements(i Interface) bool {
	return typ.GetGoType().Implements(i.GetGoType())
}
