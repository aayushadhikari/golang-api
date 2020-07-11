package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type blogPost struct {
	title   string
	content string
}

type blogPosts []blogPost

var myPosts = blogPosts{
	{
		title:   "My Life",
		content: "My Life is Good",
	},
	{
		title:   "Your Life",
		content: "Your Life is Good Too",
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
	json.NewEncoder(w).Encode(myPosts)
}
