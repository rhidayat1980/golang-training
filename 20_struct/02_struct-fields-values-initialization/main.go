package main

import (
	"fmt"
	"reflect"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

	fmt.Println(reflect.TypeOf(p1))
	fmt.Println(reflect.TypeOf(p2))

	fmt.Println(p1)
	fmt.Println(p2)

	p3 := person{first: "rachmat", last: "hidayat", age: 39}
	fmt.Println(p3)

	fmt.Println(p3.first, p3.last, p3.age)
}
