package data_types

import "fmt"

const OUTPUT_FORMAT string = "%s (%d)\n"

func getStrings() {
	var txt1 string = "Hello World"
	txt2 := "Have a nice day!"

	// String can be empty, but they are not nil.
	var txt3 string

	fmt.Printf(OUTPUT_FORMAT, txt1, len(txt1)) // Output: Hello World (11)
	fmt.Printf(OUTPUT_FORMAT, txt2, len(txt2)) // Output: Have a nice day! (16)
	fmt.Printf(OUTPUT_FORMAT, txt3, len(txt3)) // Output:  (0)

	// You can combine strings with + operator.
	message1, message2 := "Oops!", "Something went wrong"

	fmt.Printf("%s\n", (message1 + " " + message2)) // Output: Oops! Something went wrong
}
