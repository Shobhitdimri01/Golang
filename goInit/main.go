package main

import "fmt"

var WhatIsThe = AnswerToLife()

func AnswerToLife() int { // 1
	return 42
}

func init() { // 2
	WhatIsThe = 42
}

func main() { // 3
	if WhatIsThe == 42 {
		fmt.Println("It's not all a lie.")
	}
}