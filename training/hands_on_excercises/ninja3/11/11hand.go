package main

import "fmt"

func main() {
	FavSport := "FavSport"
	switch FavSport {
	case "LOL":
		fmt.Println("this should not print")
	case "FavSport":
		fmt.Println("This should Print")
	}
}
