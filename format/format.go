package format

import (
	"fmt"
	"log"
)

var (
	message1 string = "Hello, John Doe!"
	message2 string = "Welcome to Old Trafford!"
)

func PrintSomething() {
	/*
	 * The Print() function prints its arguments with their default format.
	 * It prints without a newline and doesn't add space automatically.
	 *
	 * When to use:
	 * 1. You want full control over spacing and line breaks.
	 * 2. Low-level or very simple output. */
	fmt.Print(message1)
	fmt.Print(message2)

	// You have to manage spaces and newlines yourself.
	fmt.Print(message1)
	fmt.Print(message2 + "\n")

	/*
	 * If we want to print the arguments in new lines,
	 * we need to use \n (it creates new lines). */
	fmt.Print(message1, "\n")
	fmt.Print(message2, "\n")

	// It is also possible to only use one Print() for printing multiple variables.
	fmt.Print(message1, "\n", message2)

	// If we want to add a space between string arguments, we need to use " ".
	fmt.Print(message1, " ", message2, "\n")

	x, y := 10, 20

	// Print() inserts a space between the arguments if neither are strings.
	fmt.Print(x, y)
	fmt.Print("\n")
}

func PrintSomethingWithNewLine() {
	/*
	 * fmt.Println() prints with a newline at the end
	 * and automatically adds spaces between arguments.
	 * It converts values to readable text automatically.
	 *
	 * When to use:
	 * 1. Most common choice.
	 * 2. Quick logging.
	 * 3. Debugging values.
	 *
	 * If you're not sure which one to use, use Println.
	 *
	 * When you do fmt.Println(x):
	 * 1. Detects type
	 * 2. Applies default formatting
	 * 3. Converts to string
	 * 4. Writes to stdout. */
	fmt.Println(message1, message2)
	fmt.Println("Hello", "World")
	fmt.Println(10, "Hello", true)
}

func PrintSomethingWithFormattingVerbs() {
	x, y, text, t := 10, 3.141, "Hello World", true

	/*
	 * The Printf() function first formats its argument based on
	 * the given formatting verb and then prints them.
	 * It gives you precise control over how data looks
	 * and does not add a newline automatically.
	 *
	 * Go offers several formatting verbs that can be used with the Printf() function.
	 *
	 * Format verbs start with %.
	 *
	 * The following verbs can be used with all data types:
	 * 1. %v   : Prints the value in the default format.
	 * 2. %+v  : Prints the value in its default format with the field names
	 * included when dealing specifically with structs.
	 * 3. %#v  : Prints the value in Go-syntax format.
	 * 4. %T   : Prints the type of the value.
	 * 5. %% : Prints the % sign.
	 *
	 * The + or # symbol is a flag that modifies the behavior of the general %v verb.
	 *
	 * It prints using a format string. */
	fmt.Printf("%v\n", x)     // Output: 10
	fmt.Printf("%#v\n", x)    // Output: 10
	fmt.Printf("%T\n", x)     // Output: int
	fmt.Printf("%v%%\n", x)   // Output: 10%
	fmt.Printf("%v\n", text)  // Output: Hello World
	fmt.Printf("%#v\n", text) // Output: "Hello World"
	fmt.Printf("%T\n", text)  // Output: string

	type User struct {
		Name string
		Age  int
	}

	user := User{"John Doe", 20}

	// Using %v is the same as "use Golang's default rule".
	fmt.Printf("%v\n", user)

	// It's Go-syntax representation (great for debugging).
	fmt.Printf("%#v\n", user)

	/*
	 * The primary use of %+v is to display the field names
	 * along with their corresponding values,
	 * which is extremely useful for debugging.
	 * For basic types (e.g., integers, strings, etc.),
	 * the + flag generally has no effect,
	 * and the output is the same as %v.
	 *
	 * Use %+v while learning structs. */
	fmt.Printf("%+v\n", user)

	/*
	 * It works with any type and makes %v perfect for debugging.
	 * Works automatically. No loops needed. */
	fmt.Printf("%v\n", []int{1, 2, 3})         // Output: [1 2 3]
	fmt.Printf("%v\n", map[string]int{"a": 1}) // Output: map[a:1]

	/*
	 * The following verbs can be used with the integer data type:
	 * 1. %b   : Base 2.
	 * 2. %d   : base 10.
	 * 3. %+d  : Base 10 and always show sign.
	 * 4. %o   : Base 8.
	 * 5. %O   : Base 8 with leading 0o.
	 * 6. %x   : Base 16 lowercase.
	 * 7. %X   : Base 16 uppercase.
	 * 8. %#x  : Base 16 with leading 0x.
	 * 9. %4d  : Pad with spaces (width 4, right justified).
	 * 10. %-4d: Pad with spaces (width 4, left justified).
	 * 11. %04d: Pad with zeroes (width 4). */
	fmt.Printf("%b\n", x)   // Output: 1010
	fmt.Printf("%d\n", x)   // Output: 10
	fmt.Printf("%+d\n", x)  // Output: +10
	fmt.Printf("%o\n", x)   // Output: 12
	fmt.Printf("%O\n", x)   // Output: 0o12
	fmt.Printf("%x\n", x)   // Output: a
	fmt.Printf("%X\n", x)   // Output: A
	fmt.Printf("%#X\n", x)  // Output: 0XA
	fmt.Printf("%5d\n", x)  // Output:    10
	fmt.Printf("%-5d\n", x) // Output: 10
	fmt.Printf("%05d\n", x) // Output: 00010

	/*
	 * The following verbs can be used with the string data type:
	 * 1. %s  : Prints the value as plain string.
	 * 2. %q  : Prints the value as a double-quoted string.
	 * 3. %8s : Prints the value as plain string (width 8, right justified).
	 * 4. %-8s: Prints the value as plain string (width 8, left justified).
	 * 5. %x  : Prints the value as hex dump of byte values with lowercase.
	 * 6. %X  : Prints the value as hex dump of byte values with uppercase.
	 * 7. % x : Prints the value as hex dump with spaces with lowercase.
	 * 8. % X : Prints the value as hex dump with spaces with uppercase. */
	fmt.Printf("%s\n", text)    // Output: Hello World
	fmt.Printf("%q\n", text)    // Output: "Hello World"
	fmt.Printf("%15s\n", text)  // Output:     Hello World
	fmt.Printf("%-15s\n", text) // Output: Hello World
	fmt.Printf("%x\n", text)    // Output: 48656c6c6f20576f726c64
	fmt.Printf("%X\n", text)    // Output: 48656C6C6F20576F726C64
	fmt.Printf("% x\n", text)   // Output: 48 65 6c 6c 6f 20 57 6f 72 6c 64
	fmt.Printf("% X\n", text)   // Output: 48 65 6C 6C 6F 20 57 6F 72 6C 64

	/*
	 * The following verb can be used with the boolean data type:
	 * 1. %t: Value of the boolean operator in true or false format (same as using %v). */
	fmt.Printf("%t\n", t) // Output: true

	/*
	 * The following verbs can be used with the float data type:
	 * 1. %e   : Scientific notation with 'e' as exponent.
	 * 2. %f   : Decimal point, no exponent.
	 * 3. %.2f : Default width, precision 2.
	 * 4. %6.2f: Width 6, precision 2.
	 * 5. %g   : Exponent as needed, only necessary digits. */
	fmt.Printf("%e\n", y)     // Output: 3.141000e+00
	fmt.Printf("%f\n", y)     // Output: 3.141000
	fmt.Printf("%.1f\n", y)   // Output: 3.1
	fmt.Printf("%10.2f\n", y) // Output:       3.14
	fmt.Printf("%g\n", y)     // Output: 3.141

	/*
	 * Think of fmt as Go's text formatting engine.
	 *
	 * Main printing functions you'll use:
	 * 1. fmt.Print* : Prints to stdout (terminal), Returns bytes written.
	 * 2. fmt.Fprint*: Prints any io.Writer, Returns bytes written.
	 * 3. fmt.Sprint*: Prints to string, Returns string.
	 *
	 * Rule of thumb:
	 * 1. Print  : rare, manual control.
	 * 2. Println: learning & debugging.
	 * 3. Printf : structured output.
	 *
	 * Golden rules to memorize:
	 * 1. Println: Prints daily debugging
	 * 2. Printf : Prints formatted output
	 * 3. Sprintf: Prints formatted string
	 * 4. %v     : Prints value
	 * 5. %T     : Prints type
	 * 6. %+v    : Prints struct fields
	 * 7. %#v    : Prints Go-syntax. */
	fmt.Print("XYZ\n")              // Output: XYQ\n
	fmt.Println("X", "Y")           // Output: X Y\n
	fmt.Printf("%s %d\n", "AB", 10) // Output: AB 10\n
}

func PrintSomethingWithSprintf() {
	const fullName, nickName = "John Doe", "John"

	/*
	 * In Golang, fmt package implements formatted I/O with functions.
	 * The fmt.Sprintf() function in Golang formats according to
	 * a format specifier and returns the resulting string.
	 *
	 * It returns a string instead of printing.
	 *
	 * You’ll use this for:
	 * 1. Logs
	 * 2. Errors
	 * 3. Building strings
	 * 4. APIs */
	message := fmt.Sprintf("Hello, %s! You can call me %s", fullName, nickName)
	fmt.Println(message) // Output: Hello, John Doe! You can call me John

	error := fmt.Sprintf("%s", "Something went wrong!")
	fmt.Println(error) // Output: Something went wrong!

	// You can also print with width, alignment, and precision.
	fmt.Printf("|%10s|\n", fullName)  // Output: |  John Doe|
	fmt.Printf("|%-10s|\n", fullName) // Output: |John Doe  |
}

func PrintSomethingWithLog() {
	logMessage := "User has been created"

	/*
	 * For real apps:
	 * 1. Adds timestamp
	 * 2. Writes to stderr
	 * 3. Better for production */
	log.Println(logMessage) // Output: 2026/02/10 16:30:31 User has been created
}
