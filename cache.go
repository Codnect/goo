package goo

import (
	"reflect"
	"sync"
)

var cacheTypeMap map[string]Type
var cacheTypeMu sync.RWMutex

var cacheMethodMap map[string][]Method
var cacheMethodMu sync.RWMutex

var cacheFieldMap map[string][]Field
var cacheFieldMu sync.RWMutex

func initCache() {
	cacheTypeMap = make(map[string]Type, 0)
	cacheMethodMap = make(map[string][]Method, 0)
	cacheFieldMap = make(map[string][]Field, 0)
}

func getTypeFromCache(typ reflect.Type, isPointer bool) Type {
	if reflect.Interface != typ.Kind() && reflect.Struct != typ.Kind() {
		return nil
	}
	typeName := getStructOrInterfaceFullName(typ)
	defer func() {
		cacheTypeMu.Unlock()
	}()
	var cacheKeyName string
	if isPointer {
		cacheKeyName = "$" + typeName
	} else {
		cacheKeyName = typeName
	}
	cacheTypeMu.Lock()
	if typ, ok := cacheTypeMap[cacheKeyName]; ok {
		return typ
	}
	return nil
}

func putTypeIntoCache(typ Type, isPointer bool) Type {
	if !typ.IsInterface() && !typ.IsStruct() {
		return typ
	}
	defer func() {
		cacheTypeMu.Unlock()
	}()
	cacheTypeMu.Lock()
	var cacheKeyName string
	if isPointer {
		cacheKeyName = "$" + typ.GetFullName()
	} else {
		cacheKeyName = typ.GetFullName()
	}
	cacheTypeMap[cacheKeyName] = typ
	return typ
}

func getMethodsFromCache(key string) []Method {
	defer func() {
		cacheMethodMu.Unlock()
	}()
	cacheMethodMu.Lock()
	if methods, ok := cacheMethodMap[key+"#methods"]; ok {
		return methods
	}
	return nil
}

func putMethodsIntoCache(key string, methods []Method) []Method {
	defer func() {
		cacheMethodMu.Unlock()
	}()
	cacheMethodMu.Lock()
	cacheMethodMap[key+"#methods"] = methods
	return methods
}

func getFieldsFromCache(key string) []Field {
	defer func() {
		cacheFieldMu.Unlock()
	}()
	cacheFieldMu.Lock()
	if fields, ok := cacheFieldMap[key+"#fields"]; ok {
		return fields
	}
	return nil
}

func putFieldsIntoCache(key string, fields []Field) []Field {
	defer func() {
		cacheFieldMu.Unlock()
	}()
	cacheFieldMu.Lock()
	cacheFieldMap[key+"#fields"] = fields
	return fields
}
