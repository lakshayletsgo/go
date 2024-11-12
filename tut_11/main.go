// Pointers
// They are used to point to a variable without creating its copy
// This is the syntax
// var p *int
package main

import (
	"fmt"
)

func main() {
	// var p *int

	// Or we can use this method
	myString := "Lakshay"
	myStringPr := &myString
	fmt.Printf("%v\n", *myStringPr) //This will give the value of the mystring i.e. dereference the pointer
	fmt.Printf("%v\n", &myStringPr) //This will give the address of the mystring

	// We can dereference the pointer and then re assign the same pointer
	var message *string
	messageDe := "Ok"
	message = &messageDe
	messageDe = *message
	messageDe = "hello"
	*message = messageDe
	fmt.Printf("%v\n", *message)

	fmt.Println("Pointers in GO")
}
