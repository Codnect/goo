package goo

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_isComplex(t *testing.T) {
	assert.True(t, isComplex(reflect.TypeOf(complex(14.3, 22.5))))
	assert.False(t, isComplex(reflect.TypeOf(23)))
}

func Test_getGoPointerTypeAndValueWithNil(t *testing.T) {
	assert.Panics(t, func() {
		getGoPointerTypeAndValue(nil)
	})
}

func Test_getGoTypeAndValueWithNil(t *testing.T) {
	assert.Panics(t, func() {
		getGoTypeAndValue(nil)
	})
}
