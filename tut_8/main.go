// Arrays
// There is slices in go in which we dont have any fixed size
// It is dynamically allocated
// Slices are build on array
// We can get values from array by using this
// newArr := arr[1:3]
// This will send values of arr from 1st variable to 2nd variable to newArr

// If we want to have all the values we just need to use this
// newArr:=arr[:]

// If we want to declare slice without using old array
// mySlice:=make([]int,5,10)
// make([]type, length, capacity)

// We can just use this syntax
// mySlice:=make([]int,5)

// mySlice:=[]string{"I","Am","Lakshay"}
// There is a build in length function(len) and build in capacity function (cap)

// len and cap function is safe function i.e they woudl return 0 if the arr is nil
// They wouldn't panic i.e. they won't throw runtime error

// There is a append function that we can use to add to a slice
// We can just call the append(slice,oneThing)

// If we want to append a slice we can do that also
// append(slice, slice2...)

// We can have 2d slice matrix

// we should avoid alloting slice to new slice after appending it
// As if we have mentioned the capacity then it will not allote a new memory, It will use the old memory only
package main

import (
	"fmt"
)

func returnArr() [2]int {
	return [2]int{
		1, 2,
	}
}

func returnArr2() ([]int, error) {
	values := returnArr()
	return values[0:1], nil
}

// Here ...int is same as []int i.e. it is a type of slice
func sum(nums ...int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

// Print 2d slice
func printSlice(rows, column int) [][]int {
	matrix := make([][]int, 0)
	for i := 0; i < rows; i++ {
		row := make([]int, 0)
		for j := 0; j < column; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	return matrix

}

func main() {
	fmt.Print("Slices")
	var newArr [2]int
	newArr2 := [6]int{2, 4, 1, 4, 1, 5}
	fmt.Print(newArr)
	fmt.Print(newArr2)

	var_2 := returnArr()
	fmt.Println(var_2)

	var_3 := newArr2[1:2] // Including 1 and excluding 2 value
	fmt.Println(var_3)
	fmt.Printf("Type of var_3 is %T\n", var_3)

	var_4, _ := returnArr2()
	fmt.Print(var_4)

	var_5 := make([]string, 4)
	var_5[0] = "Laksahy"
	var_5[1] = "Is"
	var_5[2] = "Jod"
	fmt.Println(var_5)

	var_6 := []string{"Lakshay", "Is", "Jod"}
	fmt.Println(var_6)
	// Here var_5 is equal to var_6

	fmt.Printf("Sum of 10 and 20 is %v", sum(10, 20))

	// If we want to pass a slice in an function and it dont have slice input
	// Then we don't have to pass the argument we can just extend the var_7
	var_7 := []int{1, 2, 3, 4}
	fmt.Printf("Its sum is %v\n", sum(var_7...))

	var_8 := 5
	var_9 := append(var_7, var_8)
	fmt.Println(var_9)

	matrix := printSlice(3, 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}

	// THis is range function to print things in range with their indexes
	diffColors := []string{"Red", "Black", "Orange"}
	for i, color := range diffColors {
		fmt.Println(i, color)
	}
}
