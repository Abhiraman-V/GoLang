package main

import "fmt"

/*
	Go language has only a single loop that is For-loop,repetition control structure
	for initialization; condition; post{
		// statements....
	}
*/

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
