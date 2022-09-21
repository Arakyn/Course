package main

import "fmt"

func main() {
	x := map[string][]string{

		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	x["khurana_Aman"] = []string{`Games`, `More Games`, `idk he dumb`}
	for k, v := range x {
		fmt.Println(k, v)
	}
}
