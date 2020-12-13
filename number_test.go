package goo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignedIntegerType(t *testing.T) {
	int8Type := GetType(int8(8))
	assert.True(t, int8Type.IsNumber())

	int8NumberType := int8Type.ToNumberType()
	assert.Equal(t, IntegerType, int8NumberType.GetType())
	assert.Equal(t, BitSize8, int8NumberType.GetBitSize())

	assert.True(t, int8NumberType.Overflow(129))
	assert.True(t, int8NumberType.Overflow(-150))
	assert.Equal(t, "120", int8NumberType.ToString(120))

	int16Type := GetType(int16(25))
	assert.True(t, int16Type.IsNumber())

	int16NumberType := int16Type.ToNumberType()
	assert.Equal(t, IntegerType, int16NumberType.GetType())
	assert.Equal(t, BitSize16, int16NumberType.GetBitSize())

	assert.True(t, int16NumberType.Overflow(35974))
	assert.True(t, int16NumberType.Overflow(-39755))
	assert.Equal(t, "1575", int16NumberType.ToString(1575))

	int32Type := GetType(int32(25))
	assert.True(t, int32Type.IsNumber())

	int32NumberType := int32Type.ToNumberType()
	assert.Equal(t, IntegerType, int32NumberType.GetType())
	assert.Equal(t, BitSize32, int32NumberType.GetBitSize())

	assert.True(t, int32NumberType.Overflow(2443252523))
	assert.True(t, int32NumberType.Overflow(-2443252523))
	assert.Equal(t, "244325", int32NumberType.ToString(244325))

	int64Type := GetType(int64(25))
	assert.True(t, int32Type.IsNumber())

	int64NumberType := int64Type.ToNumberType()
	assert.Equal(t, IntegerType, int64NumberType.GetType())
	assert.Equal(t, BitSize64, int64NumberType.GetBitSize())
	assert.Equal(t, "244325", int64NumberType.ToString(244325))
}

func TestSignedIntegerType_NewInstance(t *testing.T) {
	int8Type := GetType(int8(8))
	int8NumberType := int8Type.ToNumberType()
	val := int8NumberType.NewInstance()
	assert.NotNil(t, val.(*int8))

	int16Type := GetType(int16(25))
	int16NumberType := int16Type.ToNumberType()
	val = int16NumberType.NewInstance()
	assert.NotNil(t, val.(*int16))

	int32Type := GetType(int32(25))
	int32NumberType := int32Type.ToNumberType()
	val = int32NumberType.NewInstance()
	assert.NotNil(t, val.(*int32))

	int64Type := GetType(int64(25))
	int64NumberType := int64Type.ToNumberType()
	val = int64NumberType.NewInstance()
	assert.NotNil(t, val.(*int64))
}

func TestUnSignedIntegerType(t *testing.T) {
	int8Type := GetType(uint8(8))
	assert.True(t, int8Type.IsNumber())

	int8NumberType := int8Type.ToNumberType()
	assert.Equal(t, IntegerType, int8NumberType.GetType())
	assert.Equal(t, BitSize8, int8NumberType.GetBitSize())

	assert.True(t, int8NumberType.Overflow(uint(280)))
	assert.Equal(t, "120", int8NumberType.ToString(uint(120)))

	int16Type := GetType(uint16(25))
	assert.True(t, int16Type.IsNumber())

	int16NumberType := int16Type.ToNumberType()
	assert.Equal(t, IntegerType, int16NumberType.GetType())
	assert.Equal(t, BitSize16, int16NumberType.GetBitSize())

	assert.True(t, int16NumberType.Overflow(uint(68954)))
	assert.Equal(t, "1575", int16NumberType.ToString(uint(1575)))

	int32Type := GetType(uint32(25))
	assert.True(t, int32Type.IsNumber())

	int32NumberType := int32Type.ToNumberType()
	assert.Equal(t, IntegerType, int32NumberType.GetType())
	assert.Equal(t, BitSize32, int32NumberType.GetBitSize())

	assert.True(t, int32NumberType.Overflow(uint(2443252687523)))
	assert.Equal(t, "244325", int32NumberType.ToString(uint(244325)))

	int64Type := GetType(uint64(25))
	assert.True(t, int32Type.IsNumber())

	int64NumberType := int64Type.ToNumberType()
	assert.Equal(t, IntegerType, int64NumberType.GetType())
	assert.Equal(t, BitSize64, int64NumberType.GetBitSize())
	assert.Equal(t, "244325", int64NumberType.ToString(uint(244325)))
}

func TestUnSignedIntegerType_NewInstance(t *testing.T) {
	int8Type := GetType(uint8(8))
	int8NumberType := int8Type.ToNumberType()
	val := int8NumberType.NewInstance()
	assert.NotNil(t, val.(*uint8))

	int16Type := GetType(uint16(25))
	int16NumberType := int16Type.ToNumberType()
	val = int16NumberType.NewInstance()
	assert.NotNil(t, val.(*uint16))

	int32Type := GetType(uint32(25))
	int32NumberType := int32Type.ToNumberType()
	val = int32NumberType.NewInstance()
	assert.NotNil(t, val.(*uint32))

	int64Type := GetType(uint64(25))
	int64NumberType := int64Type.ToNumberType()
	val = int64NumberType.NewInstance()
	assert.NotNil(t, val.(*uint64))
}

func TestComplexType_GetType(t *testing.T) {

}

func TestComplexType_ToString(t *testing.T) {
	complexNumber := complex(14.3, 22.5)
	typ := GetType(complexNumber)
	assert.True(t, typ.IsNumber())

	numberType := typ.ToNumberType()
	assert.Equal(t, ComplexType, numberType.GetType())
	assert.Equal(t, "(14.300000+22.500000i)", numberType.ToString(complexNumber))
}

func TestComplexType_GetBitSize(t *testing.T) {
	complexNumber := complex(14.3, 22.5)
	typ := GetType(complexNumber)
	assert.True(t, typ.IsNumber())

	numberType := typ.ToNumberType()
	assert.Equal(t, ComplexType, numberType.GetType())
	assert.Equal(t, BitSize128, numberType.GetBitSize())
}

func TestComplexType_Overflow(t *testing.T) {
	complexNumber := complex(14.3, 22.5)
	typ := GetType(complexNumber)
	assert.True(t, typ.IsNumber())

	numberType := typ.ToNumberType()
	assert.Equal(t, ComplexType, numberType.GetType())
	assert.Panics(t, func() {
		numberType.Overflow(nil)
	})
}

func TestComplexType_NewInstance(t *testing.T) {
	complexNumber := complex(14.3, 22.5)
	typ := GetType(complexNumber)
	assert.True(t, typ.IsNumber())

	numberType := typ.ToNumberType()
	assert.Equal(t, ComplexType, numberType.GetType())

	instance := numberType.NewInstance()
	assert.NotNil(t, instance)
}

func TestComplexType_GetRealData(t *testing.T) {
	complexNumber := complex(14.3, 22.5)
	typ := GetType(complexNumber)
	assert.True(t, typ.IsNumber())

	numberType := typ.ToNumberType()
	assert.Equal(t, ComplexType, numberType.GetType())

	complexType := numberType.(Complex)
	assert.Equal(t, 14.3, complexType.GetRealData(complexNumber))
}

func TestComplexType_GetImaginaryData(t *testing.T) {
	complexNumber := complex(14.3, 22.5)
	typ := GetType(complexNumber)
	assert.True(t, typ.IsNumber())

	numberType := typ.ToNumberType()
	assert.Equal(t, ComplexType, numberType.GetType())

	complexType := numberType.(Complex)
	assert.Equal(t, 22.5, complexType.GetImaginaryData(complexNumber))
}
