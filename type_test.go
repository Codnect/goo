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

type Cell struct {
}

type Animal struct {
	Name string
	Cell
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

func (dog Dog) Test(arg string, i *int) {

}

func (dog Dog) TestOutputParam() string {
	return "TestOutputParam"
}

type Person struct {
	name    string
	Surname string
	age     int
	Address Address
	Cell
}

func (person Person) GetName() string {
	return person.name
}

func (person Person) GetSurname() string {
	return person.Surname
}

func (person Person) getAge() int {
	return person.age
}

func (person *Person) GetAddress() Address {
	return person.Address
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

func TestType_IsMethods(t *testing.T) {
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

func TestType_Instantiable(t *testing.T) {
	typ := GetType(Animal{})
	assert.True(t, typ.IsInstantiable())

	typ = GetType((*testInterface)(nil))
	assert.False(t, typ.IsInstantiable())

	typ = GetType(testFunction)
	assert.False(t, typ.IsInstantiable())
}

func TestBaseType_Equals(t *testing.T) {
	typ := GetType(Animal{})
	assert.True(t, typ.Equals(GetType(Animal{})))
	assert.False(t, typ.Equals(GetType(Dog{})))
	assert.False(t, typ.Equals(nil))
}

func TestType_GetTypeFromGoTypeWithNil(t *testing.T) {
	assert.Panics(t, func() {
		getTypeFromGoType(nil)
	})
}
