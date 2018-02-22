package main

import "fmt"

func main() {
	fmt.Println(add(3, 4))
	fmt.Println(swap("hello","GO"))


}

func add(x, y int) int {
	return x + y
}

func swap(x,y string) (string, string){
	return y, x
}