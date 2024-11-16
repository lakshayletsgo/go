// Concurrency
// Go supports concurrency that is it support multiple functions to run side by side
// A CPU is broke in n parts, depending on the type of CPU and each part can run its own code
// So we can have n different code run simultaneously
// We have a function in go in which we can make it so that we can run multiple functions same time
// To do that we use the go keyword
// go doSomething()
// This will run the doSomething function along the running function
// This can't have a return type because we can't depend on this function for its end to complete
// When the compiler is reading the code it will skip this line and send this to another core of CPU and execute the rest lines

// To take a return we can have a channel to store the data from a function that is running concurrently.
// Channel needs to have a input and ouput happening at the same time in the functions otherwise it will entire a deadlock state
// To make a channel we use this syntax
// channel_1 :=make(chan int)
// This will be a int channel

// There are also empty channels also called as tokens
// In Tokens we don't care what is passed through the channel we are just concerned is something is passed or not
// What the value is we don't care
// ->ch
// THis is the syntax in this we discard the value of the channel

// Up till now we have made channel with no storage system
// So we can make channel buffer that can store data uptil one point after that it need a reader to read it
// ch:=make(chan int,100)

// We can close a channel if we are done with it
// Channels are closed from the sending side only
// We can do that by
// close(ch)

// We can also check whether it is closed or not from reading end
// v,ok:= <-ch

// We have a select statement to listen to multiple channels at the same time

// We can make a function read only channel only by using this
// func sample(temp <- chan time)
// In this we can only read values from time and not write in it
package main

import (
	"fmt"
)

func tempChannelSwitch(names, age chan string) {
	select {
	case i, ok := <-names:
		fmt.Printf("Ok name is %v\n", ok)
		fmt.Println(i)
	case v, ok := <-age:
		fmt.Printf("Ok name is %v\n", ok)
		fmt.Println(v)
	default:
		fmt.Println("No channel is selected")
	}

}

func tempChannel(names []string) {
	isOldChan := make(chan bool)
	go func() {

		for _, e := range names {
			if e == "Lakshay" {
				isOldChan <- true
				continue
			}
			isOldChan <- false
		}
	}()
	isOld := <-isOldChan
	fmt.Printf("Your name is lakshay : %v\n", isOld)
	isOld = <-isOldChan
	fmt.Printf("Your name is lakshay : %v\n", isOld)
	isOld = <-isOldChan
	fmt.Printf("Your name is lakshay : %v\n", isOld)
}

func main() {
	fmt.Println("Concurrency Control in GO")

	names := []string{"Lakshay", "Rahul", "Lakshay"}
	tempChannel(names)
}
