// Generics in go
// Generics help to maintain the dry state(Dont Repeat Yourself)
// Generics are special type of datatype that we can have to define datatype which we dont care
// Like if we want to make a function that spilts a array in two and don't care about the datatype so we can have generics in that
// We can also have customized generics to implement we can make a interface that have a sum method and then every class that has a sum method can be used through the special type generic function

package main

import "fmt"

type sumFuncIn interface {
	add() int
}

// In this if any func any class that implement this interface
// Any child of that class can use this adder function no one else can use it
func adder[T sumFuncIn](values T) int {
	newValue := values.add() + 10
	return newValue
}

func concataa[T any](s []T) ([]T, []T) {
	length := len(s)
	return s[:length/2], s[length/2:]
}

// Here as you can see we otherwise have to make two functions to return the different array type but here we have only one function doing both

func main() {
	fmt.Println("Generics in GO`")
	names := []string{"Lakshay", "Soham", "Rahul", "Mohan"}
	age := []int{10, 20, 30, 40}
	name1, name2 := concataa(names)
	age1, age2 := concataa(age)

	fmt.Println(name1, name2)
	fmt.Println(age1, age2)

}
