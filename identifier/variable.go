package identifier

import "fmt"

/*
 * For constants, once variables declared and initialized, it cannot be changed.
 * It must be initialized and cannot use := for const.
 *
 * Syntax:
 * const X = 10
 * const Y int = 5 */
const (
	FULLNAME_TEMPLATE string = "Full Name: %s (%T)\n"
	NICKNAME_TEMPLATE string = "Nick Name: %s (%T)\n"
	AGE_TEMPLATE      string = "Age      : %d (%T)\n"
	ADDRESS_TEMPLATE  string = "Address  : %s (%T)\n"
	ISMALE_TEMPLATE   string = "Is Male  : %t (%T)\n"
)

/*
 * Variable is a placeholder of the information which can be changed at runtime.
 *
 * In Go, variables are created in two different ways:
 * 1. Using var keyword                    : Variables are created using var keyword of a particular type,
 * connected with name and provide its initial value
 * 2. Using := (short variable declaration): The local variables which are declared and initialized
 * in the functions are declared by using short variable declaration
 *
 * The big difference between var and := (short variable declaration):
 * 1. The var keyword can be used inside and outside functions (at package level).
 * It allows declaration without initialization and is more explicit.
 * Variable declaration and value assignment can be done separately.
 * 2. := can be used only inside functions and not be allowed at package level.
 * Variable declaration and value assignment cannot be done separately (must be done in the same line).
 * It must initialize immediately.
 *
 * When should you use the var keyword:
 * 1. At package level: var appVersion = "1.0.0"
 * 2. You need a default value first: var name string
 * 3. You want explicit type (clarity): var price float64 = 99.99
 * 4. Grouped declarations
 *
 * When should you use := (short variable declaration):
 * 1. Inside functions (default choice)
 * 2. Quick, local variables: sum := a + b
 * 3. Loop variables
 *
 * Go community rule of thumb: "Use := unless you need var". */
func GenerateVariablesUsingVar() {
	/*
	 * The var declaration of variables are used for those local variables
	 * which need an explicit type that differs from the initializer expression,
	 * or for those variables whose values are assigned later
	 * and the initialized value is unimportant.
	 *
	 * The type of the variable is determined by the type of the expression.
	 *
	 * In the below syntax, either type or = expression can be omitted, but not both.
	 * If the type is removed, then the type of the variable is determined by the value-initialize in the expression.
	 * If the "= expression" is omitted, then the variable value is determined by its type's default value.
	 * The default value is usually 0.
	 *
	 * The var keyword is very common at package level. */
	var fullName1 string = "Turner Johnston"
	var nickName1 string = "Turner"
	var age1 int8 = 17
	var address1 string = "Flat 23, 829 Victoria Road, Trading Estate, Nottingham, Northern Ireland, CF8 9QC, United Kingdom"
	var isMale1 bool = true

	fmt.Printf(FULLNAME_TEMPLATE, fullName1, fullName1) // Output: Full Name: Turner Johnston (string)
	fmt.Printf(NICKNAME_TEMPLATE, nickName1, nickName1) // Output: Nick Name: Turner (string)
	fmt.Printf(AGE_TEMPLATE, age1, age1)                // Output: Age      : 17 (int8)
	fmt.Printf(ADDRESS_TEMPLATE, address1, address1)    // Output: Address  : Flat 23, 829 Victoria Road, Trading Estate, Nottingham, Northern Ireland, CF8 9QC, United Kingdom (string)
	fmt.Printf(ISMALE_TEMPLATE, isMale1, isMale1)       // Output: Is Male  : true (bool)

	// Variables declared and initialized without the explicit type.
	var fullName2 = "Catherine Flores"
	var nickName2 = "Cathie"
	var age2 = 16
	var address2 = "Flat 4, 7730 Manor Road, Education Campus, London, Northern Ireland, PJ5 7QE, United Kingdom"
	var isMale2 = false

	fmt.Printf(FULLNAME_TEMPLATE, fullName2, fullName2) // Output: Full Name: Catherine Flores (string)
	fmt.Printf(NICKNAME_TEMPLATE, nickName2, nickName2) // Output: Nick Name: Cathie (string)
	fmt.Printf(AGE_TEMPLATE, age2, age2)                // Output: Age      : 16 (int)
	fmt.Printf(ADDRESS_TEMPLATE, address2, address2)    // Output: Address  : Flat 4, 7730 Manor Road, Education Campus, London, Northern Ireland, PJ5 7QE, United Kingdom (string)
	fmt.Printf(ISMALE_TEMPLATE, isMale2, isMale2)       // Output: Is Male  : false (bool)

	/*
	 * If the expression is removed, then the variable holds zero-value for the type
	 * like 0 for number, false for boolean, "" for string, and nil for interface and reference type.
	 *
	 * There is no such concept of an uninitialized variable in Go.
	 *
	 * Here is the list of default value in Go:
	 * 1. int (0)
	 * 2. float64 (0.0)
	 * 3. string (empty string)
	 * 4. bool (false)
	 * 5. pointer, slice, map, chan (nil)
	 *
	 * Variables declared and initialized without expression.
	 * No undefined like JavaScript.
	 * Go hates surprises. */
	var fullName3 string
	var nickName3 string
	var age3 int
	var address3 string
	var isMale3 bool

	fmt.Printf(FULLNAME_TEMPLATE, fullName3, fullName3) // Output: Full Name:  (string)
	fmt.Printf(NICKNAME_TEMPLATE, nickName3, nickName3) // Output: Nick Name:  (string)
	fmt.Printf(AGE_TEMPLATE, age3, age3)                // Output: Age      : 0 (int)
	fmt.Printf(ADDRESS_TEMPLATE, address3, address3)    // Output: Address  :  (string)
	fmt.Printf(ISMALE_TEMPLATE, isMale3, isMale3)       // Output: Is Male  : false (bool)

	/*
	 * If you use type, then you are allowed to declare multiple variables
	 * of the same type in the single declaration.
	 * It is only possible to declare one type of variable per line. */
	var fullName4, nickName4, address4 string = "Dennis James", "Dennis", "Flat 1, 1925 Windsor Road, Market Square, Glasgow, Merseyside, JL5 1BF, United Kingdom"

	// If the type keyword is not specified, you can declare different types of variables on the same line.
	var age4, isMale4 = 17, true

	fmt.Printf(FULLNAME_TEMPLATE, fullName4, fullName4) // Output: Full Name: Dennis James (string)
	fmt.Printf(NICKNAME_TEMPLATE, nickName4, nickName4) // Output: Nick Name: Dennis (string)
	fmt.Printf(AGE_TEMPLATE, age4, age4)                // Output: Age      : 17 (int8)
	fmt.Printf(ADDRESS_TEMPLATE, address4, address4)    // Output: Address  : Flat 1, 1925 Windsor Road, Market Square, Glasgow, Merseyside, JL5 1BF, United Kingdom (string)
	fmt.Printf(ISMALE_TEMPLATE, isMale4, isMale4)       // Output: Is Male  : true (bool)

	// The type of variables is determined by the initialized values.
	var fullName5, nickName5, age5, address5, isMale5 = "Annabella Marsh", "Anne", 16, "Flat 2, 1926 Windsor Road, Market Square, Glasgow, Merseyside, JL5 1BF, United Kingdom", false

	fmt.Printf(FULLNAME_TEMPLATE, fullName5, fullName5) // Output: Full Name: Annabella Marsh (string)
	fmt.Printf(NICKNAME_TEMPLATE, nickName5, nickName5) // Output: Nick Name: Anne (string)
	fmt.Printf(AGE_TEMPLATE, age5, age5)                // Output: Age      : 16 (int)
	fmt.Printf(ADDRESS_TEMPLATE, address5, address5)    // Output: Address  : Flat 2, 1926 Windsor Road, Market Square, Glasgow, Merseyside, JL5 1BF, United Kingdom (string)
	fmt.Printf(ISMALE_TEMPLATE, isMale5, isMale5)       // Output: Is Male  : false (bool)

	// You are allowed to initialize a set of variables by the calling function that returns multiple values.
	var message1, message2 string = greet(fullName5)
	fmt.Println(message1) // Output: Hello, Annabella Marsh!
	fmt.Println(message2) // Output: Have a nice day!

	var (
		fullName6 string = "Jessica Solis"
		nickName6 string = "Jessica"
		age6      int8   = 17
		address6  string = "Flat 50, 6529 Church Street, Housing Estate, Edinburgh, West Midlands, RD3 4MV, United Kingdom"
		isMale6   bool   = false
	)

	fmt.Printf(FULLNAME_TEMPLATE, fullName6, fullName6) // Output: Full Name: Jessica Solis (string)
	fmt.Printf(NICKNAME_TEMPLATE, nickName6, nickName6) // Output: Nick Name: Jessica (string)
	fmt.Printf(AGE_TEMPLATE, age6, age6)                // Output: Age      : 17 (int8)
	fmt.Printf(ADDRESS_TEMPLATE, address6, address6)    // Output: Address  : Flat 50, 6529 Church Street, Housing Estate, Edinburgh, West Midlands, RD3 4MV, United Kingdom (string)
	fmt.Printf(ISMALE_TEMPLATE, isMale6, isMale6)       // Output: Is Male  : false (bool)

	var (
		fullName7 = "Alex Wong"
		nickName7 = "Alex"
		age7      = 17
		address7  = "Flat 1, 6530 Church Street, Housing Estate, Edinburgh, West Midlands, RD3 4MV, United Kingdom"
		isMale7   = true
	)

	fmt.Printf(FULLNAME_TEMPLATE, fullName7, fullName7) // Output: Full Name: Alex Wong (string)
	fmt.Printf(NICKNAME_TEMPLATE, nickName7, nickName7) // Output: Nick Name: Alex (string)
	fmt.Printf(AGE_TEMPLATE, age7, age7)                // Output: Age      : 17 (int)
	fmt.Printf(ADDRESS_TEMPLATE, address7, address7)    // Output: Address  : Flat 1, 6530 Church Street, Housing Estate, Edinburgh, West Midlands, RD3 4MV, United Kingdom (string)
	fmt.Printf(ISMALE_TEMPLATE, isMale7, isMale7)       // Output: Is Male  : true (bool)

	/*
	 * If the value of a variable is known from the start,
	 * you can declare the variable and assign a value to it on one line.
	 * The variable types of age8, address8, and isMale8 are inferred from their values. */
	var (
		fullName8 string = "Ronin Wolf"                                                                                    // Type is string.
		nickName8 string = "Ronin"                                                                                         // Type is string.
		age8             = 16                                                                                              // Type is inferred.
		address8         = "Flat 3, 6530 Church Street, Housing Estate, Edinburgh, West Midlands, RD3 4MV, United Kingdom" // Type is inferred.
		isMale8          = true                                                                                            // Type is inferred.
	)

	fmt.Printf(FULLNAME_TEMPLATE, fullName8, fullName8) // Output: Full Name: Ronin Wolf (string)
	fmt.Printf(NICKNAME_TEMPLATE, nickName8, nickName8) // Output: Nick Name: Ronin (string)
	fmt.Printf(AGE_TEMPLATE, age8, age8)                // Output: Age      : 16 (int)
	fmt.Printf(ADDRESS_TEMPLATE, address8, address8)    // Output: Address  : Flat 3, 6530 Church Street, Housing Estate, Edinburgh, West Midlands, RD3 4MV, United Kingdom (string)
	fmt.Printf(ISMALE_TEMPLATE, isMale8, isMale8)       // Output: Is Male  : true (bool)

	/*
	 * It is possible to assign a value to a variable after it is declared.
	 * This is helpful for cases the value is not initially known. */
	var (
		fullName9 string
		nickName9 string
		age9      int
		address9  string
		isMale9   bool
	)

	fullName9 = "Ruby Friedman"
	nickName9 = "Ruby"
	age9 = 17
	address9 = "Flat 4, 6531 Church Street, Housing Estate, Edinburgh, West Midlands, RD3 4MV, United Kingdom"
	isMale9 = false

	fmt.Printf(FULLNAME_TEMPLATE, fullName9, fullName9) // Output: Full Name: Ruby Friedman (string)
	fmt.Printf(NICKNAME_TEMPLATE, nickName9, nickName9) // Output: Nick Name: Ruby (string)
	fmt.Printf(AGE_TEMPLATE, age9, age9)                // Output: Age      : 17 (int)
	fmt.Printf(ADDRESS_TEMPLATE, address9, address9)    // Output: Address  : Flat 4, 6531 Church Street, Housing Estate, Edinburgh, West Midlands, RD3 4MV, United Kingdom (string)
	fmt.Printf(ISMALE_TEMPLATE, isMale9, isMale9)       // Output: Is Male  : false (bool)

	// Multiple variable declarations can also be grouped together into a block for greater readability.
	var (
		fullName10 string = "Sarah Blair"
		nickName10        = "Sarah"
		age10             = 17
		address10  string
		isMale10   bool = false
	)

	address10 = "Flat 40, 7771 Victoria Street, Shopping Centre, Manchester, Wales, HG3 8ZT, United Kingdom"

	fmt.Printf(FULLNAME_TEMPLATE, fullName10, fullName10) // Output: Full Name: Sarah Blair (string)
	fmt.Printf(NICKNAME_TEMPLATE, nickName10, nickName10) // Output: Nick Name: Sarah (string)
	fmt.Printf(AGE_TEMPLATE, age10, age10)                // Output: Age      : 17 (int)
	fmt.Printf(ADDRESS_TEMPLATE, address10, address10)    // Output: Address  : Flat 40, 7771 Victoria Street, Shopping Centre, Manchester, Wales, HG3 8ZT, United Kingdom (string)
	fmt.Printf(ISMALE_TEMPLATE, isMale10, isMale10)       // Output: Is Male  : false (bool)
}

func GenerateVariablesUsingShortVarDec() {
	/*
	 * Most of the local variables are declared and initialized (at the same time)
	 * by using short variable declarations due to their brevity and flexibility.
	 *
	 * Don't confuse in between := and =, as := is a declaration and = is assignment.
	 * It is not possible to declare a variable using := without assigning a value to it.
	 *
	 * It declares and initializes at the same time.
	 * Type is inferred automatically.
	 * It only works inside functions.
	 * This is the Go-style you'll see everywhere. */
	fullName1 := "Makayla Daugherty"
	nickName1 := "Makkie"
	age1 := 18
	address1 := "Flat 24, 9260 Windsor Road, Leisure Complex, London, West Yorkshire, LU7 3BD, United Kingdom"
	isMale1 := false

	fmt.Printf(FULLNAME_TEMPLATE, fullName1, fullName1) // Output: Full Name: Makayla Daugherty (string)
	fmt.Printf(NICKNAME_TEMPLATE, nickName1, nickName1) // Output: Nick Name: Makkie (string)
	fmt.Printf(AGE_TEMPLATE, age1, age1)                // Output: Age      : 18 (int)
	fmt.Printf(ADDRESS_TEMPLATE, address1, address1)    // Output: Address  : Flat 24, 9260 Windsor Road, Leisure Complex, London, West Yorkshire, LU7 3BD, United Kingdom (string)
	fmt.Printf(ISMALE_TEMPLATE, isMale1, isMale1)       // Output: Is Male  : false (bool)

	/*
	 * Using short variable declaration you are allowed to declare multiple variables
	 * of different types in the single declaration.
	 * The type of these variables are determined by the expression. */
	fullName2, nickName2, address2 := "Jerry Edwards", "Jerry", "Flat 25, 9260 Windsor Road, Leisure Complex, London, West Yorkshire, LU7 3BD, United Kingdom"

	// Using short variable declaration you are allowed to declare multiple variables in the single declaration.
	age2, isMale2 := 17, true

	fmt.Printf(FULLNAME_TEMPLATE, fullName2, fullName2) // Output: Full Name: Jerry Edwards (string)
	fmt.Printf(NICKNAME_TEMPLATE, nickName2, nickName2) // Output: Nick Name: Jerry (string)
	fmt.Printf(AGE_TEMPLATE, age2, age2)                // Output: Age      : 17 (int)
	fmt.Printf(ADDRESS_TEMPLATE, address2, address2)    // Output: Address  : Flat 25, 9260 Windsor Road, Leisure Complex, London, West Yorkshire, LU7 3BD, United Kingdom (string)
	fmt.Printf(ISMALE_TEMPLATE, isMale2, isMale2)       // Output: Is Male  : true (bool)

	/*
	 * In a short variable declaration,
	 * you are allowed to initialize a set of variables
	 * by the calling function that returns multiple values. */
	message1, message2 := greet(fullName2)
	fmt.Println(message1) // Output: Hello, Jerry Edwards!
	fmt.Println(message2) // Output: Have a nice day!

	/*
	 * A short variable declaration acts like an assignment only when for
	 * those variables that are already declared in the same lexical block.
	 * The variables that are declared in the outer block are ignored.
	 *
	 * Here, short variable declaration acts as an assignment for alias1 variable
	 * because same variable present in the same block
	 * so the value of alias1 is changed from John Doe to Joey Greer. */
	alias1, alias2 := "John Doe", "Eliza Clayton"
	alias3, alias1 := "Jane Doe", "Joey Greer"

	/*
	 * As you know, := (short variable declaration) must declare at least one new variable.
	 * Syntax:
	 * a := 10
	 * a := 20	// Causes compile-time error (no new variables on left side of :=compiler (NoNewVar)).
	 *
	 * Syntax:
	 * a := 10
	 * a, b := 20, 30	// This is valid.
	 *
	 * Syntax:
	 * alias1, alias3 := "Boston Conrad", "Ronin Wolf"
	 * alias1 := "Ivy Velasquez"	// Causes compile-time error (no new variables on left side of :=compiler (NoNewVar)).
	 *
	 * Quick comparison between var and := (short variable declaration):
	 * 1. Type inference: var (yes), := (yes)
	 * 2. Declare without value: var (yes), := (no)
	 * 3. Package level: var (yes), := (no)
	 * 4. Outside functions: var (yes), := (no)
	 * 5. Inside functions: var (yes), := (yes)
	 * 6. Idiomatic for local: var (no), := (yes) */
	fmt.Println(alias1) // Output: Joey Greer
	fmt.Println(alias2) // Output: Eliza Clayton
	fmt.Println(alias3) // Output: Jane Doe
}

func greet(fullName string) (string, string) {
	message1, message2 := "Hello, "+fullName+"!", "Have a nice day!"
	return message1, message2
}
