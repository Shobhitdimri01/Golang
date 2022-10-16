package main

import "fmt"

type shape interface {
	getarea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sidelength float64
}

func main() {
	t:= triangle{2,3}
	s:= square{4}
	printarea(t)
	printarea(s)
}

func (t triangle) getarea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getarea() float64 {
	return s.sidelength * s.sidelength
}

func printarea(s shape){
	fmt.Println(s.getarea())
}