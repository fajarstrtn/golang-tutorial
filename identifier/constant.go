package identifier

import "fmt"

// The value of constants must be assigned when you declare it.
const PI float32 = 3.14
const SHAPE string = "Circle"
const COLOR = "Red"

// Multiple constants can be grouped together into a block for readability.
const (
	BORDER_TYPE      string = "Thick"
	BORDER_COLOR     string = "Black"
	BORDER_THICKNESS        = 2.5
)

/*
 * If a variable should have a fixed value that cannot be changed,
 * you can use the const keyword.
 *
 * The const keyword declares the variable as constant,
 * which means that it's unchangeable and read-only.
 *
 * Here are constant rules that you must know:
 * 1. Constant names follow the same naming rules as variables.
 * 2. Constant names are usually written in uppercase letters
 * (for easy identification and differentiation from variables).
 * 3. Constants can be declared both inside and outside of a function.
 *
 * There are two types of constants:
 * 1. Typed constants (declared with a defined type)
 * 2. Untyped constants (declared without a type
 * and the type of the constant is inferred from the value) */
func GenerateConstants() {
	fmt.Println(PI)               // Output: 3.14
	fmt.Println(SHAPE)            // Output: Circle
	fmt.Println(COLOR)            // Output: Red
	fmt.Println(BORDER_TYPE)      // Output: Thick
	fmt.Println(BORDER_COLOR)     // Output: Black
	fmt.Println(BORDER_THICKNESS) // Output: 2.5
}
