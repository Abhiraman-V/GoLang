package main

import (
	"fmt"
)

/*
  	A slice is declared just like an array, but it doesn’t contain the size of the slice.
  	So it can grow or shrink according to the requirement.

	Syntax:

		[]T
		or
		[]T{}
		or
		[]T{value1, value2, value3, ...value n}

	Here, T is the type of the elements. For example: 	var my_slice[]int

	• A slice contains three components:

	Pointer:
		The pointer is used to points to the first element of the array that is accessible through the slice.
	Length:
		The length is the total number of elements present in the array.
	Capacity:
		The capacity represents the maximum size upto which it can expand.

	• Creating and initializing a Slice

		Using slice literal:
			The creation of slice literal is just like an array literal,
			but with one difference you are not allowed to specify the size of the slice in the square braces[].

			var my_slice_1 = []string{"Geeks", "for", "Geeks"}

		Using an Array:
			Slice is the reference of the array. so slice can be created from the given array.
			For creating a slice from the given array first the lower and upper bound has to be specified,
			which means slice can take elements from the array starting from the lower bound to the upper bound.
			It does not include the elements above from the upper bound.

			array_name[low:high]

		Using already Existing Slice:
			It is also be allowed to create a slice from the given slice.
			To create a slice from the given slice. The syntax is same as using an array to create slice.

			slice_name[low:high]

		Using make() function:
			slice can be created using the make() function.
			This function takes three parameters,
				i.e, type, length, and capacity.
			Here, capacity value is optional.
			It assigns an underlying array with a size that is equal to the given capacity and returns a slice which refers to the underlying array.
			Generally, make() function is used to create an empty slice. Here, empty slices are those slices that contain an empty array reference.

			func make([]T, len, cap) []T
*/

func main() {

	arr := [9]string{"The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}

	fmt.Println("Array:", arr)

	slice := arr[1:6]

	fmt.Println("Slice:", slice)

	fmt.Printf("Length of the slice: %d", len(slice))

	fmt.Printf("\nCapacity of the slice: %d\n", cap(slice))

	// Creating and initializing a Slice

	// using the var keyword
	var slice_1 = []string{"Something", "fishy", "about", "this", "one"}

	fmt.Println("\nSlice 1:", slice_1)

	//using shorthand declaration
	slice_2 := []int{12, 45, 67, 56, 43, 34, 45}
	fmt.Println("Slice 2:", slice_2)

	//using Array
	slice_3 := arr[:2]
	fmt.Println("Slice 3:", slice_3)

	// Using already Existing Slice
	slice_4 := slice_1[:3]
	fmt.Println("Slice 4:", slice_4)

	// Using make() function
	var slice_5 = make([]int, 7)
	fmt.Printf("Slice 5 = %v, \n\nlength = %d, \ncapacity = %d\n", slice_5, len(slice_5), cap(slice_5))
}
