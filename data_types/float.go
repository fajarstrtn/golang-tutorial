package data_types

import (
	"fmt"
	"unsafe"
)

func getFloatNumbers() {
	var a float32 = 10.52
	fmt.Printf("%f\n", a)                // Output: 10.520000
	fmt.Printf("%T\n", a)                // Output: float32
	fmt.Printf("%d\n", unsafe.Sizeof(a)) // Output: 4

	// The float64 data type can store a larger set of numbers than float32.
	var b float64 = 125.23e-5
	fmt.Printf("%f\n", b)                // Output: 0.001252
	fmt.Printf("%T\n", b)                // Output: float64
	fmt.Printf("%d\n", unsafe.Sizeof(b)) // Output: 8

	/*
	 * The default type for float is float64.
	 * If you do not specify a type, the type will be float64. */
	c := 33.4571e2
	fmt.Printf("%f\n", c)                // Output: 3345.710000
	fmt.Printf("%T\n", c)                // Output: float64
	fmt.Printf("%d\n", unsafe.Sizeof(c)) // Output: 8

	var d, e float32 = 7.49, 12.112

	fmt.Printf("d + e = %f\n", (d + e)) // Output: d + e = 19.602000
	fmt.Printf("d - e = %f\n", (d - e)) // Output: d - e = -4.622000
	fmt.Printf("d * e = %f\n", (d * e)) // Output: d * e = 90.718880
	fmt.Printf("d / e = %f\n", (d / e)) // Output: d / e = 0.618395

	f, g := 5.25, 14.766

	fmt.Printf("f + g = %.2f\n", (f + g)) // Output: f + g = 20.02
	fmt.Printf("f - g = %.2f\n", (f - g)) // Output: f - g = -9.52
	fmt.Printf("f * g = %.2f\n", (f * g)) // Output: f * g = 77.52
	fmt.Printf("f / g = %.2f\n", (f / g)) // Output: f / g = 0.36
}
