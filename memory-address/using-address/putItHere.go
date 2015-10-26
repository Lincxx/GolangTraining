package main

import "fmt"

const metesToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)//we place the user's answer into the memory location
	yards := meters * metesToYards
	fmt.Println(meters, " meters is ", yards, " yards")
}
