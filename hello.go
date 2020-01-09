package main

import "fmt"

// Use var when outside of a function.
// If you don't set a value, zero value of the type is set.
// foo is therefore 0.

var foo int

func add(x int, y int) int {
	return x + y
}

type Dog struct {
	Animal
}
type Animal struct {
	Age int
}

func (a *Animal) Move() {
	fmt.Println("Animal moved")
}
func (a *Animal) SayAge() {
	fmt.Printf("Animal age: %d\n", a.Age)
}

func main() {
	x := 42
	for i := 0; i < 3; i++ {
		fmt.Println(add(42, x*i))
	}

	// Struct magicks

	d := Dog{}
	d.Age = 3
	d.Move()
	d.SayAge()

	// You can print the type of a variable thusly:

	fmt.Printf("%T\n", foo)
}
