package repository

import (
	"encoding/json"
	"io"
	"net/http"

	"posts_sender/internal/models"
)

// PostRepository інтерфейс для роботи з постами
type PostRepository interface {
	GetPosts() ([]models.Post, error)
	SavePosts(posts []models.Post) error
}

type jsonPlaceholderRepository struct {
	baseURL string
}

// NewJSONPlaceholderRepository створює новий екземпляр репозиторію JSONPlaceholder
func NewJSONPlaceholderRepository() PostRepository {
	return &jsonPlaceholderRepository{
		baseURL: "https://jsonplaceholder.typicode.com",
	}
}

// GetPosts отримує всі пости з JSONPlaceholder API
func (r *jsonPlaceholderRepository) GetPosts() ([]models.Post, error) {
	resp, err := http.Get(r.baseURL + "/posts")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var posts []models.Post
	if err := json.Unmarshal(body, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

// SavePosts заглушка для JSONPlaceholder репозиторію
func (r *jsonPlaceholderRepository) SavePosts(posts []models.Post) error {
	// JSONPlaceholder не підтримує реальне збереження
	return nil
}
