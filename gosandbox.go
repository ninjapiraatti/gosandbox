package main

import (
	"fmt"
	"os"

	"github.com/ninjapiraatti/gotools"
)

func basicprints() {
	gotools.Color(gotools.Green, "Here we go\n")
	x := 42
	for i := 0; i < 3; i++ {
		fmt.Println(gotools.Add(42, x*i))
	}

	animals := [2]string{"Gopher", "Honeybadger"}

	cities := []string{}
	cities = append(cities, "San Diego", "Mountain View")
	fmt.Printf("%q", cities)
	fmt.Println(len(cities))

	cities2 := []string{"San Diego", "Mountain View"}
	otherCities := []string{"Santa Monica", "Venice"}
	cities2 = append(cities2, otherCities...)
	fmt.Printf("%q", cities2)

	foods := map[string]int{
		"Carrot": 50,
		"Radish": 45,
		"Apple":  65,
	}
	for key, value := range foods {
		fmt.Printf("\nFood: %s. Value: %d.", key, value)
	}

	seepra := gotools.Creature{}
	seepra.Species = ""
	seepra.Name = "Sepi Seepra"
	seepra.Age = 0
	seepra.IsGoodBoi = true

	fmt.Println(seepra.Name)
	fmt.Println(animals[1])
	slicetest()
	fmt.Println(gotools.Multiply(2, 6))

	//gotools.AddCreatureDB(animal)
}

func main() {
	if len(os.Args) > 1 {
		basicprints()
	}
	// Display number of arguments
	fmt.Println(len(os.Args))
}

/*
func (u *gotools.Creature) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		u.Species, u.Name)
}
*/
