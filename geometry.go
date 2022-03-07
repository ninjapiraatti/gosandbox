package main

import "fmt"

type Triangle struct {
	height float64
	base   float64
}

type Square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func geometry() {
	sq := Square{sideLength: 4}
	tr := Triangle{height: 2, base: 10}

	printArea(sq)
	printArea(tr)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (tr Triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}

func (sq Square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}
