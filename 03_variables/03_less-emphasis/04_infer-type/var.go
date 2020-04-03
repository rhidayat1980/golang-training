package main

import "fmt"

func main() {
	var message = "Hello world"
	var a, b, c = 1, 2, 3

	fmt.Println(message, a, b, c)

	// checking type
	fmt.Printf("%v \t is %T\n", message, message)
	fmt.Printf("%v \t is %T\n", a, a)
	fmt.Printf("%v \t is %T\n", b, b)
	fmt.Printf("%v \t is %T\n", c, c)
}

/*
di golang kita boleh tidak menuliskan type sebuah data
dalam contoh di atas message, a, b, c di deklarasikan tanpa type mereka, namun golang
akan ototmatis memberikan data type sebelum kompilasi atau saat run time? cek hayo...
*/
