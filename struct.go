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
}

type StructType struct {
	baseType
}

func newStructType(baseTyp baseType) StructType {
	return StructType{
		baseTyp,
	}
}

func (typ StructType) GetFields() []Field {
	fields := make([]Field, 0)
	fieldCount := typ.GetFieldCount()
	for fieldIndex := 0; fieldIndex < fieldCount; fieldIndex++ {
		field := typ.typ.Field(fieldIndex)
		fields = append(fields, convertGoFieldToMemberField(field))
	}
	return fields
}

func (typ StructType) GetFieldCount() int {
	return typ.typ.NumField()
}

func (typ StructType) GetMethods() []Method {
	methods := make([]Method, 0)
	methodCount := typ.GetMethodCount()
	for methodIndex := 0; methodIndex < methodCount; methodIndex++ {
		method := typ.typ.Method(methodIndex)
		methods = append(methods, convertGoMethodToMemberMethod(method))
	}
	return putMethodsIntoCache(typ.GetFullName(), methods)
}

func (typ StructType) GetMethodCount() int {
	return typ.typ.NumMethod()
}

func (typ StructType) GetInterfaces() []Interface {
	return nil
}

func (typ StructType) GetInterfaceCount() int {
	return 0
}

func (typ StructType) GetEmbeddeds() []Type {
	return nil
}

func (typ StructType) GetEmbeddedCount() int {
	return 0
}

func (typ StructType) GetEmbeddedInterfaces() []Interface {
	return nil
}

func (typ StructType) GetEmbeddedInterfaceCount() int {
	return 0
}

func (typ StructType) GetEmbeddedStructs() []Struct {
	return nil
}

func (typ StructType) GetEmbeddedStructCount() int {
	return 0
}
