package main

import "fmt"

func main() {
	counter := 0

	// for i := 0; i < 10; i++ {
	// 	if i%3 == 0 {
	// 		counter += i
	// 	} else if i%5 == 0 {
	// 		counter += i
	// 	}
	// }
	// fmt.Println(counter)

	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			counter = counter + i
		} else if i%5 == 0 {
			counter = counter + i
		}
	}
	fmt.Println(counter)
}

/*
list all value that divisible by 3 and 5 and sum them all
*/
