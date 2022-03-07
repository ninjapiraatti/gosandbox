package main

import "fmt"

// Remember value vs reference types! Value types: int, bool, struct, float, etc. Anything
// that is written directly into the memory. Reference types: maps, slices, channels, pointers, functions.
// Anything where the copied value is not the value itself, but a type. Therefore you don't
// need pointers when you work with maps for example.

type Color struct {
	name string
}

func maps() {
	var colors map[string]string

	purple := Color{
		name: "purple",
	}
	pointerToPurple := &purple

	colors3 := map[string]string{
		"red":   "Yeah thats red",
		"white": "ffffff",
		"green": "00ff00",
	}

	colors3["black"] = "yes thats black"

	delete(colors3, "white")
	fmt.Println(colors)
	fmt.Println(purple.name)
	purple.changeColor("blue")
	fmt.Println(purple.name)
	fmt.Println(colors3)
	printMap(colors3)
	fmt.Printf("%v", *pointerToPurple)
}

func (purplePointer *Color) changeColor(newColor string) {
	(*purplePointer).name = newColor
}

func printMap(colors map[string]string) {
	fmt.Printf("%v", colors)
	for color, hex := range colors {
		fmt.Println("Hex for ", color, "is ", hex)
		fmt.Println("--------")
	}
}
