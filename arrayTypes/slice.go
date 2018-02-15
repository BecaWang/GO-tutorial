package main

import "fmt"

func main() {
	generateSliceArray()
	slice_append()
	slice_copy()
}

func generateSliceArray() {
	sliceArray := make([]float64, 5)
    fmt.Println(len(sliceArray))
}

func slice_append() {
	slice1 := []int{1,2,3,}
	slice2 := append(slice1, 4, 5,)
    fmt.Println(slice1, slice2)
}

func slice_copy(){
	slice1 := []int{1,2,3}
	slice2 := make([]int, 2)
	slice3 := make([]int, 4)
	copy(slice2, slice1)
	copy(slice3, slice1)
    fmt.Println(slice1, slice2, slice3)
}