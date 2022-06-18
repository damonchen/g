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

func (i Int8) Compare(o Int8) int {
	if i > o {
		return 1
	} else if i < o {
		return -1
	} else {
		return 0
	}
}

func (i Int8) Int8() int8 {
	return int8(i)
}

type Uint8 uint8

func (i Uint8) Equal(o Uint8) bool {
	return i == o
}

func (i Uint8) Compare(o Uint8) int {
	if i > o {
		return 1
	} else if i < o {
		return -1
	} else {
		return 0
	}
}

func (i Uint8) UInt8() uint8 {
	return uint8(i)
}

type Int16 int16

func (i Int16) Equal(o Int16) bool {
	return i == o
}

func (i Int16) Compare(o Int16) int {
	if i > o {
		return 1
	} else if i < o {
		return -1
	} else {
		return 0
	}
}

func (i Int16) Int16() int16 {
	return int16(i)
}

type Uint16 uint16

func (i Uint16) Equal(o Uint16) bool {
	return i == o
}

func (i Uint16) Compare(o Uint16) int {
	if i > o {
		return 1
	} else if i < o {
		return -1
	} else {
		return 0
	}
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

func (i Int32) Compare(o Int32) int {
	if i > o {
		return 1
	} else if i < o {
		return -1
	} else {
		return 0
	}
}

type Uint32 uint32

func (i Uint32) Equal(o Uint32) bool {
	return i == o
}

func (i Uint32) Compare(o Uint32) int {
	if i > o {
		return 1
	} else if i < o {
		return -1
	} else {
		return 0
	}
}

func (i Uint32) Uint32() uint32 {
	return uint32(i)
}

type Int int

func (i Int) Equal(o Int) bool {
	return i == o
}

func (i Int) Compare(o Int) int {
	if i > o {
		return 1
	} else if i < o {
		return -1
	} else {
		return 0
	}
}

func (i Int) Int() int {
	return int(i)
}

type Uint uint

func (i Uint) Equal(o Uint) bool {
	return i == o
}

func (i Uint) Compare(o Uint) int {
	if i > o {
		return 1
	} else if i < o {
		return -1
	} else {
		return 0
	}
}

func (i Uint) Uint() uint {
	return uint(i)
}

type Int64 int64

func (i Int64) Equal(o Int64) bool {
	return i == o
}

func (i Int64) Compare(o Int64) int {
	if i > o {
		return 1
	} else if i < o {
		return -1
	} else {
		return 0
	}
}

func (i Int) Int64() int64 {
	return int64(i)
}

type Uint64 uint64

func (i Uint64) Equal(o Uint64) bool {
	return i == o
}

func (i Uint64) Compare(o Uint64) int {
	if i > o {
		return 1
	} else if i < o {
		return -1
	} else {
		return 0
	}
}

func (i Uint64) Uint64() uint64 {
	return uint64(i)
}

type Float32 float32

func (f Float32) Equal(o Float32) bool {
	return f == o
}

func (f Float32) Float32() float32 {
	return float32(f)
}

func (i Float32) Compare(o Float32) int {
	if i > o {
		return 1
	} else if i < o {
		return -1
	} else {
		return 0
	}
}

type Float64 float64

func (f Float64) Equal(o Float64) bool {
	return f == o
}

func (i Float64) Compare(o Float64) int {
	if i > o {
		return 1
	} else if i < o {
		return -1
	} else {
		return 0
	}
}

func (f Float64) Float64() float64 {
	return float64(f)
}

type Rune rune

func (r Rune) Equal(o Rune) bool {
	return r == o
}

func (r Rune) Compare(o Rune) int {
	if r > o {
		return 1
	} else if r < o {
		return -1
	} else {
		return 0
	}
}

func (s Rune) Rune() rune {
	return rune(s)
}

type String string

func (s String) Equal(o String) bool {
	return s == o
}

func (s String) Compare(o String) int {
	if s == o {
		return 0
	}

	oo := []rune(o)
	ss := []rune(s)

	for index, e := range ss {
		c := oo[index]
		if e > c {
			return 1
		} else {
			return -1
		}
	}
	return 0
}

func (s String) String() string {
	return string(s)
}
