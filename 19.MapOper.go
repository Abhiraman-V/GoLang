package main

import "fmt"

/*
   Maps is an unordered collection of key and its associated value.

   A map is a reference type and declared in general as:
      var map1 map[keytype]valuetype
   eg  var map1 map[int]string

   Map Insert and Update operation
      Update and insert operation are similar in go map.
      If the map does not contain the provided key the insert operation will takes place and if the key is present in the map then update operation takes place.

   Map Delete operation
      You can delete an element in Go Map using delete() function.
         delete(map, key)

   Map Retrieve Element
      elem = m[key]

*/
func main() {
	x := map[string]int{"Kate": 28, "John": 37, "Raj": 20}
	fmt.Print(x)
	fmt.Println("\n", x["Raj"])

	fmt.Println("\n==============================\n")

	//Go Map Insert and Update operation
	m := make(map[string]int)
	fmt.Println(m)
	m["Key1"] = 10
	m["Key2"] = 20
	m["Key3"] = 30
	fmt.Println(m)
	m["Key2"] = 555
	fmt.Println(m)

	fmt.Println("\n==============================\n")

	//Go Map Delete operation
	fmt.Println(m)
	delete(m, "Key3")
	fmt.Println(m)

	fmt.Println("\n==============================\n")

	//Go Map Retrieve Element
	fmt.Println(m)
	fmt.Println("The value:", m["Key2"])
}
