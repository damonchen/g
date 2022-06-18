package g

type Bool bool

func (b Bool) Equal(o Bool) bool {
	return b == o
}

func (b Bool) Bool() bool {
	return bool(b)
}

type Byte byte

func (i Byte) Equal(o Byte) bool {
	return i == o
}

func (i Byte) Byte() byte {
	return byte(i)
}

type Int8 int8

func (i Int8) Equal(o Int8) bool {
	return i == o
}

func (i Int8) Int8() int8 {
	return int8(i)
}

type Uint8 uint8

func (i Uint8) Equal(o Uint8) bool {
	return i == o
}

func (i Uint8) UInt8() uint8 {
	return uint8(i)
}

type Int16 int16

func (i Int16) Equal(o Int16) bool {
	return i == o
}

func (i Int16) Int16() int16 {
	return int16(i)
}

type Uint16 uint16

func (i Uint16) Equal(o Uint16) bool {
	return i == o
}

func (i Uint16) UInt16() uint16 {
	return uint16(i)
}

type Int32 int32

func (i Int32) Equal(o Int32) bool {
	return i == o
}

func (i Int32) Int32() int32 {
	return int32(i)
}

type Uint32 uint32

func (i Uint32) Equal(o Uint32) bool {
	return i == o
}

func (i Uint32) Uint32() uint32 {
	return uint32(i)
}

type Int int

func (i Int) Equal(o Int) bool {
	return i == o
}

func (i Int) Int() int {
	return int(i)
}

type Uint uint

func (i Uint) Equal(o Uint) bool {
	return i == o
}

func (i Uint) Uint() uint {
	return uint(i)
}

type Int64 int64

func (i Int64) Equal(o Int64) bool {
	return i == o
}

func (i Int) Int64() int64 {
	return int64(i)
}

type Uint64 uint64

func (i Uint64) Equal(o Uint64) bool {
	return i == o
}

func (i Int) Uint64() uint64 {
	return uint64(i)
}

type Float32 float32

func (f Float32) Equal(o Float32) bool {
	return f == o
}

func (f Float32) Float32() float32 {
	return float32(f)
}

type Float64 float64

func (f Float64) Equal(o Float64) bool {
	return f == o
}

func (f Float64) Float64() float64 {
	return float64(f)
}

type Rune rune

func (s Rune) Equal(o Rune) bool {
	return s == o
}

func (s Rune) Rune() rune {
	return rune(s)
}

type String string

func (s String) Equal(o String) bool {
	return s == o
}

func (s String) String() string {
	return string(s)
}
