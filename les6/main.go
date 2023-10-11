package main

import (
	"fmt"

	"github.com/tidwall/sjson"
)

func main() {

	json := `{
				"age":37,
				"children": ["Sara","Alex","Jack"],
				"friends": [
				  {"age": 44, "first": "Dale", "last": "Murphy"},
				  {"age": 68, "first": "Roger", "last": "Craig"},
				  {"age": 47, "first": "Jane", "last": "Murphy"}
				],
				"name":{"first":"Tom","last":"Anderson"},
				"age": 1631 
		  }`

	Val, _ := sjson.Set(json, "name.first", "Filson")
	fmt.Println(Val)

}
