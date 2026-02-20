package data_types

import (
	"fmt"
	"unsafe"
)

func getBooleans() {
	var bool1 bool = true // Typed declaration with initial value
	var bool2 = true      // Untyped declaration with initial value
	var bool3 bool        // Typed declaration without initial value
	bool4 := false        // Untyped declaration with initial value

	fmt.Printf("bool1? %t\n", bool1) // Output: bool1? true
	fmt.Printf("bool2? %t\n", bool2) // Output: bool2? true
	fmt.Printf("bool3? %t\n", bool3) // Output: bool3? false
	fmt.Printf("bool4? %t\n", bool4) // Output: bool4? false

	result1 := bool1 == bool2
	result2 := bool2 == bool3
	result3 := bool3 == bool4
	result4 := bool4 == bool1

	fmt.Printf("bool1 == bool2? %t\n", result1) // Output: bool1 == bool2? true
	fmt.Printf("bool2 == bool3? %t\n", result2) // Output: bool2 == bool3? false
	fmt.Printf("bool3 == bool4? %t\n", result3) // Output: bool3 == bool4? true
	fmt.Printf("bool4 == bool1? %t\n", result4) // Output: bool4 == bool1? false

	fmt.Printf("%d\n", unsafe.Sizeof(result1)) // Output: 1
	fmt.Printf("%d\n", unsafe.Sizeof(result2)) // Output: 1
	fmt.Printf("%d\n", unsafe.Sizeof(result3)) // Output: 1
	fmt.Printf("%d\n", unsafe.Sizeof(result4)) // Output: 1
}
