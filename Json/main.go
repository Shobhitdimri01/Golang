package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Content string `json:"Content"`
	Summary string `json:"Summary"`
}
type red struct {
	Id      string
	Message string
}

var e red

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var post red
	post.Message = "Good" //hardcoded value

	json.Unmarshal(reqBody, &post)

	json.NewEncoder(w).Encode(post)

	newData, err := json.Marshal(post)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(newData))
	}
}

const a = ":8080"

func handleReqs() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/post", createNewArticle).Methods("POST")
	fmt.Println(fmt.Sprintf("Staring application on port %s", a))
	log.Fatal(http.ListenAndServe(a, r))
}

func main() {
	handleReqs()
}
