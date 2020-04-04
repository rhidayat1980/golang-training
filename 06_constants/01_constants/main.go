package main

import "fmt"

const p = "death & taxes"

func main() {
	const q = 42
	fmt.Println("p - ", p)
	fmt.Println("q - ", q)

	const message string = "Hello there"
	fmt.Println(message)
}

// a CONSTANT is a simple unchanging value
