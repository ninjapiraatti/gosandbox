package main

import (
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ninjapiraatti/gotools"

	"./kylpynallet"
)

func main() {
	x := 42
	for i := 0; i < 3; i++ {
		fmt.Println(gotools.Add(42, x*i))
	}

	db, err := sql.Open("mysql", "gotesti_1:Xv7NNK181fdAiBe2@tcp(dedi1450.your-server.de:3306)/gotesti")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO animals(name) VALUES('koiro')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

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

	n := kylpynallet.KylpyNallet1()
	fmt.Println(n)
	slicetest()
	fmt.Println(gotools.Multiply(2, 6))
}
