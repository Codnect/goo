package goo

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type Field interface {
	Member
	Taggable
	IsAnonymous() bool
	GetType() Type
	CanSet() bool
	SetValue(instance interface{}, value interface{})
	GetValue(instance interface{}) interface{}
}

type memberField struct {
	name        string
	typ         Type
	isAnonymous bool
	tags        reflect.StructTag
	isExported  bool
}

func newMemberField(name string, typ Type, isAnonymous bool, tags reflect.StructTag, isExported bool) memberField {
	return memberField{
		name,
		typ,
		isAnonymous,
		tags,
		isExported,
	}
}

func (field memberField) GetName() string {
	return field.name
}

func (field memberField) IsAnonymous() bool {
	return field.isAnonymous
}

func (field memberField) IsExported() bool {
	return field.isExported
}

func (field memberField) CanSet() bool {
	return field.isExported
}

func (field memberField) GetTags() []Tag {
	fieldTags := make([]Tag, 0)
	tags := field.tags
	for tags != "" {
		i := 0
		for i < len(tags) && tags[i] == ' ' {
			i++
		}
		tags = tags[i:]
		if tags == "" {
			break
		}

		i = 0
		for i < len(tags) && tags[i] > ' ' && tags[i] != ':' && tags[i] != '"' && tags[i] != 0x7f {
			i++
		}
		if i == 0 || i+1 >= len(tags) || tags[i] != ':' || tags[i+1] != '"' {
			break
		}
		name := string(tags[:i])
		tags = tags[i+1:]

		i = 1
		for i < len(tags) && tags[i] != '"' {
			if tags[i] == '\\' {
				i++
			}
			i++
		}
		if i >= len(tags) {
			break
		}
		quotedValue := string(tags[:i+1])
		tags = tags[i+1:]

		value, err := strconv.Unquote(quotedValue)
		if err != nil {
			break
		}

		fieldTag := Tag{name, value}
		fieldTags = append(fieldTags, fieldTag)
	}
	return fieldTags
}

func (field memberField) GetTagByName(name string) (Tag, error) {
	value, ok := field.tags.Lookup(name)
	if ok {
		tag := Tag{name, value}
		return tag, nil
	}
	errText := fmt.Sprintf("Tag named %s not found ", name)
	return Tag{}, errors.New(errText)
}

func (field memberField) GetType() Type {
	return field.typ
}

func (field memberField) String() string {
	return field.name
}

func (field memberField) SetValue(instance interface{}, value interface{}) {
	if !field.CanSet() {
		panic("Field cannot be set because of it is an unexported field")
	}
	typ := GetType(instance)
	if !typ.IsStruct() {
		panic("Instance must only be a struct")
	}
	if !typ.IsPointer() {
		panic("Instance type must be a pointer")
	}
	structType := typ.GetGoType()
	structValueType := typ.GetGoPointerValue()
	structFieldCount := structType.NumField()
	for fieldIndex := 0; fieldIndex < structFieldCount; fieldIndex++ {
		fieldType := structType.Field(fieldIndex)
		fieldValue := structValueType.Field(fieldIndex)
		if fieldType.Name == field.name {
			fieldValue.Set(reflect.ValueOf(value))
			break
		}
	}
}

func (field memberField) GetValue(instance interface{}) interface{} {
	typ := GetType(instance)
	if !typ.IsStruct() {
		panic("Instance must only be a struct")
	}
	structType := typ.GetGoType()
	structValueType := typ.GetGoValue()
	if typ.IsPointer() {
		structValueType = typ.GetGoPointerValue()
	}
	structFieldCount := structType.NumField()
	for fieldIndex := 0; fieldIndex < structFieldCount; fieldIndex++ {
		fieldType := structType.Field(fieldIndex)
		fieldValue := structValueType.Field(fieldIndex)

		if fieldType.Name == field.name {

			if !isExportedField(fieldType) {
				panic("Field is not exported, you cannot get the value : " + field.name)
			}

			if fieldType.Type.Kind() == reflect.Ptr {
				return fieldValue.Interface()
			}

			return fieldValue.Addr().Interface()
		}
	}
	return nil
}
