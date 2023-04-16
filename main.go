package main

import "fmt"

func adddiction(x, y int) int {
	return x + y
}
func devide(x, y int) int {
	return x / y
}
func subtraction(x, y int) int {
	return x - y
}
func multiplication(x, y int) int {
	return x * y
}
func main() {
	constanta := true
	var x, y int
	var usertext string
	for constanta {
		fmt.Print("enter x:")
		fmt.Scan(&x)
		fmt.Print("enter y:")
		fmt.Scan(&y)
		fmt.Print("enter aperation(divide,addiction,substraction,multiplication,stop):")
		fmt.Scan(&usertext)
		if usertext == "stop" {
			break
		}
		if usertext == "divide" {
			fmt.Print(devide(x, y))
		}
		if usertext == "addiction" {
			fmt.Print(adddiction(x, y))
		}
		if usertext == "subtraction" {
			fmt.Print(subtraction(x, y))
		}
		if usertext == "multiplication" {
			fmt.Print(multiplication(x, y))
		}
	}

}
