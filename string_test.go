package goo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringType_NewInstance(t *testing.T) {
	typ := GetType("")
	assert.True(t, typ.IsString())

	stringType := typ.ToStringType()
	stringVal := stringType.NewInstance()
	assert.NotNil(t, stringVal)
}

func TestStringType_ToFloat32(t *testing.T) {
	typ := GetType("")
	assert.True(t, typ.IsString())

	stringType := typ.ToStringType()
	floatVal := stringType.ToFloat32("23.22")
	assert.Equal(t, float32(23.22), floatVal)

	assert.Panics(t, func() {
		stringType.ToFloat32("")
	})
}

func TestStringType_ToFloat64(t *testing.T) {
	typ := GetType("")
	assert.True(t, typ.IsString())

	stringType := typ.ToStringType()
	floatVal := stringType.ToFloat64("23.22")
	assert.Equal(t, 23.22, floatVal)

	assert.Panics(t, func() {
		stringType.ToFloat64("")
	})
}

func TestStringType_ToNumber(t *testing.T) {

}

func TestStringType_ToInt(t *testing.T) {
	typ := GetType("")
	assert.True(t, typ.IsString())

	stringType := typ.ToStringType()
	result := stringType.ToInt("23")

	assert.Equal(t, 23, result)

	assert.Panics(t, func() {
		stringType.ToInt("")
	})
}

func TestStringType_ToInt8(t *testing.T) {
	typ := GetType("")
	assert.True(t, typ.IsString())

	stringType := typ.ToStringType()

	result := stringType.ToInt8("23")
	assert.Equal(t, int8(23), result)

	result = stringType.ToInt8("-128")
	assert.Equal(t, int8(-128), result)

	assert.Panics(t, func() {
		result = stringType.ToInt8("150")
	})

	assert.Panics(t, func() {
		result = stringType.ToInt8("-130")
	})

	assert.Panics(t, func() {
		stringType.ToInt8("")
	})
}

func TestStringType_ToInt16(t *testing.T) {
	typ := GetType("")
	assert.True(t, typ.IsString())

	stringType := typ.ToStringType()

	result := stringType.ToInt16("19421")
	assert.Equal(t, int16(19421), result)

	result = stringType.ToInt16("-15040")
	assert.Equal(t, int16(-15040), result)

	assert.Panics(t, func() {
		result = stringType.ToInt16("32980")
	})

	assert.Panics(t, func() {
		result = stringType.ToInt16("-35874")
	})

	assert.Panics(t, func() {
		stringType.ToInt16("")
	})
}

func TestStringType_ToInt32(t *testing.T) {
	typ := GetType("")
	assert.True(t, typ.IsString())

	stringType := typ.ToStringType()

	result := stringType.ToInt32("243293245")
	assert.Equal(t, int32(243293245), result)

	result = stringType.ToInt32("-243293245")
	assert.Equal(t, int32(-243293245), result)

	assert.Panics(t, func() {
		result = stringType.ToInt32("23243293245")
	})

	assert.Panics(t, func() {
		result = stringType.ToInt32("-23243293245")
	})

	assert.Panics(t, func() {
		stringType.ToInt32("")
	})
}

func TestStringType_ToInt64(t *testing.T) {
	typ := GetType("")
	assert.True(t, typ.IsString())

	stringType := typ.ToStringType()

	result := stringType.ToInt64("23243293245")
	assert.Equal(t, int64(23243293245), result)

	result = stringType.ToInt64("-23243293245")
	assert.Equal(t, int64(-23243293245), result)

	assert.Panics(t, func() {
		result = stringType.ToInt64("23545243293245741354")
	})

	assert.Panics(t, func() {
		result = stringType.ToInt64("-23545243293245741354")
	})

	assert.Panics(t, func() {
		stringType.ToInt64("")
	})
}

func TestStringType_ToUint(t *testing.T) {

}

func TestStringType_ToUint8(t *testing.T) {
	typ := GetType("")
	assert.True(t, typ.IsString())

	stringType := typ.ToStringType()

	result := stringType.ToUint8("23")
	assert.Equal(t, uint8(23), result)

	assert.Panics(t, func() {
		result = stringType.ToUint8("-150")
	})

	assert.Panics(t, func() {
		result = stringType.ToUint8("258")
	})

	assert.Panics(t, func() {
		stringType.ToUint16("")
	})
}

func TestStringType_ToUint16(t *testing.T) {
	typ := GetType("")
	assert.True(t, typ.IsString())

	stringType := typ.ToStringType()

	result := stringType.ToUint16("19874")
	assert.Equal(t, uint16(19874), result)

	assert.Panics(t, func() {
		result = stringType.ToUint16("-150")
	})

	assert.Panics(t, func() {
		result = stringType.ToUint16("68419")
	})

	assert.Panics(t, func() {
		stringType.ToUint16("")
	})
}

func TestStringType_ToUint32(t *testing.T) {
	typ := GetType("")
	assert.True(t, typ.IsString())

	stringType := typ.ToStringType()

	result := stringType.ToUint32("68941")
	assert.Equal(t, uint32(68941), result)

	assert.Panics(t, func() {
		result = stringType.ToUint32("254684571411")
	})

	assert.Panics(t, func() {
		result = stringType.ToUint32("-150")
	})

	assert.Panics(t, func() {
		stringType.ToUint32("")
	})
}

func TestStringType_ToUint64(t *testing.T) {
	typ := GetType("")
	assert.True(t, typ.IsString())

	stringType := typ.ToStringType()

	result := stringType.ToUint64("254684571411")
	assert.Equal(t, uint64(254684571411), result)

	assert.Panics(t, func() {
		result = stringType.ToUint64("254684571411656202321")
	})

	assert.Panics(t, func() {
		result = stringType.ToUint64("-150")
	})

	assert.Panics(t, func() {
		stringType.ToUint64("")
	})
}
