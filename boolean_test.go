package goo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBooleanType_NewInstance(t *testing.T) {
	typ := GetType(true)
	assert.True(t, typ.IsBoolean())

	booleanType := typ.ToBooleanType()
	instance := booleanType.NewInstance()
	assert.NotNil(t, instance)

	boolPtr := instance.(*bool)
	assert.False(t, *boolPtr)
}

func TestBooleanType_ToString(t *testing.T) {
	typ := GetType(true)
	assert.True(t, typ.IsBoolean())
	booleanType := typ.ToBooleanType()

	assert.Equal(t, "true", booleanType.ToString(true))
	assert.Equal(t, "false", booleanType.ToString(false))
}

func TestBooleanType_ToBoolean(t *testing.T) {
	typ := GetType(true)
	assert.True(t, typ.IsBoolean())
	booleanType := typ.ToBooleanType()

	assert.True(t, booleanType.ToBoolean("true"))
	assert.False(t, booleanType.ToBoolean("false"))
	assert.Panics(t, func() {
		booleanType.ToBoolean("test")
	})
}
