package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var guess int
	var random int
	random = rand.Intn(100)
	fmt.Println("Guess a number between 0 and 100")
	fmt.Println(random)
	fmt.Scanf("%d", &guess)

	for guess != random {
		if guess > random {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Too low!")
		}
		fmt.Scanf("%d", &guess)
	}
}
