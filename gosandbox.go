package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	cities = append(cities, "San Dieggo", "Mountain View")
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

	fmt.Printf("\n\n%s chatBots %s\n\n", separator, separator)
	bots()

	fmt.Printf("\n\n%s httpExercise %s\n\n", separator, separator)
	httpInterface()

	fmt.Printf("\n\n%s Interface assignment %s\n\n", separator, separator)
	geometry()

	fmt.Printf("\n\n%s Maps %s\n\n", separator, separator)
	maps()

	fmt.Printf("\n\n%s Channels and concurrency %s\n\n", separator, separator)
	channels()
	os.Exit(0)
	//gotools.AddCreatureDB(animal)
}

func main() {

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) > 0 {
		arguments(argsWithoutProg)
	} else {
		reader := bufio.NewReader(os.Stdin)
		gotools.Color(gotools.ToggleGreen, "Here we go\n")
		fmt.Println("---------------------")

		for {
			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')
			// convert CRLF to LF
			text = strings.Replace(text, "\n", "", -1)

			if strings.Compare("hi", text) == 0 {
				fmt.Println("hello, Yourself")
			}

			if strings.Compare("run", text) == 0 {
				basicprints()
			}
		}
	}

}
