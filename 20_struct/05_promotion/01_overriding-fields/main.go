package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person
	First          string
	LicencedToKill bool
}

func main() {

	p1 := doubleZero{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   39,
		},
		First:          "Double zero seven",
		LicencedToKill: true,
	}

	p2 := doubleZero{
		person: person{
			First: "Miss",
			Last:  "Moneypenny",
			Age:   38,
		},
		First:          "if looks could kill",
		LicencedToKill: false,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.First, p1.person.First)
	fmt.Println(p2.First, p2.person.First)
}
