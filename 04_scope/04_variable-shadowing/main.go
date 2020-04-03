package main

import "fmt"

func max(x int) int {
	return 42 + x
}
func main() {
	max := max(72)
	fmt.Println(max) // max is now the result, not the function
}

/*
don't do this, bad coding practice to shadow variables
*/
