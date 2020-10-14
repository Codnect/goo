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

type MemberMethod struct {
	typ        reflect.Type
	name       string
	isExported bool
	fun        reflect.Value
}

func newMemberMethod(methodType reflect.Type, name string, isExported bool, fun reflect.Value) MemberMethod {
	return MemberMethod{
		methodType,
		name,
		isExported,
		fun,
	}
}

func (method MemberMethod) GetName() string {
	return method.name
}

func (method MemberMethod) IsExported() bool {
	return method.isExported
}

func (method MemberMethod) String() string {
	return method.name
}

func (method MemberMethod) Invoke(obj interface{}, args ...interface{}) interface{} {
	return nil
}

func (method MemberMethod) GetMethodReturnTypeCount() int {
	return method.typ.NumOut()
}

func (method MemberMethod) GetMethodReturnTypes() []Type {
	returnTypes := make([]Type, 0)
	returnTypeCount := method.GetMethodReturnTypeCount()
	for returnTypeIndex := 0; returnTypeIndex < returnTypeCount; returnTypeIndex++ {
		returnType := method.typ.Out(returnTypeIndex)
		returnTypes = append(returnTypes, GetTypeFromGoType(returnType))
	}
	return returnTypes
}

func (method MemberMethod) GetMethodParameterCount() int {
	return method.typ.NumIn()
}

func (method MemberMethod) GetMethodParameterTypes() []Type {
	parameterTypes := make([]Type, 0)
	parameterCount := method.GetMethodParameterCount()
	for paramIndex := 0; paramIndex < parameterCount; paramIndex++ {
		paramTyp := method.typ.In(paramIndex)
		parameterTypes = append(parameterTypes, GetTypeFromGoType(paramTyp))
	}
	return parameterTypes
}
