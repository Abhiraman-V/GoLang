package main

import "fmt"

/* Control Statements

Control statements are used to control the flow of execution of the program based on certain conditions.

1.Decision-making statement
	If statement
	If-else statement
	Nested-if statement
	If..else..if ladder

2.Loops
3.Switch

*/

func main() {

	//If statement
	var v int = 497

	if v < 1000 {
		fmt.Println("v is less than 1000")
	}
	fmt.Printf("Value of v is : %d\n", v)

	//If-else statement

	var v1 int = 1780

	if v1 < 1000 {
		fmt.Println("\nv1 is less than 1000")
	} else {
		fmt.Println("\nv1 is greater than 1000")
	}

	fmt.Printf("Value of v1 is : %d\n", v1)

	//Nested-if statement

	var v2 int = 2500
	var v3 int = 7870

	if v2 == 2500 {

		if v3 == 7870 {
			fmt.Println("Value of v2 is 2500 and v3 is 7870\n")
		}
	}

	//If..else..if ladder

	var v4 int = 680

	if v4 == 500 {

		fmt.Println("Value of v4 is 500\n")

	} else if v4 == 600 {

		fmt.Println("Value of v4 is 600\n")

	} else if v4 == 700 {

		fmt.Println("Value of v4 is 700\n")

	} else {

		fmt.Println("None of the values is matching\n")
	}
}
