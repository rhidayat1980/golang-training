package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	fmt.Println(average(data))
}

func average(sh []float64) float64 {
	total := 0.0

	for _, v := range sh {
		total = total + v
	}

	return total / float64(len(sh))
}
