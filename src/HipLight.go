package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"strings"
	"io/ioutil"
	"log"
	//"bytes"
)

func bodytostruct(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse arguments, you have to call this by yourself
	resp, err := ioutil.ReadAll(r.Body)
	var hip *HipMsg
	oops := json.Unmarshal(resp, &hip)
	if oops != nil {
		fmt.Println("error:", oops)
	}
	fmt.Println(hip)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func main() {

	http.HandleFunc("/", bodytostruct)       // set router
	err := http.ListenAndServe(":8050", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

type HipMsg struct {
	user    string // 'json:"mention_name"'
	message string // 'json:"message"'
}
