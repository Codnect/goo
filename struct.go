package goo

import "reflect"

type Struct interface {
	Type
	Instantiable
	GetFields() []Field
	GetFieldCount() int
	GetExportedFields() []Field
	GetExportedFieldCount() int
	GetUnexportedFields() []Field
	GetUnexportedFieldCount() int
	GetAnonymousFields() []Field
	GetAnonymousFieldCount() int
	GetStructMethods() []Method
	GetStructMethodCount() int
	Implements(i Interface) bool
	EmbeddedStruct(candidate Struct) bool
}

type structType struct {
	*baseType
}

func newStructType(baseTyp *baseType) structType {
	return structType{
		baseTyp,
	}
}

func (typ structType) GetFields() []Field {
	fields := make([]Field, 0)
	fieldCount := typ.GetFieldCount()
	for fieldIndex := 0; fieldIndex < fieldCount; fieldIndex++ {
		field := typ.typ.Field(fieldIndex)
		fields = append(fields, convertGoFieldToMemberField(field))
	}
	return fields
}

func (typ structType) GetFieldCount() int {
	return typ.typ.NumField()
}

func (typ structType) GetExportedFields() []Field {
	exportedFields := make([]Field, 0)
	fields := typ.GetFields()
	for _, field := range fields {
		if field.IsExported() {
			exportedFields = append(exportedFields, field)
		}
	}
	return exportedFields
}

func (typ structType) GetExportedFieldCount() int {
	exportedFieldsCount := 0
	fields := typ.GetFields()
	for _, field := range fields {
		if field.IsExported() {
			exportedFieldsCount++
		}
	}
	return exportedFieldsCount
}

func (typ structType) GetUnexportedFields() []Field {
	unExportedFields := make([]Field, 0)
	fields := typ.GetFields()
	for _, field := range fields {
		if !field.IsExported() {
			unExportedFields = append(unExportedFields, field)
		}
	}
	return unExportedFields
}

func (typ structType) GetUnexportedFieldCount() int {
	unExportedFieldsCount := 0
	fields := typ.GetFields()
	for _, field := range fields {
		if !field.IsExported() {
			unExportedFieldsCount++
		}
	}
	return unExportedFieldsCount
}

func (typ structType) GetAnonymousFields() []Field {
	anonymousFields := make([]Field, 0)
	fields := typ.GetFields()
	for _, field := range fields {
		if field.IsAnonymous() {
			anonymousFields = append(anonymousFields, field)
		}
	}
	return anonymousFields
}

func (typ structType) GetAnonymousFieldCount() int {
	anonymousFieldCount := 0
	fields := typ.GetFields()
	for _, field := range fields {
		if field.IsAnonymous() {
			anonymousFieldCount++
		}
	}
	return anonymousFieldCount
}

func (typ structType) GetStructMethods() []Method {
	methods := make([]Method, 0)
	methodCount := typ.GetStructMethodCount()
	var method reflect.Method
	for methodIndex := 0; methodIndex < methodCount; methodIndex++ {
		if typ.isPointer {
			method = typ.ptrType.Method(methodIndex)
		} else {
			method = typ.typ.Method(methodIndex)
		}
		methods = append(methods, convertGoMethodToMemberMethod(method))
	}
	return methods
}

func (typ structType) GetStructMethodCount() int {
	if typ.isPointer {
		return typ.ptrType.NumMethod()
	}
	return typ.typ.NumMethod()
}

func (typ structType) Implements(i Interface) bool {
	if typ.isPointer {
		return typ.GetGoPointerType().Implements(i.GetGoType())
	}
	return typ.GetGoType().Implements(i.GetGoType())
}

func (typ structType) NewInstance() interface{} {
	return reflect.New(typ.GetGoType()).Interface()
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
