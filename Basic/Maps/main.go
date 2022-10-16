package main

import "fmt"

type name string
type age int
var Person map[name]age
func main() {

	Person = map[name]age{
		"Rohit":   22,
		"Shubham": 25,
		"Mayank":  21,
		"Rahul":   31,
	}
	dprint()
}

func dprint(){
	for key,value := range Person{
		fmt.Println(key,"--->",value)
	}
}