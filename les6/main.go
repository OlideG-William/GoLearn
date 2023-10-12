package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Student struct {
	Name     string
	Standard int `json:"Standard"`
}

func main() {
	StudentFile, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer StudentFile.Close()

	var studentDecoder *json.Decoder = json.NewDecoder(StudentFile)
	if err != nil {
		log.Fatal(err)
	}

	var studentList []Student

	err = studentDecoder.Decode(&studentList)
	if err != nil {
		log.Fatal(err)
	}

	for i, student := range studentList {
		fmt.Println("Student", i+1)
		fmt.Println("Student name", student.Name)
		fmt.Println("Student standart: ", student.Standard)
		fmt.Printf("\n")
	}
}
