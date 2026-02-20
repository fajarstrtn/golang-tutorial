package data_types

import (
	"fmt"
	"unsafe"
)

func getUintptr() {
	/*
	 * uintptr is an unsigned integer that large enough to store pointer,
	 * where its size depends of cpu architecture (4 bytes for 32-bit system,
	 * and 8 bytes for 64-bit system).
	 *
	 * It is primarily used for low-level memory manipulation,
	 * specifically for performing pointer arithmetic,
	 * which is not allowed on standard Go pointers or unsafe.Pointer.
	 *
	 * uintptr treats a memory address as a raw number.
	 * This allows arithmetic operations (adding or subtracting offsets)
	 * that are necessary for detailed memory control.
	 *
	 * uintptr is an integer type capable of storing the bit pattern of any pointer.
	 * It is used in unsafe operations, converting unsafe.Pointer to perform arithmetic,
	 * and back to unsafe.Pointer.
	 *
	 * Variables of uintptr type can store and print memory address values
	 * but doesn't reference objects.
	 * uintptr aims to store memory address as number, not pointer, only number.
	 *
	 * Pointers ensure type safety and prevent garbage collection,
	 * while uintptr allowing low-level memory manipulation
	 * and are not type-safe.
	 *
	 * Use uintptr in advanced, performance-critical, or unsafe operations. */
	n := 100
	ptr := unsafe.Pointer(uintptr(unsafe.Pointer(&n)) + 8)
	addr := uintptr(ptr)
	fmt.Printf("%v\n", n)    // Output: 100
	fmt.Printf("%v\n", ptr)  // Output: 0xc00000a0c8
	fmt.Printf("%v\n", addr) // Output: 824633761992
}
