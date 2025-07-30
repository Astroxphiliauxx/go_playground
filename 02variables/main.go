package main

import "fmt"

import (
	"bufio"
	"os"
)

var loginToken string = "abc123xyz"

func main() {

	var username string = "astroxphiliauxx"

	fmt.Println("Hello, " + username + "!  ")

	numberOfGophers := 20000 //you cant define the type as well as global variable outside the main function

	fmt.Println("There are", numberOfGophers, "Gophers in the world!")
	fmt.Println("Your login token is:", loginToken)

	//learn to take input from the user
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)

}
