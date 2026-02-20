package data_types

import (
	"fmt"
	"unsafe"
)

func getComplexNumbers() {
	var a complex64 = complex(5, 2)
	var b complex128 = complex(9, 7)

	fmt.Printf("%v\n", a)                // Output: (5+2i)
	fmt.Printf("%T\n", a)                // Output: complex64
	fmt.Printf("%d\n", unsafe.Sizeof(a)) // Output: 8

	fmt.Printf("%v\n", b)                // Output: (9+7i)
	fmt.Printf("%T\n", b)                // Output: complex128
	fmt.Printf("%d\n", unsafe.Sizeof(b)) // Output: 16

	/*
	 * There are few built-in functions in complex numbers:
	 * 1. complex(): Make complex numbers from two floats.
	 * 2. real()   : Get real part of the input complex number as float64.
	 * 3. imag()   : Get imaginary of the input complex number part as float64. */
	val := 10 + 5i
	realNumber := real(val)
	imagNumber := imag(val)

	fmt.Printf("%v\n", val)        // Output: (10+5i)
	fmt.Printf("%T\n", val)        // Output: complex128
	fmt.Printf("%v\n", realNumber) // Output: 10
	fmt.Printf("%T\n", realNumber) // Output: float64
	fmt.Printf("%v\n", imagNumber) // Output: 5
	fmt.Printf("%T\n", imagNumber) // Output: float64
}
