package main

import "fmt"

func main() {
	if true {
		fmt.Println("this ran")
	}

	if false {
		fmt.Println("this did not run")
	}

	if a := !false; a {
		fmt.Println("this ran")
	} else if a := false; a {
		fmt.Println("this did not run")
	} else {
		fmt.Println("this is default if above statement did not run")
	}

	if loop := true; loop {
		for i := 0; i < len(string(12345)); i++ {
			fmt.Println(i)
		}
	}
}
