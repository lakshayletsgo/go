// Interfaces
// We dont need to explicitly specify what interface to implement we just implement it methods and it automatically implements that interface
// Interfaces should be as small as possible
// If possible it should only have 2-3 methods to implement
package main

import (
	"fmt"
)

type vehicle interface {
	speed() int
	length() int
	details() string
}

// We can implement multiple interfaces
// We just have to implement the methods of all the interface we want to implement

type bike struct {
	width, breadth            int
	noOfWheels, speedPerWheel int
}

func (b bike) speed() int {
	return b.noOfWheels * b.speedPerWheel
}
func (b bike) length() int {
	return b.breadth * b.width
}
func (b bike) details() string {
	return fmt.Sprintf("Width of the bike is %d and its breadth is %d. It has %d number of wheels with each wheel of speed %d", b.width, b.breadth, b.noOfWheels, b.speedPerWheel)
}

type car struct {
	width, height, breadth      int
	noOfEngines, speedPerEngine int
}

func (c car) speed() int {
	return c.noOfEngines * c.speedPerEngine
}
func (c car) length() int {
	return c.height * c.breadth * c.width
}
func (c car) details() string {
	return fmt.Sprintf("Width of the car is %d, height is %d and breadth is %d. It has %d engines and each engine can produce speed upto %d", c.width, c.height, c.breadth, c.noOfEngines, c.speedPerEngine)
}

// We can also cast a interface to a class
type shape interface {
	permitre() float64
}
type circle struct {
	side int
}

func (c circle) permitre() float64 {
	return float64(4 * c.side)
}

func main() {
	fmt.Printf("Ok")
	sonet := car{
		width:          10,
		height:         20,
		breadth:        30,
		noOfEngines:    3,
		speedPerEngine: 200,
	}
	fmt.Println(sonet.details())
	fmt.Printf("Its speed is %d\n", sonet.speed())
	fmt.Printf("Its length is %d\n", sonet.length())

	hayabusa := bike{
		width:         20,
		breadth:       70,
		noOfWheels:    2,
		speedPerWheel: 200,
	}
	fmt.Println(hayabusa.details())
	fmt.Printf("Its top speed is %d\n", hayabusa.speed())
	fmt.Printf("Its is %d inches long\n", hayabusa.length())

	var s shape
	k, ok := s.(circle)
	if ok {
		fmt.Printf("It is a circle\n and its perimetre is %v\n", k.permitre())
	}

}
