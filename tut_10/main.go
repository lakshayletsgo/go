// Advanced Functions in Go
// Functions that take other function as parametre

// First Class Function are those functions that are simple that take input and give output
// eg simple(a string) int

// High class function are those function that uses simple function as a parametre

// We can also create function that return function

// Defer is a keyword that is used to execute some code just before the function ends
package main

import "fmt"

func divide(a, b int) int {
	return a / b
}
func sub(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func add(a, b int) int {
	return a + b
}

// This function takes function as input and returns values which is result after we give input to the function
func aggregateFunc(a, b, c int, operation func(e, f int) int) int {
	return operation(operation(a, b), c)
}

// This function take function as a input and return a function
// In this it takes arithmatic function as a input and make it a self arithmatic function
func selfMath(arithmatic func(int, int) int) func(int) int {
	return func(i int) int {
		return arithmatic(i, i)
	}
}

// func deferFunc(temp string)bool {
// 	defer delete(temp)
// 	if(temp=="Hello"){
// 		return true;
// 	}
// 	if temp=="Ok" {
// 		return false;
// 	}
// }

func concatter() func(string) string {
	doc := ""
	return func(s string) string {
		doc += s
		return doc
	}
}

func main() {
	fmt.Println("Advanced Functions")
	fmt.Println(aggregateFunc(600, 20, 10, divide))
	fmt.Println(aggregateFunc(30, 20, 10, sub))

	selfDivide := selfMath(divide)
	selfSub := selfMath(sub)
	selfAdd := selfMath(add)
	selfMul := selfMath(mul)

	fmt.Println(selfDivide(20))
	fmt.Println(selfSub(10))
	fmt.Println(selfAdd(20))
	fmt.Println(selfMul(10))

	concatinator := concatter()
	concatinator("My ")
	concatinator("name ")
	concatinator("is ")
	concatinator("Lakshay.")

	fmt.Println(concatinator(" "))
}
