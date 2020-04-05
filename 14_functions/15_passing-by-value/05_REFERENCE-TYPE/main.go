package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m)
	changeMe(m)
	fmt.Println(m)
	fmt.Println(m["Todd"])

	n := map[string]int{"shinta": 39}
	changeMe(n)
	fmt.Println(n)

	o := map[string]int{"andre": 37}
	changeMe(o)
	fmt.Println(o)

	p := make(map[string]int)
	changeMe(p)
	fmt.Println(p)
}

func changeMe(z map[string]int) {
	z["Todd"] = 44
}
