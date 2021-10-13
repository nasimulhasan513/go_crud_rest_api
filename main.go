package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Struct thar represents a Blog post type
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author Author `json:"author"`
}

// Author struct represents an Author
type Author struct {
	FullName string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var posts []Post = []Post{}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/posts", addPost).Methods("POST")
	router.HandleFunc("/posts", getAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", patchPost).Methods("PATCH")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	http.ListenAndServe(":5000", router)

}
func getPost(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	var id, err = strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not  be parsed"))
		return
	}

	// error checking
	if id >= len(posts) {
		w.WriteHeader(400)
		w.Write([]byte("No Post found"))
		return
	}

	post := posts[id]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)

}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func addPost(w http.ResponseWriter, r *http.Request) {
	// get Post value from request body
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)
	posts = append(posts, newPost)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPost)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	// get id from request body
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Parsing error"))
		return
	}
	if id >= len(posts) {
		w.WriteHeader(400)
		w.Write([]byte("Post not found"))
		return
	}

	// get updated Post value from request body
	var updatedPost Post
	json.NewDecoder(r.Body).Decode(&updatedPost)

	posts[id] = updatedPost

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPost)

}

func patchPost(w http.ResponseWriter, r *http.Request) {
	// get id from request body
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Parsing error"))
		return
	}
	if id >= len(posts) {
		w.WriteHeader(400)
		w.Write([]byte("Post not found"))
		return
	}

	// get current Post value from array
	post := &posts[id]
	json.NewDecoder(r.Body).Decode(&post)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*post)

}
func deletePost(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Parsing error"))
		return
	}
	if id >= len(posts) {
		w.WriteHeader(400)
		w.Write([]byte("Post not found"))
		return
	}

	posts = append(posts[:id], posts[id+1:]...)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Messages string `json:"messages"`
	}{
		Messages: "Successfully deleted",
	})

}
