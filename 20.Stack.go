package main

import "fmt"

/*
	Stack
		A stack is an ordered data structure that follows the Last-In-First-Out (LIFO) principle.
		Stacks are most easily implemented in Golang using slices:

		An element is pushed to the stack with the built-in append function.
		The element is popped from the stack by slicing off the top element.

*/
type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func main() {
	var stack Stack

	stack.Push("One")
	stack.Push("Two")
	stack.Push("Three!!")

	for len(stack) > 0 {
		x, y := stack.Pop()
		if y == true {
			fmt.Println(x)
		}
	}
}
