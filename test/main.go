package main

import (

	"net/http"
	"encoding/json"
	"src/github.com/gorilla/mux"
	"io/ioutil"
	"time"
	"fmt"
	"os"
)

type Json struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	input_data, _ := ioutil.ReadAll(r.Body)
	var (
		decoded_data = new(Json)
	)

	if json.Unmarshal( input_data, decoded_data) == nil {
		fmt.Println("- message received -")

		if decoded_data.Key == "com" {

			 if decoded_data.Value == "time" {
				 data_for_send := new(Json)
				 data_for_send.Key = "ans"
				 decoded_data.Value = time.Now().String()
				 out, _ := json.Marshal(decoded_data)
				 w.Write([]byte(out))
			 }
		}
		if decoded_data.Key == "--"{
			if decoded_data.Value == "destroy" {
				fmt.Println("exit")
				os.Exit(1)
			}
		}

	}
}

func main() {
	var router *mux.Router
	router = mux.NewRouter()


	router.HandleFunc("/net", Handler)
	http.ListenAndServe(":8080", router)
}



//----------------------------------------------------------------------------------

