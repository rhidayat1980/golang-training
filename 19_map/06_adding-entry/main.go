package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":  "Good morning",
		"Jane": "Bonjour",
	}

	fmt.Println(myGreeting)

	myGreeting["Harlen"] = "Howdy"

	fmt.Println(myGreeting)

	fmt.Println(len(myGreeting))
}
