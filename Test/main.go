package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(100.0, 0.0)
	if err != nil {
		log.Println(err)
		return
	} else {
		log.Println(result)
		return
	}

}

func divide(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("can't divide by zero")
	}
	result = x / y
	return result, nil

}
