package main

import (
	"fmt"
	//"reflect"
)


func main() {
	 a := 43

	fmt.Println(a)
	fmt.Println(&a)

	//	var b *int = &a
	//
	//	fmt.Println(b)
	//the above code makes b a pointer to the memory address where an int is store
	//b is of type "int pointer"

	//this is the same
//	b := &a
//
//	fmt.Println(b)

//	//this is the same as well
//	b := &a
//
//	fmt.Println(reflect.TypeOf(b))

	//And there is this way too
	b := &a

	fmt.Printf("%T\n",b)

}
