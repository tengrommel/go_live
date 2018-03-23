package main

import (
	"encoding/json"
	"fmt"
)

type MyObject struct {
	Number int `json:"number"`
	Word string `json:"string"`
}


func main() {
	jsonBytes := []byte(`{"number": 5, "string": "Packt"}`)
	var dangerousObject map[string]interface{}
	err := json.Unmarshal(jsonBytes, &dangerousObject)
	if err != nil{
		panic(err)
	}
	fmt.Printf("Number is %d, Word is %s\n", dangerousObject["number"], dangerousObject["string"])
}
