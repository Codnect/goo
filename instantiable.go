package goo

type Instantiable interface {
	NewInstance() interface{}
}
