package goo

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type Run interface {
	Run()
}

type Bark interface {
	Bark()
}

type Animal struct {
	Name string
}

func (animal Animal) SayHi() string {
	return "Hi, I'm " + animal.Name
}

type Dog struct {
	Animal
}

func (dog *Dog) Bark() {
	log.Print("Bark")
}

func (dog Dog) Run() {
	log.Print("Run")
}

type Person struct {
	Name    string
	Surname string
	age     int
	address Address
}

func (person Person) GetName() string {
	return person.Name
}

func (person Person) GetSurname() string {
	return person.Surname
}

func (person Person) getAge() int {
	return person.age
}

func (person *Person) GetAddress() Address {
	return person.address
}

type Address struct {
	city    string
	country string
}

func (address Address) GetCity() string {
	return address.city
}

func (address Address) GetCountry() string {
	return address.country
}

func TestIsMethods(t *testing.T) {
	typ := GetType(Animal{})
	assert.Equal(t, true, typ.IsStruct())
	assert.Equal(t, false, typ.IsInterface())
	assert.Equal(t, false, typ.IsFunction())
	assert.Equal(t, false, typ.IsNumber())
	assert.Equal(t, true, typ.IsInstantiable())
	assert.Equal(t, false, typ.IsMap())
	assert.Equal(t, false, typ.IsPointer())
	assert.Equal(t, false, typ.IsArray())
	assert.Equal(t, false, typ.IsString())
	assert.Equal(t, false, typ.IsBoolean())
	assert.Equal(t, false, typ.IsSlice())

	typ = GetType(&Animal{})
	assert.Equal(t, true, typ.IsStruct())
	assert.Equal(t, false, typ.IsInterface())
	assert.Equal(t, false, typ.IsFunction())
	assert.Equal(t, false, typ.IsNumber())
	assert.Equal(t, true, typ.IsInstantiable())
	assert.Equal(t, false, typ.IsMap())
	assert.Equal(t, true, typ.IsPointer())
	assert.Equal(t, false, typ.IsArray())
	assert.Equal(t, false, typ.IsString())
	assert.Equal(t, false, typ.IsBoolean())
	assert.Equal(t, false, typ.IsSlice())
}

func TestGetNamesForStruct(t *testing.T) {
	testGetNamesForStruct(t, GetType(Animal{}))
	testGetNamesForStruct(t, GetType(&Animal{}))
}

func testGetNamesForStruct(t *testing.T, typ Type) {
	assert.Equal(t, "Animal", typ.GetName())
	assert.Equal(t, "github.com.codnect.goo.Animal", typ.GetFullName())
	assert.Equal(t, "goo", typ.GetPackageName())
	assert.Equal(t, "github.com.codnect.goo", typ.GetPackageFullName())
	assert.Equal(t, typ.(Struct), typ.ToStruct())
}

func TestGetFieldsForStruct(t *testing.T) {
	testGetFieldsForStruct(t, GetType(Person{}))
	testGetFieldsForStruct(t, GetType(&Person{}))
}

func testGetFieldsForStruct(t *testing.T, typ Type) {
	structType := typ.(Struct)
	// all fields
	fieldCount := structType.GetFieldCount()
	assert.Equal(t, 4, fieldCount)
	fields := structType.GetFields()
	assert.Equal(t, 4, len(fields))

	// exported fields
	fieldCount = structType.GetExportedFieldCount()
	assert.Equal(t, 2, fieldCount)
	fields = structType.GetExportedFields()
	assert.Equal(t, 2, len(fields))

	// unexported fields
	fieldCount = structType.GetUnexportedFieldCount()
	assert.Equal(t, 2, fieldCount)
	fields = structType.GetUnexportedFields()
	assert.Equal(t, 2, len(fields))

	// anonymous fields
	fieldCount = structType.GetAnonymousFieldCount()
	assert.Equal(t, 0, fieldCount)
	fields = structType.GetAnonymousFields()
	assert.Equal(t, 0, len(fields))
}

func TestGetMethodsForStruct(t *testing.T) {
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

func TestImplementsForStruct(t *testing.T) {
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
