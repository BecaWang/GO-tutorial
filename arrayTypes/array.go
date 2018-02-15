package main

import "fmt"

func main() {
	generateBasicArray()
	generateBasicArray2()

}

func generateBasicArray() {
	var x [5]int
	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5

	sum := 0
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}

	fmt.Println(sum)

	sum2 := 0
	for _, value := range x {
		sum2 += value
	}
	fmt.Println(sum2)

}

func generateBasicArray2() {
	y := [5]int{
		1,
		2,
		3,
		4,
		5,
	}

	sum := 0
	for _, value := range y {
		sum += value

	}
	fmt.Println(sum)

}
