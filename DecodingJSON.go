package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	FirstName string
	LastName  string
	Details   struct {
		Height int
		Weight float32
	}
}

func main() {
	//Decoding to a struct
	var person People

	jsonString := `{"firstname":"Peter",
	"lastname":"Parker"}`

	err := json.Unmarshal([]byte(jsonString), &person)
	if err == nil {
		fmt.Println(person.FirstName)
		fmt.Println(person.LastName)
	} else {
		fmt.Println(err)
	}

	//Decoding to array
	var persons []People

	jsonString2 :=
		`[
	{
		"firstname":"Wei-Meng",
		"lastname":"Lee",
		"details": {
		"height":175,
		"weight":70.0
		}
	},
		{
		"firstname":"Mickey",
		"lastname":"Mouse",
		"details": {
		"height":105,
		"weight":85.5
		}
	}
	]`

	json.Unmarshal([]byte(jsonString2), &persons)
	for _, person := range persons {
		fmt.Println(person.FirstName)
		fmt.Println(person.LastName)
		fmt.Println(person.Details.Height)
		fmt.Println(person.Details.Weight)
	}
}
