package identifier

import "fmt"

/*
 * Identifiers are the user-defined names of the program components.
 * In Golang, an identifier can be a variable name, function name, constant, statement label, package name, or type.
 *
 * There is a total of six identifiers available in the code:
 * 1. identifier: Name of the package.
 * 2. CallIdentifiers: Name of the function.
 * 3. name: Name of the variable.
 * 4. name2: Name of the variable.
 * 5. _nickName: Name of the variable.
 * 6. full_name: Name of the variable.
 *
 * There are certain valid rules for defining a valid Golang identifier.
 * These rules should be followed, otherwise, we will get a compile-time error.
 *
 * The name of the identifier:
 * 1. Must begin with a letter or an underscore(_).
 * 2. May contain the letters 'a-z' or 'A-Z' or digits 0-9 as well as the character '_'.
 * 3. Should not start with a digit.
 * 4. Is case-sensitive.
 * 5. Doesn't allow keyword to become an identifier name.
 * 6. has no limit on the length of the name of the identifier,
 * but it is advisable to use an optimum length of 4–15 letters only. */
func CallIdentifiers() {
	var name string = "John Doe"
	var name2 string = "Jane Doe"
	var _nickName string = "Johnny"
	var full_name string = "John Albert Doe"

	fmt.Println(name)      // Output: John Doe
	fmt.Println(name2)     // Output: Jane Doe
	fmt.Println(_nickName) // Output: Johnny
	fmt.Println(full_name) // Output: John Albert Doe
}
