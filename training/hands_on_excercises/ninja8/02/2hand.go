package main

import (
	"encoding/json"
	"fmt"
)

type PEOPLE struct {
	First   string   // `json:"First"`
	Last    string   // `json:"Last"`
	Age     int      //`json:"Age"`
	Sayings []string //`json:"Sayings"`
}

func main() {
	var people []PEOPLE

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	bs := []byte(s)
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	for i, v := range people {

		fmt.Println("The Index position", i, "\n Has Value", v.First, v.Last, v.Age, "\n")
		for _, v1 := range v.Sayings {
			fmt.Println("\t", v1)
		}

	}
}
