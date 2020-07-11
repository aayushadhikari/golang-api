package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type blogPost struct {
	ID      int64  `json:"ID"`
	Title   string `json:"Title"`
	Content string `json:"Content"`
}

type blogPosts []blogPost

var myPosts = blogPosts{
	{
		ID:      1,
		Title:   "My Life",
		Content: "My Life is Good",
	},
	{
		ID:      2,
		Title:   "Your Life",
		Content: "Your Life is Good Too",
	},
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome Home")
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	//res, _ := json.Marshal(myPosts)
	json.NewEncoder(w).Encode(myPosts)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	var newPost blogPost
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(w, "Need title and content")
	}

	json.Unmarshal(reqBody, &newPost)
	myPosts = append(myPosts, newPost)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newPost)
}
