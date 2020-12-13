package goo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStructType_GetNames(t *testing.T) {
	testGetNamesForStruct(t, GetType(Animal{}))
	testGetNamesForStruct(t, GetType(&Animal{}))
}

func testGetNamesForStruct(t *testing.T, typ Type) {
	assert.Equal(t, "Animal", typ.GetName())
	assert.Equal(t, "github.com.procyon.projects.goo.Animal", typ.GetFullName())
	assert.Equal(t, "goo", typ.GetPackageName())
	assert.Equal(t, "github.com.procyon.projects.goo", typ.GetPackageFullName())
	assert.Equal(t, typ.(Struct), typ.ToStructType())
}

func TestStructType_GetFields(t *testing.T) {
	testGetFieldsForStruct(t, GetType(Person{}))
	testGetFieldsForStruct(t, GetType(&Person{}))
}

func testGetFieldsForStruct(t *testing.T, typ Type) {
	structType := typ.(Struct)
	// all fields
	fieldCount := structType.GetFieldCount()
	assert.Equal(t, 5, fieldCount)
	fields := structType.GetFields()
	assert.Equal(t, 5, len(fields))

	// exported fields
	fieldCount = structType.GetExportedFieldCount()
	assert.Equal(t, 3, fieldCount)
	fields = structType.GetExportedFields()
	assert.Equal(t, 3, len(fields))

	// unexported fields
	fieldCount = structType.GetUnexportedFieldCount()
	assert.Equal(t, 2, fieldCount)
	fields = structType.GetUnexportedFields()
	assert.Equal(t, 2, len(fields))

	// anonymous fields
	fieldCount = structType.GetAnonymousFieldCount()
	assert.Equal(t, 1, fieldCount)
	fields = structType.GetAnonymousFields()
	assert.Equal(t, 1, len(fields))
}

func TestStructType_GetStructMethods(t *testing.T) {
	typ := GetType(Person{})
	structType := typ.(Struct)
	methodsCount := structType.GetStructMethodCount()
	assert.Equal(t, 2, methodsCount)
	methods := structType.GetStructMethods()
	assert.Equal(t, 2, len(methods))

	typ = GetType(&Person{})
	structType = typ.(Struct)
	methodsCount = structType.GetStructMethodCount()
	assert.Equal(t, 3, methodsCount)
	methods = structType.GetStructMethods()
	assert.Equal(t, 3, len(methods))
}

func TestStructType_Implements(t *testing.T) {
	x := &Dog{}
	x.Run()
	typ := GetType(Dog{})
	structType := typ.(Struct)
	assert.Equal(t, false, structType.Implements(GetType((*Bark)(nil)).(Interface)))
	assert.Equal(t, true, structType.Implements(GetType((*Run)(nil)).(Interface)))

	typ = GetType(&Dog{})
	structType = typ.(Struct)
	assert.Equal(t, true, structType.Implements(GetType((*Bark)(nil)).(Interface)))
	assert.Equal(t, true, structType.Implements(GetType((*Run)(nil)).(Interface)))
}

func TestStructType_EmbeddedStruct(t *testing.T) {
	typ := GetType(Dog{})
	assert.True(t, typ.ToStructType().EmbeddedStruct(GetType(Animal{}).ToStructType()))
	assert.True(t, typ.ToStructType().EmbeddedStruct(GetType(Cell{}).ToStructType()))
	assert.False(t, typ.ToStructType().EmbeddedStruct(GetType(Dog{}).ToStructType()))
	assert.Panics(t, func() {
		typ.ToStructType().EmbeddedStruct(nil)
	})
}

func TestStructType_NewInstance(t *testing.T) {
	typ := GetType(Dog{})
	instance := typ.ToStructType().NewInstance()
	assert.NotNil(t, instance)
}
