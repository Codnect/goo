package goo

import (
	"fmt"
	"math"
	"math/bits"
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
	BitSize16  BitSize = 16
	BitSize32  BitSize = 32
	BitSize64  BitSize = 64
	BitSize128 BitSize = 128
)

type Number interface {
	Type
	Instantiable
	GetType() NumberType
	GetBitSize() BitSize
	Overflow(val interface{}) bool
	ToString(val interface{}) string
}

type Integer interface {
	Number
	IsSigned() bool
}

type signedIntegerType struct {
	*baseType
}

func newSignedIntegerType(baseTyp *baseType) signedIntegerType {
	return signedIntegerType{
		baseTyp,
	}
}

func (integer signedIntegerType) GetType() NumberType {
	return IntegerType
}

func (integer signedIntegerType) GetBitSize() BitSize {
	switch integer.kind {
	case reflect.Int64:
		return BitSize64
	case reflect.Int8:
		return BitSize8
	case reflect.Int16:
		return BitSize16
	case reflect.Int32:
		return BitSize32
	default:
		if bits.UintSize == 32 {
			return BitSize32
		}
		return BitSize64
	}
}

func (integer signedIntegerType) IsSigned() bool {
	return true
}

func (integer signedIntegerType) Overflow(val interface{}) bool {
	valType := GetType(val)
	if !valType.IsNumber() || IntegerType != valType.(Number).GetType() || !valType.(Integer).IsSigned() {
		panic("Given type is not compatible with signed integer")
	}
	integerValueStr := fmt.Sprintf("%d", val)
	integerValue, err := strconv.ParseInt(integerValueStr, 10, 64)
	if err != nil {
		panic(err)
	}

	bitSize := integer.GetBitSize()
	overflow := false
	if BitSize8 == bitSize && (math.MinInt8 > integerValue || math.MaxInt8 < integerValue) {
		overflow = true
	} else if BitSize16 == bitSize && (math.MinInt16 > integerValue || math.MaxInt16 < integerValue) {
		overflow = true
	} else if BitSize32 == bitSize && (math.MinInt32 > integerValue || math.MaxInt32 < integerValue) {
		overflow = true
	}
	return overflow
}

func (integer signedIntegerType) NewInstance() interface{} {
	return reflect.New(integer.GetGoType()).Interface()
}

func (integer signedIntegerType) ToString(val interface{}) string {
	valType := GetType(val)
	if !valType.IsNumber() || IntegerType != valType.(Number).GetType() || !valType.(Integer).IsSigned() {
		panic("Incompatible type : " + valType.GetName())
	}
	return fmt.Sprintf("%d", val)
}

type unsignedIntegerType struct {
	*baseType
}

func newUnsignedIntegerType(baseTyp *baseType) unsignedIntegerType {
	return unsignedIntegerType{
		baseTyp,
	}
}

func (integer unsignedIntegerType) GetType() NumberType {
	return IntegerType
}

func (integer unsignedIntegerType) GetBitSize() BitSize {
	switch integer.kind {
	case reflect.Uint64:
		return BitSize64
	case reflect.Uint8:
		return BitSize8
	case reflect.Uint16:
		return BitSize16
	case reflect.Uint32:
		return BitSize32
	default:
		if bits.UintSize == 32 {
			return BitSize32
		}
		return BitSize64
	}
}

func (integer unsignedIntegerType) IsSigned() bool {
	return false
}

func (integer unsignedIntegerType) Overflow(val interface{}) bool {
	valType := GetType(val)
	if !valType.IsNumber() || IntegerType != valType.(Number).GetType() || valType.(Integer).IsSigned() {
		panic("Given type is not compatible with unsigned integer")
	}
	integerValueStr := fmt.Sprintf("%d", val)
	integerValue, err := strconv.ParseUint(integerValueStr, 10, 64)
	if err != nil {
		panic(err)
	}

	bitSize := integer.GetBitSize()
	overflow := false
	if BitSize8 == bitSize && math.MaxUint8 < integerValue {
		overflow = true
	} else if BitSize16 == bitSize && math.MaxUint16 < integerValue {
		overflow = true
	} else if BitSize32 == bitSize && math.MaxUint32 < integerValue {
		overflow = true
	}
	return overflow
}

func (integer unsignedIntegerType) NewInstance() interface{} {
	return reflect.New(integer.GetGoType()).Interface()
}

func (integer unsignedIntegerType) ToString(val interface{}) string {
	valType := GetType(val)
	if !valType.IsNumber() || IntegerType != valType.(Number).GetType() || valType.(Integer).IsSigned() {
		panic("Incompatible type : " + valType.GetName())
	}
	return fmt.Sprintf("%d", val)
}

type Float interface {
	Number
}

type floatType struct {
	*baseType
}

func newFloatType(baseTyp *baseType) Float {
	return floatType{
		baseTyp,
	}
}

func (float floatType) GetType() NumberType {
	return FloatType
}

func (float floatType) GetBitSize() BitSize {
	switch float.kind {
	case reflect.Float32:
		return BitSize32
	default:
		return BitSize64
	}
}

func (float floatType) Overflow(val interface{}) bool {
	valType := GetType(val)
	if !valType.IsNumber() || FloatType != valType.(Number).GetType() {
		panic("Given type is not compatible with float")
	}
	floatValueStr := fmt.Sprintf("%f", val)
	floatValue, err := strconv.ParseFloat(floatValueStr, 64)
	if err != nil {
		panic(err)
	}

	bitSize := float.GetBitSize()
	overflow := false
	if BitSize32 == bitSize && math.MaxFloat32 < floatValue {
		overflow = true
	} else if BitSize64 == bitSize && math.MaxFloat64 < floatValue {
		overflow = true
	}
	return overflow
}

func (float floatType) NewInstance() interface{} {
	return reflect.New(float.GetGoType()).Interface()
}

func (float floatType) ToString(val interface{}) string {
	valType := GetType(val)
	if !valType.IsNumber() || FloatType != valType.(Number).GetType() {
		panic("Incompatible type : " + valType.GetName())
	}
	return fmt.Sprintf("%f", val)
}

type Complex interface {
	Number
	GetImaginaryData(val interface{}) interface{}
	GetRealData(val interface{}) interface{}
}

type complexType struct {
	*baseType
}

func newComplexType(baseTyp *baseType) Complex {
	return complexType{
		baseTyp,
	}
}

func (complex complexType) GetType() NumberType {
	return ComplexType
}

func (complex complexType) GetBitSize() BitSize {
	switch complex.kind {
	case reflect.Complex64:
		return BitSize64
	default:
		return BitSize128
	}
}

func (complex complexType) Overflow(val interface{}) bool {
	panic("It does not support Overflow for now")
}

func (complex complexType) GetImaginaryData(val interface{}) interface{} {
	valType := GetType(val)
	if !valType.IsNumber() || ComplexType != valType.(Number).GetType() {
		panic("Given type is not compatible with complex")
	}

	if complex.GetBitSize() == BitSize64 {
		return imag(val.(complex64))
	}
	return imag(val.(complex128))
}

func (complex complexType) GetRealData(val interface{}) interface{} {
	valType := GetType(val)
	if !valType.IsNumber() || ComplexType != valType.(Number).GetType() {
		panic("Given type is not compatible with complex")
	}

	if complex.GetBitSize() == BitSize64 {
		return real(val.(complex64))
	}
	return real(val.(complex128))
}

func (complex complexType) NewInstance() interface{} {
	return reflect.New(complex.GetGoType()).Interface()
}

func (complex complexType) ToString(val interface{}) string {
	valType := GetType(val)
	if !valType.IsNumber() || ComplexType != valType.(Number).GetType() {
		panic("Incompatible type : " + valType.GetName())
	}
	return fmt.Sprintf("%f", val)
}
