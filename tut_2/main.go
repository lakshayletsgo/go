// Conditional Statements (If-else)
package main

import "fmt"

func main() {
	var_1 := 10
	var_2 := 20
	if var_1 > var_2 {
		fmt.Println("var_1 is greater than var_2")
	} else if var_2 > var_1 {
		fmt.Printf("var_2 is greater than var_1 by %d\n", (var_2 - var_1))
	} else {
		fmt.Print("They are equal")
	}

	// We can initialize and add condition in the if block only
	if var_3 := 30; var_3 > 20 {
		fmt.Printf("Value of var_3 is %d", var_3)
	}
	// Here we cant use the var_3 because it is only defined for the if block
	// fmt.Printf("Value of var_3 is %d", var_3)

}
