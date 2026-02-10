/*
 * Every Golang file must belong to a package.
 * This main keyword is special because this program is meant to be run, not imported.
 * If you want to build an executable (e.g., go run, go build), the package must be main.
 * Think of a package as a folder of related code.
 * The main package means "this is the starting point of my app".
 *
 * The below explanation is a full execution flow of Golang:
 * 1. Go starts the program.
 * 2. Finds package main.
 * 3. Looks for func main().
 * 4. Runs code inside main().
 * 5. Prints "Hello World".
 * 6. Program ends.
 *
 * There are concerns you need to know:
 * 1. No semicolons needed (Go inserts them automatically).
 * 2. Imports must be used; unused imports cause errors.
 * 3. File must be named something.go. */
package main

import (
	"fmt"

	"github.com/fajarstrtn/golang-tutorial/comment"
	"github.com/fajarstrtn/golang-tutorial/format"
	"github.com/fajarstrtn/golang-tutorial/identifier"
	"github.com/fajarstrtn/golang-tutorial/introduction"
)

/*
 * When you run a program, Golang automatically starts executing main function.
 * No main function = nothing runs. */
func main() {
	introduction.Greet()
	comment.ReadSingleLineComment()
	comment.ReadMultiLineComment()
	identifier.GenerateIdentifiers()
	identifier.GenerateKeywords()
	identifier.GenerateVariablesUsingVar()
	identifier.GenerateVariablesUsingShortVarDec()
	identifier.GenerateConstants()
	format.PrintSomething()
	format.PrintSomethingWithNewLine()
	format.PrintSomethingWithFormattingVerbs()
	format.PrintSomethingWithSprintf()
	format.PrintSomethingWithLog()

	// Try to call an exported variable from exported_variable.go file.
	fmt.Printf("%s called from main function\n", identifier.ExportedVariable) // Output: This is an exported variable called from main function
}
