package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m)
	changeMe(m)
	fmt.Println(m)

}

func changeMe(z []string) {
	fmt.Println("changeMe start")
	z[0] = "Todd"
	fmt.Println(z)

	fmt.Println("changeMe end")
}
