package main

import "fmt"

func main() {
	x := map[string][]string{

		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	for k, v := range x { // idhar k is the key which is a string and v is a []string which is a slice
		fmt.Println("Key:", k, "Value", v)
		for i, val := range v { // []string slice ki range bata rahe v = string
			fmt.Printf("\t Index Position %v\t Value is %v\n", i, val)
		}
	}
}
