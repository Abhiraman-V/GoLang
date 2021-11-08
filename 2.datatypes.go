package main

import "fmt"

/*
	Basic type:
		Numbers
		strings
		booleans.

	Aggregate type:
		Array
		structs.

	Reference type:
		Pointers,
		slices,
		maps,
		functions,
		channels.

	Interface type

	The Basic Data Types are further categorized into three subcategories which are:
	Numbers:
		Un/signed Integers,
		Float,
		Complex
	Booleans
	Strings
*/

func main() {

	//Numbers
	// 8-bit unsigned int
	var A uint8 = 25
	fmt.Println("\nA = ", A, "\nA+1 =", A+1)

	// 16-bit signed int
	var B int16 = 377
	fmt.Println("\nB = ", B, "\nB+2 =", B+2, "\nB-2 =", B-2)

	// float
	a := 19.99
	b := 70.89

	c := b - a

	fmt.Printf("\nResult is: %f", c)
	fmt.Printf("\nThe type of c is : %T\n", c) // Display the type

	// complex
	var m complex128 = complex(6, 2)
	var n complex64 = complex(9, 2)

	fmt.Println("\n", m)
	fmt.Println("\n", n)

	// Display the type

	fmt.Printf("\nThe type of m is %T and the type of n is %T\n", m, n)

	//Boolean

	o := 5
	p := 8

	fmt.Println("\no == p:", o == p)
	fmt.Println("o != p:", o != p)

	//Strings

	string1 := "Go"
	string2 := "was designed in Google"
	result := string1 + " " + string2

	fmt.Println("\n", result)
	fmt.Printf("This is the type of %T", result)

}
