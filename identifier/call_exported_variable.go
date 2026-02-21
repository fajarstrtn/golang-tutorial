package identifier

import "fmt"

// Try to call an exported variable from exported_variable.go file.
func CallExportedVariable() {
	fmt.Printf("%s called from main function\n", ExportedVariable) // Output: This is an exported variable called from main function
}
