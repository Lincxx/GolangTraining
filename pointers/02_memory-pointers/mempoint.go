package main

//Write a program tha uses memory addresses and pointers

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	//the above code makes b a pointer to the memory address where an int is store
	//b is of type "int pointer"
	fmt.Println(b)
	fmt.Println(&b)

	*b = 42  //b says, "The value at this address, change it to 42"
	fmt.Println(a)

}
