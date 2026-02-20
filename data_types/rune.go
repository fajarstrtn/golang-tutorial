package data_types

import (
	"fmt"
	"unsafe"
)

func getRune() {
	/*
	 * The rune data type is exactly the same as int32, where its size is 4 bytes.
	 *
	 * Used to represent unicode characters (not ASCII only, all languages).
	 * For example:
	 * 1. Latin   : A
	 * 2. Japanese: あ
	 * 3. Emoji   : 😀
	 *
	 * Use byte for raw data and ASCII, rune for character and unicode. */
	var char rune = 'A'
	fmt.Printf("%d\n", char)                // Output: 65
	fmt.Printf("%c\n", char)                // Output: A
	fmt.Printf("%T\n", char)                // Output: int32
	fmt.Printf("%d\n", unsafe.Sizeof(char)) // Output: 4

	en := "A"
	fmt.Printf("%v\n", len(en)) // Output: 1

	// rune exists because UTF-8 characters can be multiple bytes. あ stores 3 bytes.
	jp := "あ"
	fmt.Printf("%v\n", len(jp)) // Output: 3

	// But rune counts correctly; In Go, string indexing gives byte, not character.
	for _, r := range jp {
		fmt.Println(r)        // Output: 12354
		fmt.Printf("%c\n", r) // Output: あ
	}

	var emoji rune = '😀'
	fmt.Println(emoji)        // Output: 128512
	fmt.Printf("%c\n", emoji) // Output: 😀

	str := "Aあ"
	for _, r := range str {
		fmt.Printf("%v\n", string(r)) // Output: A\nあ
	}

	// A for 1 byte and あ for 3 bytes.
	fmt.Println(len(str)) // Output: 4
}
