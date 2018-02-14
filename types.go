package main

import "fmt"

func main() {
	intContent()
	floatContent()
	complexContent()
	stringContent()
	boolContent()

}

func intContent() {
	//int
	//int8 -128~127
	//int16 -32768~32767
	//unit16 0~65535
	//int32 -2147483648~2147483647
	//unit32 0~4294967295
	//int64 -9223372036854775808~9223372036854775807
	//unit64 0~18446744073709551615

	fmt.Println("1 + 1 = ", 1 + 1 )

}

func floatContent() {
	//defauult float64

	var floatValue float64 = 7.0
	var floatValue2 = 3.0
	var floatValue3 = floatValue + floatValue2
	fmt.Println("floatValue + floatValue2 = ", floatValue3 )

	//float32  float64 can't calculate meantimes , need to convert
	var floatValue4 float64 = 1.3
	var floatValue5 float32 = 2.7
	fmt.Println("floatValue4 + floatValue5 =", float32(floatValue4) + floatValue5)

	//fmt.Println("floatValue4 + floatValue5 =", floatValue4 + floatValue5)
	//invalid operation: floatValue4 + floatValue5 (mismatched types float64 and float32)

}

func complexContent() {
	//default complex128

	var complexValue complex64 = 1.2 + 12i
	fmt.Println("complexValue =", complexValue)
	fmt.Println("complexValue real =", real(complexValue))
	fmt.Println("complexValue imag =", imag(complexValue))
	
}

func stringContent() {
	stringValue := "hello world!"
	fmt.Println(stringValue)
}

func  boolContent() {
	//default fault
	var boolValue bool
	fmt.Println("boolValue", boolValue)
}

