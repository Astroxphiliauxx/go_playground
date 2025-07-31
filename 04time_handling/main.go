package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	currTime := time.Now()
	fmt.Println("Current time:", currTime)

	fmt.Println("Formatted time:", currTime.Format("01-02-2006 15:04:05 Monday"))
	// well i m still surprised that why using this format !!!!

	diceNumber := rand.Intn(6) + 1 // generates a random number between 1 and 6
	fmt.Println("You rolled a dice and got:", diceNumber)

}
