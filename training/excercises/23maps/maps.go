package main

import "fmt"

func main() {
	y := map[string]int{

		"Amandeep": 9871398314,
		"Divesh":   1233432453,
		"James":    23332435435,
	}
	y["Dad"] = 2123234

	fmt.Println(y)

	fmt.Println(y["Amandeep"])
	fmt.Println(y["Barny"])
	v, ok := y["Barny"]
	fmt.Println(ok, v)
	v1, ok1 := y["Amandeep"]
	fmt.Println(ok1, v1)

	for k, v := range y {
		fmt.Println(k, v)
	}

	intslice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for l, c := range intslice {
		fmt.Println(l, c)
	}

	// heroes := map[int]map[string]string{
	// 	1: {
	// 		"Name":     "Superman",
	// 		"Realname": "Clark Kent",
	// 	},
	// }
	// fmt.Println(heroes)

	// deleiting a value
	delete(y, "Dad")
	for v, k := range y {
		fmt.Println(v, k)
	}

	if v2, ok2 := y["Divesh"]; ok2 == true {
		fmt.Println("Value is", v2)
		delete(y, "Divesh")

	}
	for v, k := range y {
		fmt.Println(v, k)
	}
}
