package main

import "fmt"

func main() {
	testExceptionSample()
}

func testExceptionSample() {
	defer func() {
		e := recover()
		fmt.Print(e)
		fmt.Print("\n")

	}()
	panic("this is exception!")

}
