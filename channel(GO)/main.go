package main

import (
	"log"
	"github.com/Shobhitdimri01/CHANNEL/helpers"
)

const numpool = 10

func CalculateValue(intchan chan int) {
	randomnum := helpers.RandomNumber(numpool)
	intchan <- randomnum
}
func main() {
	intchan := make(chan int)
	defer close(intchan)
	go CalculateValue(intchan) //Concurrency
	num := <-intchan
	log.Println(num)
}
