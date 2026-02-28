package data_types

import (
	"fmt"
	"unsafe"
)

const MESSAGE_TEMPLATE = "Rune %d is '%c' (Unicode: U+%04X)\n"

func getRune() {
	/*
	 * Go has rune data type which is comparable to "char" type in java.
	 * It represents a single Unicode code point, meaning it represents a character
	 * in its entirety from Unicode character set.
	 *
	 * If you convert a string to slice of rune, you can iterate through each character
	 * which is just what you need for string manipulation.
	 *
	 * The rune data type is exactly the same as int32, where its size is 4 bytes.
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

	// rune exists because UTF-8 characters can be multiple bytes, あ stores 3 bytes.
	jp := "あ"
	fmt.Printf("%v\n", len(jp)) // Output: 3

	/*
	 * Finally, Go provides a nifty way of iterating through string
	 * without caring for these details when you use range.
	 * In case of range knows the internals of string and
	 * implicitly converts string to []rune.
	 * As a result, below code prints the correct
	 * Unicode codepoint for the character 'val' at any index 'idx'.
	 *
	 * Output:
	 * Rune 0 is 'F' (Unicode: U+0046)
	 * Rune 1 is 'l' (Unicode: U+006C)
	 * Rune 2 is 'a' (Unicode: U+0061)
	 * Rune 3 is 'b' (Unicode: U+0062)
	 * Rune 4 is 'b' (Unicode: U+0062)
	 * Rune 5 is 'e' (Unicode: U+0065)
	 * Rune 6 is 'r' (Unicode: U+0072)
	 * Rune 7 is 'g' (Unicode: U+0067)
	 * Rune 8 is 'a' (Unicode: U+0061)
	 * Rune 9 is 's' (Unicode: U+0073)
	 * Rune 10 is 't' (Unicode: U+0074)
	 * Rune 11 is 'e' (Unicode: U+0065)
	 * Rune 12 is 'd' (Unicode: U+0064) */
	txt := "Flabbergasted"
	iterateString(txt)

	/*
	 * But rune counts correctly.
	 * In Go, string indexing gives byte, not character. */
	for _, r := range jp {
		fmt.Printf("%d\n", r) // Output: 12354
		fmt.Printf("%c\n", r) // Output: あ
	}

	// A for 1 byte and あ for 3 bytes.
	str := "Aあ"
	fmt.Println(len(str)) // Output: 4
<<<<<<< HEAD
<<<<<<< HEAD

	/*
	 * Output:
	 * A
	 * 65
	 * int32
	 * あ
	 * 12354
	 * int32 */
	for _, r := range str {
		fmt.Printf("%v\n", string(r))
		fmt.Printf("%v\n", r)
		fmt.Printf("%T\n", r)

	var emoji rune = '😀'
	fmt.Println(emoji)        // Output: 128512
	fmt.Printf("%c\n", emoji) // Output: 😀

	txt = "Hello, 世界"
	runes := []rune(txt)
	fmt.Printf("Length in bytes: %d\n", len(txt))   // Output: Length in bytes: 13
	fmt.Printf("Length in runes: %d\n", len(runes)) // Output: Length in runes: 9

	/*
	 * Output:
	 * Rune 0 is 'H' (Unicode: U+0048)
	 * Rune 1 is 'e' (Unicode: U+0065)
	 * Rune 2 is 'l' (Unicode: U+006C)
	 * Rune 3 is 'l' (Unicode: U+006C)
	 * Rune 4 is 'o' (Unicode: U+006F)
	 * Rune 5 is ',' (Unicode: U+002C)
	 * Rune 6 is ' ' (Unicode: U+0020)
	 * Rune 7 is '世' (Unicode: U+4E16)
	 * Rune 8 is '界' (Unicode: U+754C) */
	iterateRunes(runes)
}

func iterateString(txt string) {
	/*
	 * For simple iteration without in-place modification,
	 * using for...range on the string directly is generally preferred
	 * as it avoids memory allocation for a new []rune slice.
	 * This produces the exact same rune values as ranging over a []rune slice. */
	for i, r := range txt {
		fmt.Printf(MESSAGE_TEMPLATE, i, r, r)
	}
}

func iterateRunes(runes []rune) {
	for i, r := range runes {
		fmt.Printf(MESSAGE_TEMPLATE, i, r, r)
	}
}
