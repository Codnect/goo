package goo

import (
	"reflect"
	"runtime"
	"strings"
	"unicode"
)

func sanitizedName(str string) string {
	name := strings.ReplaceAll(str, "/", ".")
	name = strings.ReplaceAll(name, "-", ".")
	name = strings.ReplaceAll(name, "_", ".")
	return name
}

func getActualTypeFromBaseType(baseTyp *baseType) Type {
	if baseTyp.IsFunction() {
		return newFunctionType(baseTyp)
	} else if baseTyp.IsInterface() {
		return newInterfaceType(baseTyp)
	} else if baseTyp.IsStruct() {
		return newStructType(baseTyp)
	} else if baseTyp.IsNumber() {
		if isSignedInteger(baseTyp.typ) {
			return newSignedIntegerType(baseTyp)
		} else if isUnsignedInteger(baseTyp.typ) {
			return newUnsignedIntegerType(baseTyp)
		} else if isFloat(baseTyp.typ) {
			return newFloatType(baseTyp)
		} else if isComplex(baseTyp.typ) {
			return newComplexType(baseTyp)
		}
	} else if baseTyp.IsString() {
		return newStringType(baseTyp)
	} else if baseTyp.IsBoolean() {
		return newBooleanType(baseTyp)
	} else if baseTyp.IsMap() {
		return newMapType(baseTyp)
	} else if baseTyp.IsArray() {
		return newArrayType(baseTyp)
	} else if baseTyp.IsSlice() {
		return newSliceType(baseTyp)
	}
	panic(baseTyp.GetName() + " isn't supported for now")
}

func createBaseType(typ reflect.Type, val reflect.Value) *baseType {
	return newBaseType(typ, val)
}

func getTypeName(typ reflect.Type, val reflect.Value) string {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	switch typ.Kind() {
	case reflect.Struct:
		return getBaseTypeName(typ)
	case reflect.Interface:
		return getBaseTypeName(typ)
	case reflect.Func:
		return getFunctionName(val)
	}
	return typ.Name()
}

func getGoTypeAndValue(obj interface{}) (reflect.Type, reflect.Value, bool) {
	typ := reflect.TypeOf(obj)
	if typ == nil {
		panic("Type cannot be determined as the given object is nil")
	}
	isPointer := false
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		isPointer = true
	}
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return typ, val, isPointer
}

func getGoPointerTypeAndValue(obj interface{}) (reflect.Type, reflect.Value) {
	var pointerType reflect.Type
	var pointerVal reflect.Value
	typ := reflect.TypeOf(obj)
	if typ == nil {
		panic("Type cannot be determined as the given object is nil")
	}
	if typ.Kind() == reflect.Ptr {
		pointerType = typ
	}
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		pointerVal = val.Elem()
	}
	return pointerType, pointerVal
}

func getBaseTypeName(typ reflect.Type) string {
	name := typ.Name()
	if name == "" {
		name = typ.String()
	}
	return name
}

func getPackageName(typ reflect.Type, val reflect.Value) string {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	if reflect.Func == typ.Kind() {
		return getFunctionPackageName(val)
	}
	dotLastIndex := strings.LastIndex(typ.String(), ".")
	if dotLastIndex != -1 {
		return typ.String()[:dotLastIndex]
	}
	return ""
}

func getPackageFullName(typ reflect.Type, val reflect.Value) string {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	if reflect.Func == typ.Kind() {
		return getFunctionPackageFullName(val)
	}
	return sanitizedName(typ.PkgPath())
}

func getFunctionName(val reflect.Value) string {
	fullName := runtime.FuncForPC(val.Pointer()).Name()
	dotLastIndex := strings.LastIndex(fullName, ".")
	if dotLastIndex != -1 {
		fullName = fullName[dotLastIndex+1:]
	}
	return fullName
}

func getFunctionPackageFullName(val reflect.Value) string {
	fullName := runtime.FuncForPC(val.Pointer()).Name()
	dotLastIndex := strings.LastIndex(fullName, ".")
	if dotLastIndex != -1 {
		fullName = fullName[:dotLastIndex]
	}
	return sanitizedName(fullName)
}

func getFunctionPackageName(val reflect.Value) string {
	fullName := getFunctionPackageFullName(val)
	dotLastIndex := strings.LastIndex(fullName, ".")
	if dotLastIndex != -1 {
		fullName = fullName[dotLastIndex+1:]
	}
	return fullName
}

func convertGoFieldToMemberField(goField reflect.StructField) Field {
	field := newMemberField(goField.Name,
		getTypeFromGoType(goField.Type),
		goField.Anonymous,
		goField.Tag,
		isExportedField(goField))
	return field
}

func isExportedField(structField reflect.StructField) bool {
	return unicode.IsUpper(rune(structField.Name[0]))
}

func isExportedMethod(method reflect.Method) bool {
	return unicode.IsUpper(rune(method.Name[0]))
}

func convertGoMethodToMemberMethod(goMethod reflect.Method) Method {
	method := newMemberMethod(goMethod.Type,
		goMethod.Name,
		isExportedMethod(goMethod),
		goMethod.Func,
	)
	return method
}

func isNumber(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	case reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return true
	}
	return false
}

func isSignedInteger(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	}
	return false
}

func isUnsignedInteger(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	}
	return false
}

func isFloat(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

func isComplex(typ reflect.Type) bool {
	switch typ.Kind() {
	case reflect.Complex64, reflect.Complex128:
		return true
	}
	return false
}
