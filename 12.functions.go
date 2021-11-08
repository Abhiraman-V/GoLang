package main

import "fmt"

/*
	Function is a collection of statements that perform some specific task and return the result to the caller

	func function_name(Parameter-list)(Return_type){
    // function body.....
	}

	Function Arguments
		The parameters passed to a function are called actual parameters, whereas the parameters received by a function are called formal parameters.

		Call by value:
			The values of actual parameters are copied to function’s formal parameters and the two types of parameters are stored in different memory locations.
			So any changes made inside functions are not reflected in actual parameters of the caller.

		Call by reference:
			Both the actual and formal parameters refer to the same locations,
			so any changes made inside the function are actually reflected in actual parameters of the caller.
*/
func area(length, breadth int) int {

	Ar := length * breadth
	return Ar
}

func main() {

	fmt.Printf("Area of rectangle is : %d\n", area(12, 10)) //Function Call

	var p int = 10
	var q int = 20
	fmt.Printf("\nP = %d and Q = %d", p, q)

	// call by values
	swapvalue(p, q)
	fmt.Printf("\nP = %d and Q = %d\n", p, q)

	var r int = 10
	var s int = 20
	fmt.Printf("\nR = %d and S = %d", r, s)

	// call by reference
	swapreference(&r, &s)
	fmt.Printf("\nR = %d and S = %d", r, s)
}

func swapvalue(a, b int) int {
	var o int
	o = a
	a = b
	b = o
	return o
}

func swapreference(a, b *int) int {
	var o int
	o = *a
	*a = *b
	*b = o
	return o
}
