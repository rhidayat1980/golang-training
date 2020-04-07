package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("%T %v\n", myAge, myAge)

	var yourAge int
	yourAge = 29
	fmt.Printf("%T %v \n", yourAge, yourAge)

	// since myAge and yourAge are different type
	// this does not work:
	// fmt.Println(myAge + yourAge)

	fmt.Println(int(myAge) + yourAge)
}
