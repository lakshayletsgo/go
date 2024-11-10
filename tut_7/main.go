// Loops
// For 	Loops in go are similar to all other loops
// Initialisation Condition Increment
// In this it is just that it will run n-1 times
// Here first incrementation happen then condition is checked

// In go condition is optional
// Its our choice whether to include it or not
// We can skip it and then the loop will go forever until it is explictly stopped

// In go we dont have a while keyword
// If we want to use while loop we can use it using the for only we just have to pass the condition

// In go we also have logical statments (||, &&)
// We also have modulo operator(%)

// We can also use break and continue clauses in loops in go
package main

import "fmt"

func main() {
	var_1 := 2
	fmt.Print("Loops")
	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {

			fmt.Print("*")
		}
		fmt.Println()
	}

	// for i := 0; ; i++ {
	// 	fmt.Println(i)
	// 	if i == 20 {
	// 		return
	// 	}

	// }

	// This is a while loop
	for var_1 < 5 {
		fmt.Print(var_1)
		var_1++
	}

	for i := 1; i < 100; i++ {
		if i%2 == 0 && i%5 == 0 {
			fmt.Println("Multiple of 10")
		} else if i%2 == 0 {
			fmt.Println("Multiple of 2")

		} else if i%5 == 0 {
			fmt.Println("Multiple of 5")
		} else {
			fmt.Println(i)
		}
	}

	// Prime Function
	primm := 13
	for i := 0; i < primm; i++ {
		if i == 2 {
			fmt.Println(i)
		}
		if i%2 != 0 {
			temp := 0
			for j := 2; j*j < i+1; j++ {
				if i%j == 0 {
					temp++
				}

			}
			if temp == 0 {
				fmt.Println(i)
			}
		}

	}
}
