package main

import "fmt"

type contact struct {
	greeting string
	name     string
}

// SwitchOnType function, exported
func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("Unknown")
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("McLeod")

	var t = contact{"Good to see you", "Tim"}
	SwitchOnType(t)

	SwitchOnType(t.greeting)
	SwitchOnType(t.name)

	fmt.Println("good to see you", t.name)
}
