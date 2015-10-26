package main

import	"fmt"

func main() {
	var num1, num2 int

	fmt.Print("Enter in the first number ")
	fmt.Scan(&num1)
	fmt.Print("Enter in the second number ")
	fmt.Scan(&num2)

	num3 := num1%num2
	fmt.Println("The remaining is ", num3)
}
