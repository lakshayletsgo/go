// Fun Fact Nibbles is type of variable that is 1 byte = 2 nibbles
//  1 nibble = 4 bits
// This is the first code of go
// Go is a compiled programming language
// That means it create a byte code first and then run it
// main function is the entry point of any go code

package main

import "fmt"

// func main() {
// 	fmt.Println("Hello From Lakshay")
// }

// Variable
// bool string int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // alias for uint8
// rune // alias for int32
// Represent a unicode code point
// float32 float64
// complex64 complex128

func main() {
	// This is how you declare a variable
	firstVar := 2

	// if we want a particular assignement then we need to use
	var lakshay string

	lakshay = "My name is lakshay"
	fmt.Println(firstVar, lakshay, firstVar)

	// To get the type of any variable
	// Here we need to use the printf that means formating library we need to use the type function
	fmt.Printf("The type function is %T\n", lakshay)

	// We can declare multiple variables in a single line
	var_1, var_2 := 32, 43
	fmt.Println(var_1, var_2)

	//We can type convert in this by using this
	var_3 := float32(var_1)
	fmt.Println(var_3)

	// Constants are a thing in go
	// We can't change the constants
	const name = "Lakshay"
	const age = 21
	fmt.Println(name, age)

	// Formating operators
	// %v is the default format operator
	fmt.Printf("My name is %v", "lakshay")
	fmt.Printf("My age is %v\n", 20)

	// If we dont want the default one we can use the
	// %s for string
	// %d for decimal(int) form
	// %f for float
	// If we use any decimal like %.2f --> this means 2 decimal places include

	msg := fmt.Sprintf("My name is %s and my age is %d", "Lakshay", 20)
	//Sprintf means we want to just format the line and not print it out

	fmt.Printf(msg)
}
