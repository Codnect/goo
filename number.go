package goo

import "reflect"

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

func (integer UnsignedInteger) GetByteSize() BitSize {
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
