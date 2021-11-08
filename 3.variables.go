package main

import (
	"fmt"
)

/*
	Rules for Naming Variables:

		• Variable names must begin with a letter or an underscore(_). And the names may contain the letters ‘a-z’ or ’A-Z’ or digits 0-9 as well as the character ‘_’.
			Geeks, geeks, _geeks23  // valid variable
			123Geeks, 23geeks      // invalid variable
		• A variable name should not start with a digit.
			234geeks  // illegal variable
		• The name of the variable is case sensitive.
			geeks and Geeks are two different variables
		• Keywords is not allowed to use as a variable name.
		• There is no limit on the length of the name of the variable, but it is advisable to use an optimum length of 4 – 15 letters only.

	Declaring a Variable
		Variables can be created in two ways:

		Using var keyword: var variable_name type = expression

		Using short variable declaration: variable_name:= expression
*/

func main() {

	var a1 = 487
	var b2 = "GoLang"
	var c3 = 454.80

	fmt.Printf("The value and type of a1 is : %d - %T\n", a1, a1)
	fmt.Printf("The value and type of b2 is : %s - %T\n", b2, b2)
	fmt.Printf("The value and type of c3 is : %f - %T\n", c3, c3)

	d4 := 985
	e5 := "Google"
	f6 := 685.66

	fmt.Printf("The value and type of d4 is : %d - %T\n", d4, d4)
	fmt.Printf("The value and type of e5 is : %s - %T\n", e5, e5)
	fmt.Printf("The value and type of f6 is : %f - %T\n", f6, f6)

}
