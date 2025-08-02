package main

import (
	"encoding/json"
	"fmt"
)

// you can also change the struct's field names to match the JSON keys
type course struct {
	Name     string   `json:"name"` //these backticks define and definately have lots of usage
	Price    int      `json:"price"`
	Domain   string   `json:"domain"`
	Rating   int      `json:"rating"`
	Duration string   `json:"duration"`
	Tags     []string `json:"tags"` //slice of strings
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
