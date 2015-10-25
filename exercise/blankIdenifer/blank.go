package main

import (
	"fmt"
)

var a string = "this is stored in the variable a"
var b, _ = "Stored in b", "stored and then thrown away"

func main() {
	fmt.Println(a)
	fmt.Println(b)
}
