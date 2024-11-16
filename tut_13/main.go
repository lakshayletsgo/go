package main

import (
	"fmt"

	Sum "github.com/lakshayletsgo/go/sample2"
	// If we have to import the package then we have to first upload it to github and then get it
	// We have to initialize it also like we first need to write go mod init github.com/lakshayletsgo/go/sample in sample file command prompt and then push it
	// Then we have to
	// replace (
	// 	github.com/lakshayletsgo/go/sample v1.0.0 => ../sample
	// 	github.com/lakshayletsgo/go/sample2 v1.0.0 => ../sample2
	// )

	// require (
	// 	github.com/lakshayletsgo/go/sample v1.0.0
	// 	github.com/lakshayletsgo/go/sample2 v1.0.0
	// )

	// This is go.mod file of this folder and then we can use the package

	"github.com/lakshayletsgo/go/sample"
)

func main() {

	fmt.Println(Sum.Sum2(1, 2))
	fmt.Print(Sum3.Sum(1, 2))
	// fmt.Println(Sum3.Sum(1,2))

}
