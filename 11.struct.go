package main

import "fmt"

/*
   A structure or struct is a user-defined type that allows to group/combine items of possibly different types into a single type.
   It can be termed as a lightweight class that does not support inheritance but supports composition.

   Declaring a structure:

       type Address struct {
           name string
           street string
           city string
           state string
           Pincode int
       }
   or
       type Address struct {
       name, street, city, state string
       Pincode int
       }

   To Define a structure:
       Syntax:
           var a Address

           We can also initialize a variable of a struct type using a struct literal as shown below:
           var a = Address{"Akshay", "PremNagar", "Dehradun", "Uttarakhand", 252636}

           Go also supports the name:value syntax for initializing a struct (the order of fields is irrelevant when using this syntax).
           And this allows you to initialize only a subset of fields. All the uninitialized fields are set to their corresponding zero value.
           Example:

           var a = Address{Name:”Akshay”, street:”PremNagar”, state:”Uttarakhand”, Pincode:252636}     //city:””


*/

type Car struct {
	Name, Model, Color string
	WeightInKg         float64
}

func main() {

	// Declaring a variable of a `struct` type. All the struct fields are initialized with their zero value.
	var a Car
	fmt.Println("Car A: ", a)

	// Declaring and initializing a struct using a struct literal
	b := Car{"Aston Martin", "DBX", "Matt-Black", 2357}
	fmt.Println("Car B: ", b)

	// Naming fields while initializing a struct
	c := Car{Name: "Ferrari", Model: "GTC4", Color: "Red", WeightInKg: 1920}
	fmt.Println("Car C: ", c)

	// Uninitialized fields are set to their corresponding zero-value
	d := Car{Name: "BMW"}
	fmt.Println("Car D: ", d)

	// Accessing struct fields using the dot operator
	fmt.Println("Car C Name: ", c.Name)
	fmt.Println("Car C Color: ", c.Color)

	// Assigning a new value to a struct field
	c.Color = "Black"

	// Displaying the result
	fmt.Println("Car C: ", c)
}
