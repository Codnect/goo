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

func (method memberMethod) Invoke(obj interface{}, args ...interface{}) []interface{} {
	typ := GetType(obj)
	if !typ.IsStruct() {
		panic("obj must be a struct instance")
	}

	structType := typ.ToStructType()
	structMethods := structType.GetStructMethods()
	for _, structMethod := range structMethods {
		if method.GetName() == structMethod.GetName() {
			parameterCount := method.GetMethodParameterCount()
			if (args == nil && parameterCount != 1) || (args != nil && len(args) != parameterCount-1) {
				panic("Parameter counts don't match argument counts")
			}

			inputs := make([]reflect.Value, 0)

			tempArray := make([]interface{}, 1)
			tempArray[0] = obj

			args = append(tempArray[:], args[:]...)

			for index, arg := range args {
				if arg == nil {
					paramType := method.GetMethodParameterTypes()[index]
					if paramType.IsPointer() {
						inputs = append(inputs, reflect.New(paramType.GetGoPointerType()).Elem())
					} else {
						inputs = append(inputs, reflect.New(paramType.GetGoType()).Elem())
					}
				} else {
					inputs = append(inputs, reflect.ValueOf(arg))
				}
			}

			outputs := make([]interface{}, 0)
			results := structMethod.(memberMethod).fun.Call(inputs)
			for _, outputParam := range results {
				outputs = append(outputs, outputParam.Interface())
			}
			return outputs
		}
	}
	panic("Method not found")
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
