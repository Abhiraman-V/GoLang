package main

import (
	"fmt"
	"strings"
)

/*
Variadic Functions
    The function that called with the varying number of arguments is known as variadic function

    Syntax:
        function function_name(para1, para2...type)type{
        // code...
        }

Anonymous functions
    An anonymous function is a function which doesn’t contain any name.(inline function)

    Syntax:
        func(parameter_list)(return_type){
        // code..

        // Use return statement if return_type are given
        // if return_type is not given, then do not use return statement
        return
        }()

*/

func joinstr(element ...string) string {
	return strings.Join(element, ".")
}

func main() {

	fmt.Println(joinstr())

	fmt.Println(joinstr("Android", "IOS"))
	fmt.Println(joinstr("Marshmallow", "Nougat", "Oreo", "Pie"))
	fmt.Println(joinstr("S", "E", "L", "i", "n", "u", "x"))

	// Anonymous function
	func() {
		fmt.Println("An anonymous function is also known as function literal.")
	}()

}
