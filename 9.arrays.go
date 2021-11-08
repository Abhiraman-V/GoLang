package main

import (
	"fmt"
)

/*
  Arrays is used to store a collection of data of the same type
    Array declaration
      Using var keyword:
          Var array_name[length]Type
          or
          var array_name[length]Typle{item1, item2, item3, ...itemN}
          or
          Var array_name[index] = element

      Using shorthand declaration:
          array_name:= [length]Type{item1, item2, item3,...itemN}

      Multi-Dimensional Array:
          Array_name[Length1][Length2]..[LengthN]Type
*/

func main() {

	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}
	arr3 := [...]float32{1.2, 2.3, 3.4}

	arr4 := [3][3]string{{"C#", "C", "Python"},
		{"Java", "Scala", "Perl"},
		{"C++", "Go", "HTML"}}

	fmt.Println("Element of arr1 with index 1 :")
	fmt.Println(arr1[1])
	fmt.Println("Length of arr2 :")
	fmt.Println(len(arr2))
	fmt.Println("Elements of arr3 Array :")
	fmt.Println(arr3)
	fmt.Println("Elements of Multi-Dimensional arr4 Array :")
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			fmt.Println(arr4[x][y])
		}
	}
}
