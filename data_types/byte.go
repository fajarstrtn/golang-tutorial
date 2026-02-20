package data_types

import (
	"fmt"
	"unsafe"
)

func getBytes() {
	/*
	 * The byte data type is exactly the same as uint8,
	 * where its size is 1 byte (8 bits).
	 *
	 * Used to present:
	 * 1. Raw binary data
	 * 2. ASCII Characters
	 * 3. File data
	 * 4. Network data
	 *
	 * The reason of using byte instead of uint8 is
	 * because byte makes your code more meaningful. */
	var b byte = 255
	fmt.Printf("%d\n", b)                // Output: 255
	fmt.Printf("%c\n", b)                // Output: ÿ
	fmt.Printf("%T\n", b)                // Output: uint8
	fmt.Printf("%d\n", unsafe.Sizeof(b)) // Output: 1

	// In Go, string is slice of bytes.
	text := "Hello World"
	fmt.Printf("%v\n", text[0]) // Output: 72
	fmt.Printf("%c\n", text[0]) // Output: H

	binary := []byte(text)
	fmt.Printf("%v\n", binary) // Output: [72 101 108 108 111 32 87 111 114 108 100]
}
