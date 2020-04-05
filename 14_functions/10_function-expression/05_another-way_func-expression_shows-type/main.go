package main

import "fmt"

func main() {
	greet := makeGreeter()

	fmt.Println(greet())
	fmt.Printf("%T\n", greet)

	fmt.Println("---------------------")

	greet2 := makeGreeter2("hello world")
	fmt.Println(greet2())
	fmt.Printf("%T\n", greet2)
}

func makeGreeter() func() string {
	return func() string {
		return "Hello world"
	}
}

func makeGreeter2(s string) func() string {
	return func() string {
		return s
	}
}
