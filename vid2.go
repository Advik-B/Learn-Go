package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello, World!")
	var name string = "Advik"
	var age int8 = 12
	fmt.Println("Hello " + name)
	fmt.Println("Your are is", age)
	// this can be done in as:
	name1 := string("Advik")
	age1 := int8(15)
	fmt.Println("Hello " + name1)
	fmt.Println("Your are is", age1)
	/*
	This is a multi line comment
	Isnt it cool?
	*/
	// this is a single line comment
	
	
}