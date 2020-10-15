package goo

import (
	"errors"
	"math"
	"math/bits"
	"reflect"
	"strconv"
)

type String interface {
	Type
	ToNumber(val string, number Number) (interface{}, error)
	ToInt(val string) int
	ToInt8(val string) int8
	ToInt16(val string) int16
	ToInt32(val string) int32
	ToInt64(val string) int64
	ToUint(val string) uint
	ToUint8(val string) uint8
	ToUint16(val string) uint16
	ToUint32(val string) uint32
	ToUint64(val string) uint64
	ToFloat32(val string) float32
	ToFloat64(val string) float64
}

type stringType struct {
	baseType
}

func newStringType(baseTyp baseType) stringType {
	return stringType{
		baseTyp,
	}
}

func (str stringType) ToNumber(val string, number Number) (interface{}, error) {
	if number == nil {
		panic("Number must not be null")
	}
	numberType := number.GetNumberType()
	if IntegerType == numberType {
		return str.getIntegerValue(val, number.(Integer))
	} else if FloatType == numberType {
		return str.getFloatValue(val, number.(Float))
	} else if ComplexType == numberType {
		return nil, errors.New("complex numbers does not support for now")
	}
	panic("number type is not valid")
}

func (str stringType) ToInt(val string) int {
	sizeInBits := bits.UintSize
	var result interface{}
	var err error
	if sizeInBits == 32 {
		result, err = str.getIntegerValueByBitSize(val, BitSize32, true)
	} else {
		result, err = str.getIntegerValueByBitSize(val, BitSize64, true)
	}
	if err != nil {
		panic(err)
	}
	return result.(int)
}

func (str stringType) ToInt8(val string) int8 {
	result, err := str.getIntegerValueByBitSize(val, BitSize8, true)
	if err != nil {
		panic(err)
	}
	return result.(int8)
}

func (str stringType) ToInt16(val string) int16 {
	result, err := str.getIntegerValueByBitSize(val, BitSize16, true)
	if err != nil {
		panic(err)
	}
	return result.(int16)
}

func (str stringType) ToInt32(val string) int32 {
	result, err := str.getIntegerValueByBitSize(val, BitSize32, true)
	if err != nil {
		panic(err)
	}
	return result.(int32)
}

func (str stringType) ToInt64(val string) int64 {
	result, err := str.getIntegerValueByBitSize(val, BitSize64, true)
	if err != nil {
		panic(err)
	}
	return result.(int64)
}

func (str stringType) ToUint(val string) uint {
	sizeInBits := bits.UintSize
	var result interface{}
	var err error
	if sizeInBits == 32 {
		result, err = str.getIntegerValueByBitSize(val, BitSize32, false)
	} else {
		result, err = str.getIntegerValueByBitSize(val, BitSize64, false)
	}
	if err != nil {
		panic(err)
	}
	return result.(uint)
}

func (str stringType) ToUint8(val string) uint8 {
	result, err := str.getIntegerValueByBitSize(val, BitSize8, false)
	if err != nil {
		panic(err)
	}
	return result.(uint8)
}

func (str stringType) ToUint16(val string) uint16 {
	result, err := str.getIntegerValueByBitSize(val, BitSize16, false)
	if err != nil {
		panic(err)
	}
	return result.(uint16)
}

func (str stringType) ToUint32(val string) uint32 {
	result, err := str.getIntegerValueByBitSize(val, BitSize32, false)
	if err != nil {
		panic(err)
	}
	return result.(uint32)
}

func (str stringType) ToUint64(val string) uint64 {
	result, err := str.getIntegerValueByBitSize(val, BitSize64, false)
	if err != nil {
		panic(err)
	}
	return result.(uint64)
}

func (str stringType) ToFloat32(val string) float32 {
	return 0
}

func (str stringType) ToFloat64(val string) float64 {
	return 0
}

func (str stringType) getIntegerValue(strValue string, integer Integer) (resultValue interface{}, err error) {
	var value interface{}
	var signedValue int64
	var unsignedValue uint64
	if integer.IsSigned() {
		signedValue, err = strconv.ParseInt(strValue, 10, 64)
		value = signedValue
	} else {
		unsignedValue, err = strconv.ParseUint(strValue, 10, 64)
		value = unsignedValue
	}
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(r.(string))
		}
	}()
	if integer.Overflow(value) {
		return nil, errors.New("The given value is out of range of the integer type : " + integer.String())
	}
	integerVal := reflect.New(integer.GetGoType())
	if integer.IsSigned() {
		integerVal.SetInt(signedValue)
	} else {
		integerVal.SetUint(unsignedValue)
	}
	resultValue = integerVal.Interface()
	return
}

func (str stringType) getIntegerValueByBitSize(strValue string, bitSize BitSize, isSigned bool) (resultValue interface{}, err error) {
	if BitSize128 == bitSize {
		panic("BitSize does not support 128")
	}
	var signedValue int64
	var unsignedValue uint64
	if isSigned {
		signedValue, err = strconv.ParseInt(strValue, 10, 64)
	} else {
		unsignedValue, err = strconv.ParseUint(strValue, 10, 64)
	}
	if err != nil {
		return nil, err
	}
	overflow := false
	if isSigned {
		if BitSize8 == bitSize && (math.MinInt8 > signedValue || math.MaxInt8 > signedValue) {
			overflow = true
		} else if BitSize16 == bitSize && (math.MinInt16 > signedValue || math.MaxInt16 > signedValue) {
			overflow = true
		} else if BitSize32 == bitSize && (math.MinInt32 > signedValue || math.MaxInt32 > signedValue) {
			overflow = true
		}
	} else {
		if BitSize8 == bitSize && math.MaxUint8 > unsignedValue {
			overflow = true
		} else if BitSize16 == bitSize && math.MaxUint16 > unsignedValue {
			overflow = true
		} else if BitSize32 == bitSize && math.MaxUint32 > unsignedValue {
			overflow = true
		}
	}
	if overflow {
		return nil, errors.New("the given value is out of range of the integer type")
	}
	if isSigned {
		if BitSize8 == bitSize {
			return int8(signedValue), nil
		} else if BitSize16 == bitSize {
			return int16(signedValue), nil
		} else if BitSize32 == bitSize {
			return int32(signedValue), nil
		}
		return signedValue, nil
	} else {
		if BitSize8 == bitSize {
			return uint8(unsignedValue), nil
		} else if BitSize16 == bitSize {
			return uint16(unsignedValue), nil
		} else if BitSize32 == bitSize {
			return uint32(unsignedValue), nil
		}
		return unsignedValue, nil
	}
}

func (str stringType) getFloatValue(strValue string, float Float) (resultValue interface{}, err error) {
	var value float64
	value, err = strconv.ParseFloat(strValue, 64)
	if err != nil {
		return nil, nil
	}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(r.(string))
		}
	}()
	if float.Overflow(value) {
		return nil, errors.New("The given value is out of range of the float type : " + float.String())
	}
	floatValue := reflect.New(float.GetGoType())
	floatValue.SetFloat(value)
	resultValue = floatValue.Interface()
	return
}

func (str stringType) getFloatValueByBitSize(strValue string, bitSize BitSize) (resultValue interface{}, err error) {
	var value float64
	value, err = strconv.ParseFloat(strValue, 64)
	if err != nil {
		return nil, nil
	}
	overflow := false
	if BitSize32 == bitSize {
		// todo
	} else if BitSize64 == bitSize {
		// todo
	} else {
		panic("BitSize supports only 32 and 64")
	}
	if overflow {
		return nil, errors.New("the given value is out of range of the float type")
	}
	// todo
	return value, nil
}