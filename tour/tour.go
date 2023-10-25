package tour

import (
	"fmt"
	"math"
)

var (
	type_b          bool
	type_string     string
	type_int        int  //32 or 64 and types int8  int16  int32  int64
	type_uint       uint //32 or 64 and types uint8  uint16  uint32  uint64
	type_uintptr    uintptr
	type_byte       byte // alias for uint8
	type_rune       rune // alias for int32 represents a Unicode code point
	type_float32    float32
	type_float64    float64
	type_complex64  complex64
	type_complex128 complex128
)

func IntMinMax() {
	i32 := "int"
	u32 := "uint"
	i64 := ""
	u64 := ""
	if math.MaxInt == math.MaxInt64 {
		i32 = ""
		u32 = ""
		i64 = "int"
		u64 = "uint"
	}
	// int
	fmt.Printf("Integer types min and max:\n")
	// int
	fmt.Printf("  int8: %21d .. %d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf(" int16: %21d .. %d\n", math.MinInt16, math.MaxInt16)
	fmt.Printf(" int32: %21d .. %d %s alias rune\n", math.MinInt32, math.MaxInt32, i32)
	fmt.Printf(" int64: %21d .. %d %s\n", math.MinInt64, math.MaxInt64, i64)
	// unsigned
	fmt.Printf(" uint8: %21d .. %d alias byte\n", 0, math.MaxUint8)
	fmt.Printf("uint16: %21d .. %d\n", 0, math.MaxUint16)
	fmt.Printf("uint32: %21d .. %d %s\n", 0, math.MaxUint32, u32)
	fmt.Printf("uint64: %21d .. %d %s\n", 0, uint64(math.MaxUint64), u64)
}

func SystemInt() {
	systemInt := 32
	if math.MaxInt == math.MaxInt64 {
		systemInt = 64
	}
	fmt.Printf("This is %d bit system.\n", systemInt)
}
