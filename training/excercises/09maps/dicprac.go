package main

import "fmt"

func main() {

	superhero :=
		map[string]map[string]string{
			"Superman": map[string]string{

				"realname": "Clark Kent",
				"City":     "Metroplois",
				"Age":      "21",
			},

			"Batman": map[string]string{
				"realname": "Bruce Wayne",
				"City":     "Gotham City",
				"Age":      "23",
			},
		}
	fmt.Println("Superman"[Age])
	if temp, hero := superhero["Superman"]; hero {
		fmt.Println("Name =", temp["realname"], "City =", temp["City"], "Age =", temp["Age"])
	}
}
