package main

import "fmt"

func main() {
	visit([]int{1, 2, 3, 4, 5}, func(n int) {
		fmt.Println(n)
	})

}

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

/*
callback: passing a func as an argument
*/
