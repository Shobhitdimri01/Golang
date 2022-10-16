package main

import ("fmt"
		"github.com/Shobhitdimri01/GoProgram/helper"
)

type Animal interface {
	says() string
	NumberofLegs() int
}
type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name  string
	Teeth int
	color string
}

func main() {
	var user helper.Someone 
	user.Person ="Ram"
	fmt.Println(user.Person)

	dog := Dog{
		Name:  "Tommy",
		Breed: "German-Shepherd",
	}
	gorilla:= Gorilla{
		Name:"Pik",
		color: "Black",
		Teeth: 65,
	}
	PrintInfo(&dog)
	PrintInfo( &gorilla)

}
func PrintInfo(a Animal) {
	fmt.Println("The animal says", a.says(), "and has", a.NumberofLegs(), "of legs")
}
func (d *Dog) says() string {
	return "Woof"
}
//In Go most of the reciever are pointer type it makes process fast
func (d *Dog) NumberofLegs() int {
	return 4
}

func(G *Gorilla) says()string{
	return "Anghaaa!!"
}
func(G *Gorilla) NumberofLegs()int{
	return 2
}