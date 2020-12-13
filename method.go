package goo

import "reflect"

type Method interface {
	Member
	Invokable
	GetMethodReturnTypeCount() int
	GetMethodReturnTypes() []Type
	GetMethodParameterCount() int
	GetMethodParameterTypes() []Type
}

type memberMethod struct {
	typ        reflect.Type
	name       string
	isExported bool
	fun        reflect.Value
}

func newMemberMethod(methodType reflect.Type, name string, isExported bool, fun reflect.Value) memberMethod {
	return memberMethod{
		methodType,
		name,
		isExported,
		fun,
	}
}

func (method memberMethod) GetName() string {
	return method.name
}

func (method memberMethod) IsExported() bool {
	return method.isExported
}

func (method memberMethod) String() string {
	return method.name
}

func (method memberMethod) Invoke(obj interface{}, args ...interface{}) interface{} {
	return nil
}

func (method memberMethod) GetMethodReturnTypeCount() int {
	return method.typ.NumOut()
}

func (method memberMethod) GetMethodReturnTypes() []Type {
	returnTypes := make([]Type, 0)
	returnTypeCount := method.GetMethodReturnTypeCount()
	for returnTypeIndex := 0; returnTypeIndex < returnTypeCount; returnTypeIndex++ {
		returnType := method.typ.Out(returnTypeIndex)
		returnTypes = append(returnTypes, getTypeFromGoType(returnType))
	}
	return returnTypes
}

func (method memberMethod) GetMethodParameterCount() int {
	return method.typ.NumIn()
}

func (method memberMethod) GetMethodParameterTypes() []Type {
	parameterTypes := make([]Type, 0)
	parameterCount := method.GetMethodParameterCount()
	for paramIndex := 0; paramIndex < parameterCount; paramIndex++ {
		paramTyp := method.typ.In(paramIndex)
		parameterTypes = append(parameterTypes, getTypeFromGoType(paramTyp))
	}
	return parameterTypes
}
