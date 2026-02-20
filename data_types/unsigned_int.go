package data_types

import (
	"fmt"
	"unsafe"
)

func getUnsignedIntegers() {
	var a uint8 = 8
	fmt.Printf("%d\n", a) // Output: 8
	fmt.Printf("%T\n", a) // Output: uint8

	var b uint16 = 16_000
	fmt.Printf("%d\n", b) // Output: 16000
	fmt.Printf("%T\n", b) // Output: uint16

	var c uint32 = 32_000_000
	fmt.Printf("%d\n", c) // Output: 32000000
	fmt.Printf("%T\n", c) // Output: uint32

	var d uint64 = 64_000_000_000
	fmt.Printf("%d\n", d) // Output: 64000000000
	fmt.Printf("%T\n", d) // Output: uint64

	var e uint8 = 215
	var f uint16 = uint16(e)

	fmt.Printf("%d\n", f) // Output: 215
	fmt.Printf("%T\n", f) // Output: uint16

	var g, h uint = 1200, 1800

	fmt.Printf("Size of g: %v\n", unsafe.Sizeof(g)) // Output: Size of g: 8
	fmt.Printf("Size of h: %v\n", unsafe.Sizeof(h)) // Output: Size of h: 8

	fmt.Printf("g + h = %v\n", (g + h))  // Output: g + h = 3000
	fmt.Printf("g - h = %v\n", (g - h))  // Output: g - h = 18446744073709551016
	fmt.Printf("g * h = %v\n", (g * h))  // Output: g * h = 2160000
	fmt.Printf("g / h = %v\n", (g / h))  // Output: g / h = 0
	fmt.Printf("g %% h = %v\n", (g % h)) // Output: g % h = 1200
}
