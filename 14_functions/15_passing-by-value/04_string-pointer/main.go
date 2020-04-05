package main

import "fmt"

func main() {

	name := "Todd"
	fmt.Println(&name)

	changeMe(&name)

	fmt.Println(&name)

	fmt.Println(name)
}

func changeMe(z *string) {
	fmt.Println("start function")

	fmt.Println(z)
	fmt.Println(*z)

	*z = "Rocky"
	fmt.Println(z)
	fmt.Println(*z)
}

/*
jika tujuan akhir untuk merubah sebuah nilai, maka gunakan pointer
contoh diatas terlihat function changeMe nilai awal nya "Todd", setelah dijalan kan akan menghasilkan "Rocky"
*/
