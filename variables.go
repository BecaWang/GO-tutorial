package main

import "fmt"

var x ,y , z int = 3, 4, 5
var c bool;
var d string = "hello~"

func main() {
	fmt.Println(x ,y ,z, c, d)

    short()

}

func short() {
	e := 6
	f := "world"
	fmt.Println(e, f)
}