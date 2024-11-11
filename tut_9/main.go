// Maps in go
// We have hashmap in go
// Maps are used to access any element in 0(1) time
// Map keys can't be anything
// They should be something that can be compared like string boolean integer
// That can't be a key is map, function and slices
// We can also have nested maps in go
// Or we can use the struct as a key
// In go we have rune instead of character
package main

import (
	"errors"
	"fmt"
)

type user struct {
	age   int
	name  string
	caste string
}

// This function return map for the color code
func colorToCodeMap(color []string, code []int) (map[string]int, error) {
	colorCode := make(map[string]int)
	if len(color) != len(code) {
		return nil, errors.New("NUMBERS OF COLOURS ARE NOT EQUAL TO NUMBERS OF CODES")
	}
	for i := 0; i < len(color); i++ {
		colors := color[i]
		codes := code[i]
		colorCode[colors] = codes //We can also use a struct if we want to
	}
	return colorCode, nil
}

// This function make a map for the occurance of the name with rune as the intial of the name
func namesToMap(names []string) map[rune]map[string]int {
	namesMap := make(map[rune]map[string]int)
	for _, name := range names {
		if name == "" {
			continue
		}
		firChar := name[0]
		_, ok := namesMap[rune(firChar)]
		if !ok {
			namesMap[rune(firChar)] = make(map[string]int)
		}
		namesMap[rune(firChar)][name]++

	}
	return namesMap

}

func adultCheck(userMap map[string]user, name string) (deleted bool, err error) {
	userExistOrNot, ok := userMap[name]
	if !ok {
		return false, errors.New("This user don't exist")
	}
	if userExistOrNot.age < 18 {
		return false, nil
	}
	return true, nil
}

func main() {
	fmt.Println("Maps in go")

	// This is the syntax for the map
	// In this we first write the key type then the value type
	ages := map[string]int{
		"John": 37,
		"Mary": 21,
	}
	fmt.Println(len(ages))

	// Or we can create a map
	// colors:=make(map[string]int)

	color := []string{"Black", "Orange", "Green"}
	code := []int{10, 20, 30}
	colorCodeMap, err := colorToCodeMap(color, code)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(colorCodeMap["Black"])
	}

	// We can get an element by their key
	// colorCodeMap["Green"]

	// We can also delete an element by using the buil in delete function
	delete(colorCodeMap, "Orange")
	fmt.Println(colorCodeMap["Orange"])

	// We can use this syntax to get an element
	// Here if the element is present then the bool ok will be true else false
	elem, ok := colorCodeMap["Orange"]
	if ok {

		fmt.Println(elem)
	} else {
		fmt.Println("No Element found")
	}

	user1 := user{
		name:  "Lakshay",
		age:   20,
		caste: "Baniya",
	}
	user2 := user{
		name:  "Rahul",
		age:   17,
		caste: "Punjabi",
	}
	user3 := user{
		name:  "Mohan",
		age:   21,
		caste: "Baniya",
	}
	newMap := map[string]user{
		"Lakshay": user1,
		"Rahul":   user2,
		"Mohan":   user3,
	}
	adultOrNot, err := adultCheck(newMap, "Rahul")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(adultOrNot)
	}

}
