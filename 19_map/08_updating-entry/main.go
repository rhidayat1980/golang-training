package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning",
		"Jenny": "Bonjour",
	}

	myGreeting["Harlen"] = "Howdy"
	fmt.Println(myGreeting)

	myGreeting["Harlen"] = "Gidday"

	fmt.Println(myGreeting)
}
