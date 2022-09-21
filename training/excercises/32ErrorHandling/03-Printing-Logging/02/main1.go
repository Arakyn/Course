package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("lol.txt")
	if err != nil {
		fmt.Println("err happaned", err)

	}
	log.SetOutput(f)
	defer f.Close()
	f2, err1 := os.Open("no-name.txt")
	if err1 != nil {
		log.Println("Err Happened", err1)
	}
	defer f2.Close()
	fmt.Println("Check the File in the directory")

}
