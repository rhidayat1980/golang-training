package main

import (
	"fmt"
	"reflect"
)

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning",
		"Jenny": "Bonjour",
	}

	myGreeting["Harleen"] = "Howdy"

	fmt.Println(myGreeting)

	fmt.Println(reflect.TypeOf(myGreeting))

	for k, v := range myGreeting {

		fmt.Println(k, v)
	}

	fmt.Println(len(myGreeting))
}
