package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Структура для постів
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// API для отримання даних з JSONPlaceholder
func getPostsFromPlaceholder() ([]Post, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var posts []Post
	if err := json.Unmarshal(body, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

// Обробник API запиту
func getPostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	posts, err := getPostsFromPlaceholder()
	if err != nil {
		http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(posts)
}

func main() {
	// Реєстрація маршруту
	http.HandleFunc("/posts", getPostsHandler)

	// Запуск сервера
	port := ":8080"

	log.Fatal(http.ListenAndServe(port, nil))
}
