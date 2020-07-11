package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type blogPost struct {
	Id      int64  `json:"Id"`
	Title   string `json:"Title"`
	Content string `json:"Content"`
}

type blogPosts []blogPost

var myPosts = blogPosts{
	{
		Id:      1,
		Title:   "My Life",
		Content: "My Life is Good",
	},
	{
		Id:      2,
		Title:   "Your Life",
		Content: "Your Life is Good Too",
	},
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/posts", getPosts)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome Home")
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	//res, _ := json.Marshal(myPosts)
	json.NewEncoder(w).Encode(myPosts)
}
