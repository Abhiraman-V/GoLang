package main

import (
	"fmt"
)

/*
	The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.

	Constant Rules
		• Constant names follow the same naming rules as variables
		• Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
		• Constants can be declared both inside and outside of a function

	Constant Types

		• Typed constants
		• Untyped constants

	Typed Constants
		Typed constants are declared with a defined type:
			const A int = 1
	Untyped Constants
		Untyped constants are declared without a type:
			const A = 1

*/

const A int = 1 //Typed constants
const PI = 3.14 //Untyped constants

func main() {

	fmt.Println(A)
	fmt.Println(PI)
}
