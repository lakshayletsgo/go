// Go is not a object oriented language
// Structure
// It is similar to classes in the java and other languages
package main

import (
	"fmt"
)

// We can use the variable using the dot notation like person.age
type person struct {
	age   int
	color string
	name  string
}

// WE can also make struct inside another struct
type beings struct {
	noOfLegs int
	isPerson person
}

// Anonymous struct
// In this we don't have to explicitly
// We use it because we want to use it once and not again
type car struct {
	color   string
	company string
	Wheels  struct {
		size  int
		brand string
	}
}

// If we want to make so that we don't need the dot notation we can just write
type chair struct {
	noOfLegs int
	color    string
	capacity
}
type capacity struct {
	holdingCapacity int
	width           int
}

// Now if we want to access the variable of capacity inside chair we dont need to write char.capacity.width
// WE can just write chair.width

type sq struct {
	side int
}

func (s1 sq) perimetre() int {
	return 4 * s1.side
}

func main() {
	// To initialise the struct we need to use this syntax
	lakshay := person{}
	lakshay.age = 20
	lakshay.color = "white"
	fmt.Printf("Lakshay age is %d and his colour is %v\n", lakshay.age, lakshay.color)

	livingBeing := beings{}
	livingBeing.noOfLegs = 2
	livingBeing.isPerson.age = 20
	livingBeing.isPerson.color = "brown"
	livingBeing.isPerson.name = "Lakshay"
	fmt.Printf("%v is the name of the being and he has %d legs and is of color %v\n", livingBeing.isPerson.name, livingBeing.noOfLegs, livingBeing.isPerson.color)

	wagonr := car{
		color:   "black",
		company: "Maruti Suzuki",
		Wheels: struct {
			size  int
			brand string
		}{
			size:  21,
			brand: "MRF"}}
	fmt.Printf("Wagonr is of color %v and of brand %v and it has wheels of size %d and of brand %v\n", wagonr.color, wagonr.company, wagonr.Wheels.size, wagonr.Wheels.brand)

	chair_1 := chair{
		noOfLegs: 4,
		color:    "brown",
		capacity: capacity{holdingCapacity: 5,
			width: 10,
		},
	}
	fmt.Printf("There is a chair in my room, it is of color %v and it has %d legs\n It can hold %d and is of width %d\n", chair_1.color, chair_1.noOfLegs, chair_1.holdingCapacity, chair_1.width)

	s1 := sq{
		side: 10,
	}
	fmt.Println(s1.perimetre())
}
