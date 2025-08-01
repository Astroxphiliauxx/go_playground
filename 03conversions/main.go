package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter a number: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("You entered:", input)

	numRating, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64) //conversion of string into int64, trimming is necessary to avoid any errors.

	//error handling
	if err != nil {
		fmt.Println("Error converting input to number:", err)
		return
	}
	fmt.Println("You entered the number:", numRating+1)

}
