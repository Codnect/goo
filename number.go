package goo

import (
	"fmt"
	"reflect"
	"strconv"
)

type NumberType int

const (
	IntegerType NumberType = iota
	FloatType
	ComplexType
)

type BitSize int

const (
	BitSize8   BitSize = 8
	BitSize16          = 16
	BitSize32          = 32
	BitSize64          = 64
	BitSize128         = 128
)

type Number interface {
	Type
	GetNumberType() NumberType
	GetBitSize() BitSize
	Overflow(val interface{}) bool
}

type Integer interface {
	Number
	IsSigned() bool
}

type SignedInteger struct {
	baseType
}

func newSignedInteger(baseTyp baseType) SignedInteger {
	return SignedInteger{
		baseTyp,
	}
}

func (integer SignedInteger) GetNumberType() NumberType {
	return IntegerType
}

func (integer SignedInteger) GetBitSize() BitSize {
	switch integer.kind {
	case reflect.Int:
		return BitSize32
	case reflect.Int64:
		return BitSize64
	case reflect.Int8:
		return BitSize8
	case reflect.Int16:
		return BitSize16
	}
	panic("this is not a signed integer type")
}

func (integer SignedInteger) IsSigned() bool {
	return true
}

func (integer SignedInteger) Overflow(val interface{}) bool {
	valType := GetType(val)
	if !valType.IsNumber() || IntegerType != valType.(Number).GetNumberType() || !valType.(Integer).IsSigned() {
		panic("Given type is not compatible with signed integer")
	}
	integerValueStr := fmt.Sprintf("%d", val)
	integerValue, err := strconv.ParseInt(integerValueStr, 10, 64)
	if err != nil {
		panic(err)
	}
	return integer.GetGoValue().OverflowInt(integerValue)
}

type UnsignedInteger struct {
	baseType
}

func newUnsignedInteger(baseTyp baseType) UnsignedInteger {
	return UnsignedInteger{
		baseTyp,
	}
}

func (integer UnsignedInteger) GetNumberType() NumberType {
	return IntegerType
}

func (integer UnsignedInteger) GetBitSize() BitSize {
	switch integer.kind {
	case reflect.Uint:
		return BitSize32
	case reflect.Uint64:
		return BitSize64
	case reflect.Uint8:
		return BitSize8
	case reflect.Uint16:
		return BitSize16
	}
	panic("this is not a unsigned integer type")
}

func (integer UnsignedInteger) IsSigned() bool {
	return false
}

func (integer UnsignedInteger) Overflow(val interface{}) bool {
	valType := GetType(val)
	if !valType.IsNumber() || IntegerType != valType.(Number).GetNumberType() || valType.(Integer).IsSigned() {
		panic("Given type is not compatible with unsigned integer")
	}
	integerValueStr := fmt.Sprintf("%d", val)
	integerValue, err := strconv.ParseUint(integerValueStr, 10, 64)
	if err != nil {
		panic(err)
	}
	return integer.GetGoValue().OverflowUint(integerValue)
}

type Float interface {
	Number
}

type BaseFloat struct {
	baseType
}

func newFloat(baseTyp baseType) Float {
	return BaseFloat{
		baseTyp,
	}
}

func (float BaseFloat) GetNumberType() NumberType {
	return FloatType
}

func (float BaseFloat) GetBitSize() BitSize {
	switch float.kind {
	case reflect.Float32:
		return BitSize32
	case reflect.Float64:
		return BitSize64
	}
	panic("this is not a float type")
}

func (float BaseFloat) Overflow(val interface{}) bool {
	valType := GetType(val)
	if !valType.IsNumber() || FloatType != valType.(Number).GetNumberType() {
		panic("Given type is not compatible with float")
	}
	floatValueStr := fmt.Sprintf("%f", val)
	floatValue, err := strconv.ParseFloat(floatValueStr, 64)
	if err != nil {
		panic(err)
	}
	return float.GetGoValue().OverflowFloat(floatValue)
}

type Complex interface {
	Number
}

type BaseComplex struct {
	baseType
}

func newComplex(baseTyp baseType) Complex {
	return BaseComplex{
		baseTyp,
	}
}

func (complex BaseComplex) GetNumberType() NumberType {
	return ComplexType
}

func (complex BaseComplex) GetBitSize() BitSize {
	switch complex.kind {
	case reflect.Complex64:
		return BitSize64
	case reflect.Complex128:
		return BitSize128
	}
	panic("this is not a complex type")
}

func (complex BaseComplex) Overflow(val interface{}) bool {
	panic("It does not support Overflow for now")
}
