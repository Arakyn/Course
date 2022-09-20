package main

import (
	"os"
)

func main() {
	_, err := os.Open("no-name.txt") // this will give error because there is no file
	if err != nil {
		// fmt.Println("err happened", err)
		// log.Println("Err Happened with log", err)
		// log.Fatalln(err)
		panic(err)

		// fmt.Println( ) does the same shit but log ke sath it calls os.exit which exits the program
		// 2022/09/17 13:41:12 open no-name.txt: The system cannot find the file specified. ye error dega
	}
}
