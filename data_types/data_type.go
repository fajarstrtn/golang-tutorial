/*
 * The data types specify the types of data that variables can hold,
 * divided into four categories which are as follows:
 * 1. Basic type    : Numbers, strings, and booleans come under this category.
 * 2. Aggregate type: Array and structs come under this category.
 * 3. Reference type: Pointers, slices, maps, functions, and channels come under this category.
 * 4. Interface type
 *
 * The basic data types are further categorized into three subcategories which are:
 * 1. Numbers
 * 2. Strings
 * 3. Booleans */
package data_types

/*
 * Numbers are divided into three sub-categories that are:
 * 1. Integers              : Both signed (can store both positive and negative values)
 * and unsigned (can only store non-negative values) integers
 * are available in four different sizes.
 * 2. Floating-Point Numbers: Used to store positive and negative numbers with a decimal point.
 * 3. Complex Numbers       : The in-built function creates a complex number
 * from its imaginary and real part and in-built imaginary and real function extract those parts.
 * float32 and float64 are also part of these complex numbers. */
func GenerateNumbers() {
	/*
	 * These are five types of signed integers:
	 * 1. int  : Depends on platform (32 bits in 32-bit system and 64 bits in 64-bit system),
	 * Range -2147483648 to 2147483647 in 32-bit system
	 * and -9223372036854775808 to 9223372036854775807 in 64-bit system.
	 * 2. int8 : 8-bit signed integer, Range -128 to 127.
	 * 3. int16: 16-bit signed integer, Range -32768 to 32767.
	 * 4. int32: 32-bit signed integer, Range -2147483648 to 2147483647.
	 * 5. int64: 64-bit signed integer, Range -9223372036854775808 to 9223372036854775807. */
	getSignedIntegers()

	/*
	 * These are five types of unsigned integers:
	 * 1. uint  : Depends on platform (32 bits in 32-bit system and 64 bits in 64-bit system),
	 * Range 0 to 4294967295 in 32-bit system
	 * and 0 to 18446744073709551615 in 64-bit system.
	 * 2. uint8 : 8-bit unsigned integer, Range 0 to 255.
	 * 3. uint16: 16-bit unsigned integer, Range 0 to 65535.
	 * 4. uint32: 32-bit unsigned integer, Range 0 to 4294967295.
	 * 5. uint64: 64-bit unsigned integer, Range 0 to 18446744073709551615. */
	getUnsignedIntegers()

	/*
	 * There are also other special built-in integer data types:
	 * 1. byte   : Synonym of uint8, Range 0 to 255.
	 * 2. rune   : Synonym of int32 and also represent Unicode code points,
	 * Range -2147483648 to 2147483647 (full range) and 0 to 1114111 (unicode range).
	 * 3. uintptr: An unsigned integer type; Its width is not defined
	 * (depends on your system architecture) but its can hold all the bits of a pointer value,
	 * Range 0 to 4,294,967,295 (32-bit system).
	 *
	 * They are actually aliases for other integer types,
	 * but each has a specific purpose and meaning.
	 *
	 * Understanding them is very important,
	 * especially when dealing with text,
	 * binary data, and low-level programming. */
	getBytes()
	getRune()
	getUintptr()

	/*
	 * There are two types of complex numbers:
	 * 1. complex64 : Contains float32 as a real and imaginary component.
	 * 2. complex128: Contains float64 as a real and imaginary component. */
	getComplexNumbers()

	/*
	 * There are two types of floating-point data types:
	 * 1. float32: 32-bit IEEE 754 floating-point number, Range -3.4e+38 to 3.4e+38.
	 * 2. float64: 64-bit IEEE 754 floating-point number, Range -1.7e+308 to +1.7e+308.
	 *
	 * Three literal styles are available:
	 * 1. decimal    : 3.15
	 * 2. exponential: 12e18 or 3E10
	 * 3. mixed      : 13.16e12 */
	getFloatNumbers()
}

/*
 * String a sequence of variable-width characters where
 * every character is represented by one or more bytes using UTF-8 Encoding.
 *
 * In other words, strings are the immutable chain of arbitrary bytes
 * (including bytes with zero value) or string is a read-only slice of bytes
 * and the bytes of the strings can be represented
 * in the Unicode text using UTF-8 encoding.
 *
 * Due to UTF-8 encoding, Go string can contain a text
 * which is the mixture of any language present in the world,
 * without any confusion and limitation of the page.
 *
 * Generally, strings are enclosed in double-quotes (""). */
func GenerateStrings() {
	getStrings()
}

/*
 * The boolean data type is declared with the bool keyword.
 * It represents only one bit of information either true or false.
 *
 * The values of type boolean are not converted
 * implicitly or explicitly to any other type.
 * The default value of a boolean data type is false.
 *
 * Boolean values are mostly used for conditional testing. */
func GenerateBooleans() {
	getBooleans()
}
