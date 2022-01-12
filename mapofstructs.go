package main

import (
	"fmt"
	"sort"
)

type dob struct {
	day   int
	month int
	year  int
}

type people struct {
	name  string
	email string
	dob   dob
}

var members map[int]people

func main() {
	members := make(map[int]people)

	members[1] = people{
		name:  "Mary Smith",
		email: "marysmith@example.com",
		dob: dob{
			day:   17,
			month: 3,
			year:  1990,
		},
	}
	members[2] = people{
		name:  "John Smith",
		email: "johnsmith@example.com",
		dob: dob{
			day:   9,
			month: 12,
			year:  1988,
		},
	}
	members[3] = people{
		name:  "Janet Doe",
		email: "janetdoe@example.com",
		dob: dob{
			day:   1,
			month: 12,
			year:  1988,
		},
	}
	members[4] = people{
		name:  "Adam Jones",
		email: "adamjones@example.com",
		dob: dob{
			day:   19,
			month: 8,
			year:  2001,
		},
	}

	for k, v := range members {
		fmt.Println(k, v.name, v.email, v.dob)
	}

	var keys []int
	for k := range members {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	var sliceMembers []people
	for _, k := range keys {
		sliceMembers = append(sliceMembers, members[k])
	}

	sort.SliceStable(sliceMembers, func(i, j int) bool {
		if sliceMembers[i].dob.year !=
			sliceMembers[j].dob.year {
			return sliceMembers[i].dob.year <
				sliceMembers[j].dob.year
		}
		if sliceMembers[i].dob.month !=
			sliceMembers[j].dob.month {
			return sliceMembers[i].dob.month <
				sliceMembers[j].dob.month
		}
		return sliceMembers[i].dob.day <
			sliceMembers[j].dob.day
	})

	for _, v := range sliceMembers {
		fmt.Println(v.name, v.email, v.dob)
	}
}
