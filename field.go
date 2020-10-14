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
}

type MemberField struct {
	name        string
	typ         Type
	isAnonymous bool
	tags        reflect.StructTag
	isExported  bool
}

func newMemberField(name string, typ Type, isAnonymous bool, tags reflect.StructTag, isExported bool) MemberField {
	return MemberField{
		name,
		typ,
		isAnonymous,
		tags,
		isExported,
	}
}

func (field MemberField) GetName() string {
	return field.name
}

func (field MemberField) IsAnonymous() bool {
	return field.isAnonymous
}

func (field MemberField) IsExported() bool {
	return field.isExported
}

func (field MemberField) GetTags() []Tag {
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

func (field MemberField) GetTagByName(name string) (Tag, error) {
	value, ok := field.tags.Lookup(name)
	if ok {
		tag := Tag{name, value}
		return tag, nil
	}
	errText := fmt.Sprintf("Tag named %s not found ", name)
	return Tag{}, errors.New(errText)
}

func (field MemberField) GetType() Type {
	return field.typ
}

func (field MemberField) String() string {
	return field.name
}
