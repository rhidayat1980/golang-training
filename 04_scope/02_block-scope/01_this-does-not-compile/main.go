package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}

/*
this does not compile because of x is not accessible from foo function
to make it works we need to declare x in package level, not in block/function level
*/
