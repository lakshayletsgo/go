// When we want to declare a variable we use := and when we have to alter the value we use =
// Functions
package main

import (
	"errors"
	"fmt"
)

func sum(a int, b int) int {
	return a + b
}

// We can also give parameters like this
// func sum(a,b int) int{}

func incre(c, d int) int {
	c = c + d
	return c
}

// In Go we can have multiple return values
func multipleReturn(e, f int) (int, int) {
	multi := e * f
	div := e / f
	return multi, div
}

// We can declare the variable in the return type only
func multiReturn2(g, h int) (i, j int) {
	i = g % h
	j = g ^ h
	return
}

// If we dont pass any value in the end then still it will pass the initiated values of the return type
func noReturn(k, l int) (m, n int) {
	return
}

// We should pass the return like this i.e Explicitly specify the return values
func noReturn2(o, p int) (q, r int) {
	return q, r
}

// We can also early return in this just like other languages
func earlyRetur(s, t int) (int, error) {
	if s == 0 || t == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return s / t, nil
}

func main() {
	fmt.Printf("Sum of 10 and 20 is %d\n", sum(10, 20))

	taxAndExpense := 20
	const actualPrice = 100
	taxAndExpense = incre(taxAndExpense, actualPrice)
	fmt.Printf("%d\n", taxAndExpense)

	// If we are not using any variable and have just declared and assigned it a value we cant have it
	// We need to remove it
	// For example if i remove the division variable from the output then i need to use
	// multiplication, _ := multipleReturn(var_1,var_2)
	// In this we are ignoring the value of second return
	var_1 := 20
	var_2 := 10
	multiplication, division := multipleReturn(var_1, var_2)
	fmt.Printf("%d is the multiplication and %d is the division \n", multiplication, division)

	// This is incorrect as earlyReturn is returning two variables
	// fmt.Printf("%d is the value when 1 is divided by 10 and if the value is 0 then this is the output %v", earlyRetur(0, 1))

}
