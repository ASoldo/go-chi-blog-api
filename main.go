package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Post struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Likes    int       `json:"likes"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

var (
	posts     = make(map[int]Post)
	postID    = 1
	commentID = 1
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/posts", getAllPosts)
	r.Get("/posts/{id}", getPostByID)
	r.Post("/posts", createPost)
	r.Put("/posts/{id}", updatePostByID)
	r.Delete("/posts/{id}", deletePostByID)
	r.Post("/posts/{id}/like", likePostByID)
	r.Post("/posts/{id}/comment", commentOnPost)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to listend and serve: %v", err)
	}
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	postsList := []Post{}
	for _, post := range posts {
		postsList = append(postsList, post)
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(postsList); err != nil {
		log.Fatalf("Unable to parse the JSON. %v", err)
	}
}

func getPostByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	post, ok := posts[id]
	if !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(post); err != nil {
		log.Fatalf("Unable to parse JSON. %v", err)
	}
}

func createPost(w http.ResponseWriter, r *http.Request) {
	var post Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	post.ID = postID
	posts[post.ID] = post
	postID++
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(post); err != nil {
		log.Fatalf("Unable to parse JSON. %v", err)
	}
}

func updatePostByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	post, ok := posts[id]
	if !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	var updatedPost Post
	if err := json.NewDecoder(r.Body).Decode(&updatedPost); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	post.Title = updatedPost.Title
	post.Content = updatedPost.Content
	posts[id] = post
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(post); err != nil {
		log.Fatalf("Unable to parse JSON. %v", err)
	}
}

func deletePostByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if _, ok := posts[id]; !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	delete(posts, id)
	w.WriteHeader(http.StatusNoContent)
}

func likePostByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	post, ok := posts[id]
	if !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	post.Likes++
	posts[id] = post
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(post); err != nil {
		log.Fatalf("Unable to parse JSON. %v", err)
	}
}

func commentOnPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	post, ok := posts[id]
	if !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	var comment Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	comment.ID = commentID
	commentID++
	post.Comments = append(post.Comments, comment)
	posts[id] = post
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(post); err != nil {
		log.Fatalf("Unable to parse JSON. %v", err)
	}
}
