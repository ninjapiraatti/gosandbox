package main

import "fmt"

// Use var when outside of a function.
// If you don't set a value, zero value of the type is set.
// foo is therefore 0.

var foo int

// User is a user
type User struct {
	ID             int
	Name, Location string
}

// Greetings greet ppl
func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		u.Name, u.Location)
}

// Player is a player of games
type Player struct {
	User
	GameID int
}

func add(x int, y int) int {
	return x + y
}

func main() {
	x := 42
	for i := 0; i < 3; i++ {
		fmt.Println(add(42, x*i))
	}

	// Struct magicks

	p := Player{}
	p.ID = 42
	p.Name = "Matt"
	p.Location = "LA"
	fmt.Println(p.Greetings())

	// You can print the type of a variable thusly:

	fmt.Printf("%T\n", foo)
}
