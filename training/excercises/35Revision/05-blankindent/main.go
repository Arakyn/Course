package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Here we will be using the func http.Get
	// to get the content of a web page.

	x, _ := http.Get("http://www.google.com")
	// here after x we are using a blank indentififer
	fmt.Println(x)
	page, _ := ioutil.ReadAll(x.Body)
	x.Body.Close()
	fmt.Println(string(page))
}
