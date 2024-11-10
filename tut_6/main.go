// Errors
// They are just interfaces
// They return message and we just have to check whether they are error or not
// This error handling is better than javascript handling error as we dont have to use the try catch statment and we can have different error handling for different calls
// We can implement error by ourself but there is also a package that is dedicated to errors and if just want to get a string output we can use this package
// Else we have to write so many lines to get the output

package main

import (
	"errors"
	"fmt"
)

// We just have to write errors.New("What we want to print")

// type error interface{
// 	Error() string
// }
// This is the error interface

// We can use our own struct to implement the error interface
// We just have to implement the error method in that

type bag struct {
	color string
}

func (b bag) Error() string {
	return fmt.Sprintf("This bag color is not available %v\n", b.color)
}

func getHeight() (int, error) {
	return 20, nil

	// To throw an error we can write
	// return 20, fmt.Errorf("We can't get height for below 5,11 foot")
	// return 20, fmt.Errorf("OK We UNDERSTAND THAT")
}
func getError(color string) error {
	if color == "Magenta" {
		return bag{color: color}
	}
	return nil
}

func getError2(color string) error {
	if color == "purple" {
		return errors.New("OUT OF STOCK")
	}
	return nil
}

func main() {
	// For eg if we know we can have error in the call getHeight()
	height, err := getHeight()
	if err != nil {
		fmt.Println("There is an error")
		fmt.Println(err)
		return
	}
	fmt.Println(height)

	puma := bag{}
	puma.color = "Magenta"

	// This will throw an error
	fmt.Print(getError(puma.color))

	addidas := bag{
		color: "purple",
	}

	fmt.Print(getError2(addidas.color))

}
