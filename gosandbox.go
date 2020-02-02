package main

import (
	"fmt"

	"./kylpynallet"
)

// Use var when outside of a function.
// If you don't set a value, zero value of the type is set.
// answer is therefore 0.

var answer int

func add(x int, y int) int {
	return x + y
}

func main() {
	x := 42
	for i := 0; i < 3; i++ {
		fmt.Println(add(42, x*i))
	}

	// Plain old array
	//animals := [2]string{"Gopher", "Honeybadger"}

	cities := []string{}
	cities = append(cities, "San Diego", "Mountain View")
	fmt.Printf("%q", cities)
	fmt.Println(len(cities))

	cities2 := []string{"San Diego", "Mountain View"}
	otherCities := []string{"Santa Monica", "Venice"}
	cities2 = append(cities2, otherCities...)
	fmt.Printf("%q", cities2)

	// Struct magicks

	/*
		p := Player{}
		p.ID = 42
		p.Name = animals[0]
		p.Location = "LA"
		fmt.Println(p.Greetings())
	*/

	// You can print the type of a variable thusly:

	fmt.Printf("%T\n", answer)
	n := kylpynallet.KylpyNallet1()
	fmt.Println(n)
	slicetest()
}
