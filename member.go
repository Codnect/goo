package goo

type Exportable interface {
	IsExported() bool
}

type Invokable interface {
	Invoke(obj interface{}, args ...interface{}) []interface{}
}

type Member interface {
	Exportable
	GetName() string
	String() string
}
