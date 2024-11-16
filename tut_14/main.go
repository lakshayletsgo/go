// Installing package from outside
package main

import (
	"fmt"
	"time"

	// We will run the command go get (package location -->github.com/wagslane/go-tinytime)
	// go get github.com/wagslane/go-tinytime

	// This will download the packages and then add them to the go mod file
	tinytime "github.com/wagslane/go-tinytime"
)

func main() {
	fmt.Println("Hello")
	tt := tinytime.New(13412)
	tt = tt.Add(time.Hour * 48)
	fmt.Print(tt)

}
