package main

import "fmt"


type bot interface{
	Getgreeting() string
}
type englishbot struct{}
type spanishbot struct{}

func main() {
	eb := englishbot{}
	sb := spanishbot{}
	PrintGret(eb)
	PrintGret(sb)
}

func PrintGret(b bot) {
	fmt.Println(b.Getgreeting())
}

func (eb englishbot) Getgreeting() string {
	return "Hi!"
}

func (sb spanishbot) Getgreeting() string {
	return "Ola!!"
}

