package main

import "fmt"

/*
  The simplest way to implement the queue data structure in Golang is to use a slice.
  Since a queue follows a FIFO (First-In-First-Out) structure, the dequeue and enqueue operations can be performed as follows:

    Use the built-in append function to enqueue.
    Slice off the first element to dequeue.

*/
func enqueue(queue []int, element int) []int {
	queue = append(queue, element)
	fmt.Println("Enqueued:", element)
	return queue
}

func dequeue(queue []int) []int {
	element := queue[0]
	fmt.Println("Dequeued:", element)
	return queue[1:]
}

func main() {
	var queue []int

	queue = enqueue(queue, 10)
	queue = enqueue(queue, 20)
	queue = enqueue(queue, 30)

	fmt.Println("Queue:", queue)

	queue = dequeue(queue)
	fmt.Println("Queue:", queue)

	queue = enqueue(queue, 40)
	fmt.Println("Queue:", queue)
}
