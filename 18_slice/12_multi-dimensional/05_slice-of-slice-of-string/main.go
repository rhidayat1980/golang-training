package main

import "fmt"

func main() {

	var records [][]string
	// student 1
	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.0"
	student1[3] = "74.00"

	// store the record
	records = append(records, student1)

	student2 := make([]string, 4)
	student2[0] = "Gomez"
	student2[1] = "Lisa"
	student2[2] = "92.00"
	student2[3] = "96.00"

	// store the record
	records = append(records, student2)

	// print
	fmt.Println(records)

	for i, v := range student1 {
		fmt.Println(i, v)
	}

	for i, v := range student2 {
		fmt.Println(i, v)
	}

	for i, v := range records {
		fmt.Println(i, v)
	}
}
