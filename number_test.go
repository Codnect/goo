package goo

import (
	"github.com/stretchr/testify/assert"
	"math/bits"
	"testing"
)

func TestSignedIntegerType(t *testing.T) {
	intType := GetType(8)
	assert.True(t, intType.IsNumber())

	intNumberType := intType.ToNumberType()
	assert.Equal(t, IntegerType, intNumberType.GetType())
	if bits.UintSize == BitSize32 {
		assert.Equal(t, BitSize32, intNumberType.GetBitSize())
	} else {
		assert.Equal(t, BitSize64, intNumberType.GetBitSize())
	}

	assert.Panics(t, func() {
		intNumberType.Overflow(&Animal{})
	})

	assert.Panics(t, func() {
		intNumberType.ToString("test")
	})

	int8Type := GetType(int8(8))
	assert.True(t, int8Type.IsNumber())

	int8NumberType := int8Type.ToNumberType()
	assert.Equal(t, IntegerType, int8NumberType.GetType())
	assert.Equal(t, BitSize8, int8NumberType.GetBitSize())

	assert.True(t, int8NumberType.Overflow(129))
	assert.True(t, int8NumberType.Overflow(-150))
	assert.Equal(t, "120", int8NumberType.ToString(120))

	assert.Panics(t, func() {
		int8NumberType.ToString("test")
	})

	int16Type := GetType(int16(25))
	assert.True(t, int16Type.IsNumber())

	int16NumberType := int16Type.ToNumberType()
	assert.Equal(t, IntegerType, int16NumberType.GetType())
	assert.Equal(t, BitSize16, int16NumberType.GetBitSize())

	assert.True(t, int16NumberType.Overflow(35974))
	assert.True(t, int16NumberType.Overflow(-39755))
	assert.Equal(t, "1575", int16NumberType.ToString(1575))

	assert.Panics(t, func() {
		int16NumberType.ToString("test")
	})

	int32Type := GetType(int32(25))
	assert.True(t, int32Type.IsNumber())

	int32NumberType := int32Type.ToNumberType()
	assert.Equal(t, IntegerType, int32NumberType.GetType())
	assert.Equal(t, BitSize32, int32NumberType.GetBitSize())

	assert.True(t, int32NumberType.Overflow(2443252523))
	assert.True(t, int32NumberType.Overflow(-2443252523))
	assert.Equal(t, "244325", int32NumberType.ToString(244325))

	assert.Panics(t, func() {
		int32NumberType.ToString("test")
	})

	int64Type := GetType(int64(25))
	assert.True(t, int32Type.IsNumber())

	int64NumberType := int64Type.ToNumberType()
	assert.Equal(t, IntegerType, int64NumberType.GetType())
	assert.Equal(t, BitSize64, int64NumberType.GetBitSize())
	assert.Equal(t, "244325", int64NumberType.ToString(244325))

	assert.Panics(t, func() {
		int64NumberType.ToString("test")
	})
}

func TestSignedIntegerType_NewInstance(t *testing.T) {
	intType := GetType(8)
	intNumberType := intType.ToNumberType()
	val := intNumberType.NewInstance()
	assert.NotNil(t, val.(*int))

	int8Type := GetType(int8(8))
	int8NumberType := int8Type.ToNumberType()
	val = int8NumberType.NewInstance()
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
	intType := GetType(uint(8))
	assert.True(t, intType.IsNumber())

	intNumberType := intType.ToNumberType()
	assert.Equal(t, IntegerType, intNumberType.GetType())
	if bits.UintSize == BitSize32 {
		assert.Equal(t, BitSize32, intNumberType.GetBitSize())
	} else {
		assert.Equal(t, BitSize64, intNumberType.GetBitSize())
	}

	assert.Panics(t, func() {
		intNumberType.Overflow(&Animal{})
	})

	assert.Panics(t, func() {
		intNumberType.ToString("test")
	})

	int8Type := GetType(uint8(8))
	assert.True(t, int8Type.IsNumber())

	int8NumberType := int8Type.ToNumberType()
	assert.Equal(t, IntegerType, int8NumberType.GetType())
	assert.Equal(t, BitSize8, int8NumberType.GetBitSize())

	assert.True(t, int8NumberType.Overflow(uint(280)))
	assert.Equal(t, "120", int8NumberType.ToString(uint(120)))

	assert.Panics(t, func() {
		int8NumberType.ToString("test")
	})

	int16Type := GetType(uint16(25))
	assert.True(t, int16Type.IsNumber())

	int16NumberType := int16Type.ToNumberType()
	assert.Equal(t, IntegerType, int16NumberType.GetType())
	assert.Equal(t, BitSize16, int16NumberType.GetBitSize())

	assert.True(t, int16NumberType.Overflow(uint(68954)))
	assert.Equal(t, "1575", int16NumberType.ToString(uint(1575)))

	assert.Panics(t, func() {
		int16NumberType.ToString("test")
	})

	int32Type := GetType(uint32(25))
	assert.True(t, int32Type.IsNumber())

	int32NumberType := int32Type.ToNumberType()
	assert.Equal(t, IntegerType, int32NumberType.GetType())
	assert.Equal(t, BitSize32, int32NumberType.GetBitSize())

	assert.True(t, int32NumberType.Overflow(uint(2443252687523)))
	assert.Equal(t, "244325", int32NumberType.ToString(uint(244325)))

	assert.Panics(t, func() {
		int32NumberType.ToString("test")
	})

	int64Type := GetType(uint64(25))
	assert.True(t, int32Type.IsNumber())

	int64NumberType := int64Type.ToNumberType()
	assert.Equal(t, IntegerType, int64NumberType.GetType())
	assert.Equal(t, BitSize64, int64NumberType.GetBitSize())
	assert.Equal(t, "244325", int64NumberType.ToString(uint(244325)))

	assert.Panics(t, func() {
		int64NumberType.ToString("test")
	})
}

func TestUnSignedIntegerType_NewInstance(t *testing.T) {
	intType := GetType(uint(8))
	intNumberType := intType.ToNumberType()
	val := intNumberType.NewInstance()
	assert.NotNil(t, val.(*uint))

	int8Type := GetType(uint8(8))
	int8NumberType := int8Type.ToNumberType()
	val = int8NumberType.NewInstance()
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

func TestComplexType_ToString(t *testing.T) {
	complexNumber := complex(14.3, 22.5)
	typ := GetType(complexNumber)
	assert.True(t, typ.IsNumber())

	numberType := typ.ToNumberType()
	assert.Equal(t, ComplexType, numberType.GetType())
	assert.Equal(t, "(14.300000+22.500000i)", numberType.ToString(complexNumber))

	assert.Panics(t, func() {
		numberType.ToString(23)
	})
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

	assert.Panics(t, func() {
		complexType.GetRealData(23)
	})
}

func TestComplexType_GetImaginaryData(t *testing.T) {
	complexNumber := complex(14.3, 22.5)
	typ := GetType(complexNumber)
	assert.True(t, typ.IsNumber())

	numberType := typ.ToNumberType()
	assert.Equal(t, ComplexType, numberType.GetType())

	complexType := numberType.(Complex)
	assert.Equal(t, 22.5, complexType.GetImaginaryData(complexNumber))

	assert.Panics(t, func() {
		complexType.GetImaginaryData(23)
	})
}

func TestFloatType_GetType(t *testing.T) {
	float32Number := float32(23.2)
	typ := GetType(float32Number)
	assert.True(t, typ.IsNumber())

	numberType := typ.ToNumberType()
	assert.Equal(t, FloatType, numberType.GetType())

	float64Number := 23.2
	typ = GetType(float64Number)
	assert.True(t, typ.IsNumber())

	numberType = typ.ToNumberType()
	assert.Equal(t, FloatType, numberType.GetType())
}

func TestFloatType_GetBitSize(t *testing.T) {
	float32Number := float32(23.2)
	typ := GetType(float32Number)
	assert.True(t, typ.IsNumber())

	numberType := typ.ToNumberType()
	assert.Equal(t, BitSize32, numberType.GetBitSize())

	float64Number := 23.2
	typ = GetType(float64Number)
	assert.True(t, typ.IsNumber())

	numberType = typ.ToNumberType()
	assert.Equal(t, BitSize64, numberType.GetBitSize())
}

func TestFloatType_NewInstance(t *testing.T) {
	float32Number := float32(23.2)
	typ := GetType(float32Number)
	assert.True(t, typ.IsNumber())

	numberType := typ.ToNumberType()
	val := numberType.NewInstance()
	assert.NotNil(t, val)

	float64Number := 23.2
	typ = GetType(float64Number)
	assert.True(t, typ.IsNumber())

	numberType = typ.ToNumberType()
	val = numberType.NewInstance()
	assert.NotNil(t, val)

}

func TestFloatType_Overflow(t *testing.T) {
	float32Number := float32(23.2)
	typ := GetType(float32Number)
	assert.True(t, typ.IsNumber())

	numberType := typ.ToNumberType()
	assert.Panics(t, func() {
		numberType.Overflow(&Animal{})
	})
}

func TestFloatType_ToString(t *testing.T) {
	floatNumber := 23.2
	typ := GetType(floatNumber)
	assert.True(t, typ.IsNumber())

	numberType := typ.ToNumberType()

	assert.Panics(t, func() {
		numberType.ToString(&Animal{})
	})

	assert.Equal(t, "23.200000", numberType.ToString(floatNumber))
}
