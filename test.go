package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var number = 41

	for {
		var n = rand.Intn(100) + 1
		if n < number {
			fmt.Printf("%v слишком маленькое число.\n", n)
		} else if n > number {
			fmt.Printf("%v слишком большое число.\n", n)
		} else {
			fmt.Printf("Угадал! %v\n", n)
			break
		}
	}
}
