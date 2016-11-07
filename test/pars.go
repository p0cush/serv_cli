package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
)

func main() {
	//response, err := http.Get("http://golang.org/")
	response, err := http.Get("https://vk.com/id100478977")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(contents))
	}
}