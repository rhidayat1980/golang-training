package main

import "fmt"

func hello() {
	fmt.Println("hello ")
}

func world() {
	fmt.Println("world")
}
func main() {
	defer hello()
	world()
}

// defer hello() means hold hello() function untill wolrd() done.
