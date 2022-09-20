package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := `[
		{"First":"Aman","Last":"Singh","Age":41},
		{"First":"Miss","Last":"Moneypenny","Age":31}
]`
	bs := []byte(s)
	type people struct {
		First string
		Last  string
		Age   int
	}
	var persons []people
	err := json.Unmarshal(bs, &persons)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(persons)
	for i, v := range persons {
		fmt.Println("The Index", i, "\nTHe value", v.First, v.Last, v.Age)

	}

}
