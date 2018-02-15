package main

import "fmt"

func main() {
	generateMap()
	deletdMap()
}

func generateMap() {
	x := make(map[string]int)
	x["k"] = 5
	fmt.Println(x)
	fmt.Println(x["k"])

}

func deletdMap() {
	y := make(map[string]int)
	y["gg"] = 99
	fmt.Println("before delete: ", y["gg"])
	delete(y, "gg")
	fmt.Println("after delete: ", y["gg"])

}
