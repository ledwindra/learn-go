package main

import "fmt"

func main() {
	x := [5]float64{
		11,
		13,
		7,
		19,
		23,
	}

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}
