package service

import (
	"posts_sender/internal/models"
	"posts_sender/internal/repository"
)

// PostService інтерфейс для бізнес-логіки роботи з постами
type PostService interface {
	GetPosts() ([]models.Post, error)
	SyncPosts() error
}

type postService struct {
	dbRepo    repository.PostRepository
	apiRepo   repository.PostRepository
}

// NewPostService створює новий екземпляр сервісу постів
func NewPostService(dbRepo, apiRepo repository.PostRepository) PostService {
	return &postService{
		dbRepo:  dbRepo,
		apiRepo: apiRepo,
	}
}

// GetPosts отримує всі пости з бази даних
func (s *postService) GetPosts() ([]models.Post, error) {
	return s.dbRepo.GetPosts()
}

// SyncPosts синхронізує пости з API в базу даних
func (s *postService) SyncPosts() error {
	posts, err := s.apiRepo.GetPosts()
	if err != nil {
		return err
	}

	return s.dbRepo.SavePosts(posts)
}
