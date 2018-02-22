package main

import "fmt"
import "time"


func loop(n int) {
	for i:=0; i<10; i++{
		fmt.Println(i)
	}
	
}

func main() {
go loop(0)
time.Sleep(time.Second * 1)
}