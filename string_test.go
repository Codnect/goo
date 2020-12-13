package goo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
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
	typ := GetType("")
	assert.True(t, typ.IsString())

	stringType := typ.ToStringType()

	// float32
	val, err := stringType.ToNumber("23.75", GetType(float32(0)).ToNumberType())
	assert.Nil(t, err)
	assert.Equal(t, float32(23.75), val)

	val, err = stringType.ToNumber(fmt.Sprintf("%f", math.MaxFloat64-1.0), GetType(float32(0)).ToNumberType())
	assert.NotNil(t, err)

	val, err = stringType.ToNumber("", GetType(float32(0)).ToNumberType())
	assert.NotNil(t, err)

	// float64
	val, err = stringType.ToNumber("23.75", GetType(float64(0)).ToNumberType())
	assert.Nil(t, err)
	assert.Equal(t, 23.75, val)

	val, err = stringType.ToNumber("", GetType(float64(0)).ToNumberType())
	assert.NotNil(t, err)

	// int
	val, err = stringType.ToNumber("23", GetType(0).ToNumberType())
	assert.Nil(t, err)
	assert.Equal(t, 23, val)

	val, err = stringType.ToNumber("", GetType(0).ToNumberType())
	assert.NotNil(t, err)

	assert.Panics(t, func() {
		stringType.ToNumber("", nil)
	})

	// int8
	val, err = stringType.ToNumber("23", GetType(int8(0)).ToNumberType())
	assert.Nil(t, err)
	assert.Equal(t, int8(23), val)

	val, err = stringType.ToNumber("-128", GetType(int8(0)).ToNumberType())
	assert.Nil(t, err)
	assert.Equal(t, int8(-128), val)

	val, err = stringType.ToNumber("-150", GetType(int8(0)).ToNumberType())
	assert.NotNil(t, err)

	val, err = stringType.ToNumber("", GetType(int8(0)).ToNumberType())
	assert.NotNil(t, err)

	// int16
	val, err = stringType.ToNumber("19421", GetType(int16(0)).ToNumberType())
	assert.Nil(t, err)
	assert.Equal(t, int16(19421), val)

	val, err = stringType.ToNumber("-15040", GetType(int16(0)).ToNumberType())
	assert.Nil(t, err)
	assert.Equal(t, int16(-15040), val)

	val, err = stringType.ToNumber("32980", GetType(int16(0)).ToNumberType())
	assert.NotNil(t, err)

	val, err = stringType.ToNumber("", GetType(int16(0)).ToNumberType())
	assert.NotNil(t, err)

	// int32
	val, err = stringType.ToNumber("243293245", GetType(int32(0)).ToNumberType())
	assert.Nil(t, err)
	assert.Equal(t, int32(243293245), val)

	val, err = stringType.ToNumber("-243293245", GetType(int32(0)).ToNumberType())
	assert.Nil(t, err)
	assert.Equal(t, int32(-243293245), val)

	val, err = stringType.ToNumber("23243293245", GetType(int32(0)).ToNumberType())
	assert.NotNil(t, err)

	val, err = stringType.ToNumber("", GetType(int64(0)).ToNumberType())
	assert.NotNil(t, err)

	// int64
	val, err = stringType.ToNumber("23243293245", GetType(int64(0)).ToNumberType())
	assert.Nil(t, err)
	assert.Equal(t, int64(23243293245), val)

	val, err = stringType.ToNumber("-23243293245", GetType(int64(0)).ToNumberType())
	assert.Nil(t, err)
	assert.Equal(t, int64(-23243293245), val)

	val, err = stringType.ToNumber("23545243293245741354", GetType(int64(0)).ToNumberType())
	assert.NotNil(t, err)

	val, err = stringType.ToNumber("", GetType(int64(0)).ToNumberType())
	assert.NotNil(t, err)

	// unit8
	val, err = stringType.ToNumber("23", GetType(uint8(0)).ToNumberType())
	assert.Nil(t, err)
	assert.Equal(t, uint8(23), val)

	val, err = stringType.ToNumber("-150", GetType(uint8(0)).ToNumberType())
	assert.NotNil(t, err)

	val, err = stringType.ToNumber("258", GetType(uint8(0)).ToNumberType())
	assert.NotNil(t, err)

	val, err = stringType.ToNumber("", GetType(uint8(0)).ToNumberType())
	assert.NotNil(t, err)

	// uint16
	val, err = stringType.ToNumber("19874", GetType(uint16(0)).ToNumberType())
	assert.Nil(t, err)
	assert.Equal(t, uint16(19874), val)

	val, err = stringType.ToNumber("-150", GetType(uint16(0)).ToNumberType())
	assert.NotNil(t, err)

	val, err = stringType.ToNumber("68419", GetType(uint16(0)).ToNumberType())
	assert.NotNil(t, err)

	val, err = stringType.ToNumber("", GetType(uint16(0)).ToNumberType())
	assert.NotNil(t, err)

	// uint32
	val, err = stringType.ToNumber("68941", GetType(uint32(0)).ToNumberType())
	assert.Nil(t, err)
	assert.Equal(t, uint32(68941), val)

	val, err = stringType.ToNumber("-150", GetType(uint32(0)).ToNumberType())
	assert.NotNil(t, err)

	val, err = stringType.ToNumber("254684571411", GetType(uint32(0)).ToNumberType())
	assert.NotNil(t, err)

	val, err = stringType.ToNumber("", GetType(uint32(0)).ToNumberType())
	assert.NotNil(t, err)

	// uint64
	val, err = stringType.ToNumber("254684571411", GetType(uint64(0)).ToNumberType())
	assert.Nil(t, err)
	assert.Equal(t, uint64(254684571411), val)

	val, err = stringType.ToNumber("-150", GetType(uint64(0)).ToNumberType())
	assert.NotNil(t, err)

	val, err = stringType.ToNumber("254684571411656202321", GetType(uint64(0)).ToNumberType())
	assert.NotNil(t, err)

	val, err = stringType.ToNumber("", GetType(uint64(0)).ToNumberType())
	assert.NotNil(t, err)

	val, err = stringType.ToNumber("", GetType(complex(1, 2)).ToNumberType())
	assert.NotNil(t, err)
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
	typ := GetType("")
	assert.True(t, typ.IsString())

	stringType := typ.ToStringType()

	result := stringType.ToUint("68941")
	assert.Equal(t, uint(68941), result)

	assert.Panics(t, func() {
		result = stringType.ToUint("-150")
	})
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
