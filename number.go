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

type signedInteger struct {
	baseType
}

func newSignedInteger(baseTyp baseType) signedInteger {
	return signedInteger{
		baseTyp,
	}
}

func (integer signedInteger) GetNumberType() NumberType {
	return IntegerType
}

func (integer signedInteger) GetBitSize() BitSize {
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

func (integer signedInteger) IsSigned() bool {
	return true
}

func (integer signedInteger) Overflow(val interface{}) bool {
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

type unsignedInteger struct {
	baseType
}

func newUnsignedInteger(baseTyp baseType) unsignedInteger {
	return unsignedInteger{
		baseTyp,
	}
}

func (integer unsignedInteger) GetNumberType() NumberType {
	return IntegerType
}

func (integer unsignedInteger) GetBitSize() BitSize {
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

func (integer unsignedInteger) IsSigned() bool {
	return false
}

func (integer unsignedInteger) Overflow(val interface{}) bool {
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

type baseFloat struct {
	baseType
}

func newFloat(baseTyp baseType) Float {
	return baseFloat{
		baseTyp,
	}
}

func (float baseFloat) GetNumberType() NumberType {
	return FloatType
}

func (float baseFloat) GetBitSize() BitSize {
	switch float.kind {
	case reflect.Float32:
		return BitSize32
	case reflect.Float64:
		return BitSize64
	}
	panic("this is not a float type")
}

func (float baseFloat) Overflow(val interface{}) bool {
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

type baseComplex struct {
	baseType
}

func newComplex(baseTyp baseType) Complex {
	return baseComplex{
		baseTyp,
	}
}

func (complex baseComplex) GetNumberType() NumberType {
	return ComplexType
}

func (complex baseComplex) GetBitSize() BitSize {
	switch complex.kind {
	case reflect.Complex64:
		return BitSize64
	case reflect.Complex128:
		return BitSize128
	}
	panic("this is not a complex type")
}

func (complex baseComplex) Overflow(val interface{}) bool {
	panic("It does not support Overflow for now")
}
