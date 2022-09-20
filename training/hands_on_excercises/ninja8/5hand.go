package main

import (
	"fmt"
	"sort"
)

type ByName []user

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].First < bn[j].First }

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)
	sort.Sort(ByAge(users))
	fmt.Println(users)
	sort.Sort(ByName(users))
	for i, v := range users {
		fmt.Println("\n", i, v.First, v.Last, v.Age, "\t")
		sort.Strings(v.Sayings)
		for _, v1 := range v.Sayings {
			fmt.Println("\n\t", v1)
		}
	}
	sort.Sort(ByAge(users))
	for i, v := range users {
		fmt.Println("\n", i, v.First, v.Last, v.Age, "\t")
		sort.Strings(v.Sayings)
		for _, v1 := range v.Sayings {
			fmt.Println("\n\t", v1)
		}
	}

	// your code goes here

}
