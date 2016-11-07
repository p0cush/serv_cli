package main

import (

	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"os"
)

type Json struct {
	Key string `json:"key"`
	Value string `json:"value"`
}


func main() {

	res := new(Json)
	fmt.Scanln(&res.Key)
	fmt.Scanln(&res.Value)

	b, _ := json.Marshal(res)
	req, err := http.NewRequest("net", "http://127.0.0.1:8080/net", bytes.NewBuffer(b))

	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}


	input_data, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	if input_data == nil {
		fmt.Println("goodbye")
		os.Exit(1)
	}
	dat, err := ioutil.ReadAll(input_data.Body)

	decoded_data := new(Json)

	json.Unmarshal(dat, decoded_data)
	fmt.Printf("%#v", decoded_data.Value)
}