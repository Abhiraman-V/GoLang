package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/*
	File Operations
		Golang offers a vast inbuilt library that can be used to perform read and write operations on files.
		In order to read from files on the local system, the io/ioutil module is put to use. The io/ioutil module is also used to write content to the file.
		The fmt module implements formatted I/O with functions to read input from the stdin and print output to the stdout.
		The log module implements simple logging package. It defines a type, Logger, with methods for formatting the output.
		The os module provides the ability to access native operating-system features.
		The bufio module implements buffered I/O which helps to improve the CPU performance.
*/

func CreateFile(filename, text string) {

	fmt.Printf("Writing to a file in Go lang\n")

	file, err := os.Create(filename)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	defer file.Close()

	len, err := file.WriteString(text)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}

func ReadFile(filename string) {

	fmt.Printf("\n\nReading a file in Go lang\n")

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", filename)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)

}

func main() {

	fmt.Println("Enter filename: ")
	var filename string
	fmt.Scanln(&filename)

	fmt.Println("Enter text: ")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')

	CreateFile(filename, input)
	ReadFile(filename)
}
