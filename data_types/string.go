package data_types

import (
	"fmt"
	"strings"
)

func getStrings() {
	/*
	 * First thing to note is that a UTF-8 encoded string
	 * can be made of any character in Unicode.
	 * Since UTF-8 supports up to 4 bytes long character,
	 * this means a character's length could be anywhere between 1 to 4 bytes.
	 * Second thing to note is interestingly in Go, strings are immutable slices of bytes.
	 * This means if the character "å" takes up 3 bytes when encoded in UTF-8 standard,
	 * it will occupy 3 indices in string array and index 2 will return only part of this character.
	 * This means if we iterate through a string using its indices,
	 * we'd get the byte value of that index.
	 * This is no good when we want to compare characters
	 * to validate a palindrome as we need the entire character.
	 *
	 * A string is a sequence of bytes.
	 * A string is UTF-8 encoded and immutable (you cannot modify a string directly).
	 *  */
	alphabets := "a b c d A B C D"
	fmt.Printf("%s\n", alphabets)

	/*
	 * This string value "Fajar" is [70 97 106 97 114] (bytes UTF-8).
	 * Here below, Go automatically knows it's string. */
	var name = "Fajar"
	fmt.Printf("%s\n", name) // Output: Fajar

	// String can be empty, but they are not nil.
	var txt1 string
	fmt.Printf("%s (%d)\n", txt1, len(txt1)) // Output:

	// Don't assume len() = character count; Wrong for UTF-8.
	var txt2 string = "Hello World"
	fmt.Printf("%s (%d)\n", txt2, len(txt2)) // Output:

	txt3 := "Have a nice day!"
	fmt.Printf("%s (%d)\n", txt3, len(txt3)) // Output:

	// You can combine strings with + operator.
	message1, message2 := "Oops!", "Something went wrong"
	fmt.Printf("%s\n", (message1 + " " + message2)) // Output: Oops! Something went wrong

	/*
	 * Go has 2 types of string literals:
	 * 1. Double quotes ""
	 * 2. Backticks (raw string) ``
	 *
	 * Output:
	 * Lorem ipsum
	 * dolor sit amet */
	stc := "Lorem ipsum\ndolor sit amet"
	fmt.Printf("%s\n", stc)

	/*
	 * Raw string preserves format and ignores escape characters.
	 * Syntax:
	 * query := `SELECT * FROM users`
	 *
	 * Best for:
	 * 1. Multiline text
	 * 2. JSON
	 * 3. SQL
	 *
	 * Output:
	 * Lorem ipsum
	 *     dolor sit amet */
	stc = `Lorem ipsum
	dolor sit amet`
	fmt.Printf("%s\n", stc)

	/*
	 * A string is immutable, modify string directly causes error.
	 * Syntax:
	 * city := "jakarta"
	 * name[0] = 'J' */
	city := "jakarta"
	bytes := []byte(city)
	bytes[0] = 'J'
	city = string(bytes)
	fmt.Printf("%s\n", city) // Output: Jakarta

	/*
	 * You can find string length using len() function.
	 * Be careful, this counts bytes not characters, because of UTF-8. */
	fmt.Printf("%d\n", len(city)) // Output: 7
	fmt.Printf("%d\n", len("é"))  // Output: 2

	// Accessing character using variable[index], because it's byte.
	fmt.Printf("%d\n", city[0]) // Output: 74
	fmt.Printf("%c\n", city[0]) // Output: J

	/*
	 * Output:
	 * 0 J
	 * 1 a
	 * 2 k
	 * 3 a
	 * 4 r
	 * 5 t
	 * 6 a */
	for i, c := range city {
		fmt.Println(i, string(c))
	}

	s1 := "Japan"
	s2 := "Japan"
	fmt.Printf("%t\n", (s1 == s2)) // Output: true

	city = "Tokyo"
	// String can use rune slice (important for Unicode).
	runes := []rune(city)
	fmt.Printf("%v\n", runes)    // Output: [84 111 107 121 111]
	fmt.Printf("%c\n", runes)    // Output:
	fmt.Printf("%v\n", runes[0]) // Output: 84
	fmt.Printf("%c\n", runes[0]) // Output: T
}

func manipulateStrings() {
	// Don't use + repeatedly for many concatenations, use Builder; Better performance.
	var builder strings.Builder
	builder.WriteString("One Two One Two")
	builder.WriteString(" ")
	builder.WriteString("Three Two One Three")
	txt := builder.String()
	fmt.Printf("%s\n", txt)

	containsSubstr := strings.Contains(txt, "One")
	fmt.Printf("%t\n", containsSubstr)

	toUpper := strings.ToUpper(txt)
	fmt.Printf("%s\n", toUpper)

	toLower := strings.ToLower(txt)
	fmt.Printf("%s\n", toLower)

	replace := strings.Replace(txt, "One", "1", 2)
	fmt.Printf("%s\n", replace)

	replaceAll := strings.ReplaceAll(txt, "One", "1")
	fmt.Printf("%s\n", replaceAll)

	split := strings.Split(txt, " ")
	fmt.Printf("%v\n", split)
	fmt.Printf("%v\n", split[0])

	join := strings.Join(split, "-")
	fmt.Printf("%v\n", join)

	trimSpace := strings.TrimSpace("   Hello World              ")
	fmt.Printf("%s\n", trimSpace)

	hasPrefix := strings.HasPrefix(txt, "On")
	fmt.Printf("%t\n", hasPrefix)

	hasSuffix := strings.HasSuffix(txt, "ree")
	fmt.Printf("%t\n", hasSuffix)
}
