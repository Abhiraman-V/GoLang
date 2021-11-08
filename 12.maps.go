package main

import "fmt"

/*
    Golang Maps is a collection of unordered pairs of key-value.
    It is widely used as it provides fast lookups and values that can retrieve, update or delete with the help of keys.

    Creating and initializing Maps

        map[Key_Type]Value_Type{} //(var mymap map[int]string)
		or
        // Map with key-value pair
        map[Key_Type]Value_Type{key1: value1, ..., keyN: valueN}

		or using make() function
			make(map[Key_Type]Value_Type, initial_Capacity)
			make(map[Key_Type]Value_Type)
*/

func main() {

	// Creating and initializing empty map
	// Using var keyword
	var map_1 map[int]int

	fmt.Println("Map-1: ", map_1)
	// Using shorthand declaration and
	// using map literals
	map_2 := map[int]string{
		1: "Alpha",
		2: "Bravo",
		3: "Charlie",
		4: "Delta",
		5: "Echo",
	}
	fmt.Println("Map-2: ", map_2)

	// Using make() function
	var map_3 = make(map[float64]string)
	fmt.Println("Map-3: ", map_3)

	map_3[1.1] = "Inch"
	map_3[1.2] = "Feet"
	fmt.Println("Map-3: ", map_3)
}
