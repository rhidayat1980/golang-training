package main

import "fmt"

func main() {
	a := 43

	fmt.Println("a =", a)
	fmt.Println("&a =", &a)

	var b = &a
	fmt.Println("b =", b)
	fmt.Println("*b =", *b)

	*b = 42

	fmt.Println("a =", a)
	fmt.Println("*b =", *b)
	fmt.Println(&a)
	fmt.Println(&b)
}
