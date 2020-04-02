package main

import (
	"fmt"
	"os"

	"github.com/ninjapiraatti/gotools"
)

func basicprints() {
	gotools.Color(gotools.Magenta, "Here we go\n")

	separator := "**********"
	fmt.Printf("\n\n%s gotools.Countletters %s\n\n", separator, separator)
	fmt.Println(gotools.Countletters("abcdffxxyyz", 'z'))

	fmt.Printf("\n\n%s gotools.Fprime %s\n\n", separator, separator)
	fmt.Println(gotools.Fprime(7))

	fmt.Printf("\n\n%s gotools.Add %s\n\n", separator, separator)
	x := 42
	for i := 0; i < 3; i++ {
		fmt.Println(gotools.Add(42, x*i))
	}

	fmt.Printf("\n\n%s Various basic stuff. %s\n\n", separator, separator)
	animals := [2]string{"Gopher", "Honeybadger"}
	cities := []string{}
	cities = append(cities, "San Diego", "Mountain View")
	fmt.Printf("%q", cities)
	fmt.Println(len(cities))

	cities2 := []string{"San Diego", "Mountain View"}
	otherCities := []string{"Santa Monica", "Venice"}
	cities2 = append(cities2, otherCities...)
	fmt.Printf("%q", cities2)

	fmt.Printf("\n\n%s Maps %s\n\n", separator, separator)
	foods := map[string]int{
		"Carrot": 50,
		"Radish": 45,
		"Apple":  65,
	}
	for key, value := range foods {
		fmt.Printf("\nFood: %s. Value: %d.", key, value)
	}

	fmt.Printf("\n\n%s gotools.Creature %s\n\n", separator, separator)
	seepra := gotools.Creature{}
	seepra.Species = ""
	seepra.Name = "Sepi Seepra"
	seepra.Age = 0
	seepra.IsGoodBoi = true
	fmt.Println(seepra.Name)
	fmt.Println(animals[1])
	slicetest()

	fmt.Printf("\n\n%s gotools.Multiply %s\n\n", separator, separator)
	fmt.Println(gotools.Multiply(2, 6))

	fmt.Printf("\n\n%s gotools.WordsToMap %s\n\n", separator, separator)
	fmt.Println(gotools.WordsToMap("foo bar buzz foo"))

	fmt.Printf("\n\n%s gotools.IsItFortytwo %s\n\n", separator, separator)
	fmt.Println(gotools.IsItFortytwo(126))

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
