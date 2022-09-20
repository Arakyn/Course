// creating maps in go
package main

import "fmt"

func main() {
	// studentage := make(map[string]int)
	// studentage["Arya"] = 23
	// studentage["Aman"] = 18
	// studentage["Yash"] = 15
	// studentage["Lakshit"] = 16
	// studentage["Jagmet"] = 17
	// studentage["jog"] = 132
	// fmt.Println(studentage)
	// fmt.Println(studentage["Aman"])

	// this will give this output
	// =========================================================
	//map[Aman:18 Arya:23 Jagmet:17 Lakshit:16 Yash:15 jog:132]
	// 18
	//==========================================================

	// studentage := make(map[int]string)
	// studentage[1] = "Aman"
	// studentage[2] = "Kirti"
	// studentage[3] = "Rahul"
	// studentage[4] = "Jogesh"
	// studentage[5] = "Sachin"
	// studentage[6] = "BIVHT"

	// fmt.Println(studentage[2], studentage[1], studentage[6])
	//	Kirti Aman BIVHT

	// fmt.Println(len(studentage))
	//  6

	superhero
	map[string]map[string]string{

		"Superman": map[string]string{
			"realname": "Clark Kent",
			"city":     "Delhi",
		},

		"Batman": map[string]string{

			"realname": "Bruce Wayne",
			"city":     "Gotham City",
		},
	}
	if temp, hero := temp["Superman"]; hero {
		fmt.Println(hero["realname"], hero["city"])
	}

}
