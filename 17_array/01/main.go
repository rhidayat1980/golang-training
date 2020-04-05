package main

import "fmt"

func main() {
	var x [58]int
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[47] = 777
	fmt.Println(x)
	fmt.Println(x[47])

}
