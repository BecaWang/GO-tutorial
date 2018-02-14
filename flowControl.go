package main

import "fmt"
import "math"


func  main() {

	forFunc()
	ifFunc()
	fmt.Println(
            pow(3, 2, 10),
            pow(3, 3, 20),
        )
	switchFunc(2)
	switchExtendFunc(4)
}

func forFunc() {
        sum := 0
        for i := 0; i < 5; i++ {
            sum += i
        }
        fmt.Println(sum)

        sum2 := 1
        for sum2 < 5 {
            sum2 += sum2
        }
        fmt.Println(sum2)

}

func ifFunc() {
	
	a := 0
	if a > 0 {
        fmt.Println("ok")		
	}else{
        fmt.Println("not ok")		

	}
}

func pow(x, n, lim float64 ) float64 {
	if v:= math.Pow(x, n); v < lim {
		return v
	}else{
		return lim
	}
}

func switchFunc(i int) {
		switch i {
		case 1:
	        fmt.Println("1")
		case 2:
	        fmt.Println("2")
		case 3:
	        fmt.Println("3")	
		}
	}

func  switchExtendFunc(i int) {
	switch  {
    case i > 2:
        fmt.Println(i, "> 2")
    case i < 2:
        fmt.Println(i, "< 2")
    case i == 2:
        fmt.Println(i, "= 2")
    }
}

