package main

import (
	"fmt"
	"time"
)

func main() {

	currTime := time.Now()
	fmt.Println("Current time:", currTime)

	fmt.Println("Formatted time:", currTime.Format("01-02-2006 15:04:05 Monday"))
	// well i m still surprised that why using this format !!!!

}
