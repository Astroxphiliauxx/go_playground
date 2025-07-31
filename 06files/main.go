package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("files.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	content := "This sample of file is created using GO programming language.\n"
	sizeOfFile, err := io.WriteString(file, content)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("File 'files.txt' created successfully with %d bytes written.\t", sizeOfFile)

	fmt.Println("Content of the file is:", readFile("files.txt"))

}

func readFile(sname string) string {

	sol, err := os.ReadFile(sname)

	if err != nil {
		panic(err)
	}
	return string(sol)

}
