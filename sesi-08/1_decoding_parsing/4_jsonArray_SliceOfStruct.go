package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	var jsonString = `
		[
			{
				"full_name": "Iron man",
				"email" : "ironman@mail.com",
				"age" : 70
			},
			{
				"full_name": "Spiderman",
				"email" : "spiderman@mail.com",
				"age" : 70
			},
			{
				"full_name": "Hulk",
				"email" : "hulk@mail.com",
				"age" : 70
			}
		]
	`

	var result []Employee

	var err = json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		// fmt.Println(err.Error())
		fmt.Printf("error %s", err.Error())
		return
	}

	for i, v := range result {
		fmt.Printf("Index %d: %v\n", i+1, v)
	}

}
