package main

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(data interface{}) {
	var p []byte
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	student := Student{
		ID:   1,
		Name: "Cao Thanh Tung",
	}
	fmt.Println(student)
	PrettyPrint(student)
}
