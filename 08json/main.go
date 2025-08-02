package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string
	Price    int
	Domain   string
	Rating   int
	Duration string
	Tags     []string //slice of strings
}

func main() {
	fmt.Println("Welcome to the JSON handling in Go!")
	fmt.Println("This is the unstructured JSON encoding example:")
	EncodeJsonUnstructured()
	fmt.Println("This is the structured JSON encoding example:")
	EncodeJsonStructured()
}

func EncodeJsonUnstructured() {

	mycourse := []course{
		{"Go Programming", 100, "Programming", 5, "4 weeks", []string{"Go", "Programming", "Backend"}},
		{"Python Programming", 80, "Programming", 4, "6 weeks", []string{"Python", "Programming", "Data Science"}},
		{"JavaScript Programming", 90, "Programming", 4, "5 weeks", []string{"JavaScript", "Programming", "Web Development"}},
	}

	finalData, _ := json.Marshal(mycourse) //converts the struct to json format (which is unstructured)
	fmt.Println(string(finalData))
}

// uses MarshalIndent to format the Json output with indentation
func EncodeJsonStructured() {

	mycourse := []course{
		{"Go Programming", 100, "Programming", 5, "4 weeks", []string{"Go", "Programming", "Backend"}},
		{"Python Programming", 80, "Programming", 4, "6 weeks", []string{"Python", "Programming", "Data Science"}},
		{"JavaScript Programming", 90, "Programming", 4, "5 weeks", []string{"JavaScript", "Programming", "Web Development"}},
	}

	//MarshalIndent uses 3 params, middle one is kinda useless, leave blank. The last one is the indentation string
	finalData, _ := json.MarshalIndent(mycourse, "", "\t")
	fmt.Println(string(finalData))
}
