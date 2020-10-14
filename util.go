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

func getActualTypeFromBaseType(baseTyp baseType) Type {
	if baseTyp.IsFunction() {
		return newFunctionType(baseTyp)
	} else if baseTyp.IsInterface() {
		return newInterfaceType(baseTyp)
	} else if baseTyp.IsStruct() {
		return newStructType(baseTyp)
	}
	return baseTyp
}

func createBaseType(typ reflect.Type, val reflect.Value) baseType {
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
	return ""
}

func GetGoTypeAndValue(obj interface{}) (reflect.Type, reflect.Value) {
	typ := reflect.TypeOf(obj)
	if typ == nil {
		panic("Type cannot be determined as the given object is nil")
	}
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return typ, val
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

func getStructOrInterfaceFullName(typ reflect.Type) string {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	return sanitizedName(typ.PkgPath()) + "." + getBaseTypeName(typ)
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
		GetTypeFromGoType(goField.Type),
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
