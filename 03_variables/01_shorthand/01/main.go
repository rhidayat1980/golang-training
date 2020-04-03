package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `do you like my hat?`
	g := 'M'

	fmt.Printf("%v is %T\n", a, a)
	fmt.Printf("%v is %T\n", b, b)
	fmt.Printf("%v is %T\n", c, c)
	fmt.Printf("%v is %T\n", d, d)
	fmt.Printf("%v is %T\n", e, e)
	fmt.Printf("%v is %T\n", f, f)
	fmt.Printf("%v is %T\n", g, g)
}

/*
basic data types in golang:
	int
	float
	boolean
	string
	rune or int32
*/
