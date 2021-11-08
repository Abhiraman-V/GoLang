package main

import (
	"fmt"
	"time"
)

/*
	A switch statement is a multiway branch statement.
	2 types in Go
		Expression Switch
		Type Switch

		Expression Switch
			switch optstatement; optexpression{
			case expression1: Statement..
			case expression2: Statement..
			...
			default: Statement..}

		Type Switch
			switch optstatement; typeswitchexpression{
			case typelist 1: Statement..
			case typelist 2: Statement..
			...
			default: Statement..}

*/

func main() {
	now := time.Now()

	fmt.Printf("Today is ")
	// Expression Switch
	switch day := now.Day(); day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid")
	}

	// Type Switch

	var value interface{}
	switch q := value.(type) {
	case bool:
		fmt.Println("value is of boolean type")
	case float64:
		fmt.Println("value is of float64 type")
	case int:
		fmt.Println("value is of int type")
	default:
		fmt.Printf("value is of type: %T", q)
	}
}
