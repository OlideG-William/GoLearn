package main

import (
	"encoding/json"
	"fmt"

	//"github.com/tidwall/sjson"
)


type Messages struct{
	Name string
	Body string
	Time int64
}





func main() {


	m:= Messages{"Alice", "Hello", 33223423323432}
	b, err := json.Marshal(m)
	if err!= nil {
		panic("error")
	}
	fmt.Println(string(b))

	/*var m Messages
	err:= json.Unmarshal(b, &m)*/
	
}
