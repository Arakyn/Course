package main

import "fmt"

func main() {
	x := struct {
		name    string
		age     int
		address map[string]string
	}{
		name: "Amandeep",
		age:  18,
		address: map[string]string{
			"Flat":   "48-c",
			"Colony": "Ashok Vihar",
		},
	}
	fmt.Println(x)
	for k, v := range x.address {
		fmt.Println(k, v)
	}
}
