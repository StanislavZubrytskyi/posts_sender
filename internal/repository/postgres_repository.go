package repository

import (
	"database/sql"
	"posts_sender/internal/models"
)

type postgresRepository struct {
	db *sql.DB
}

// NewPostgresRepository створює новий екземпляр PostgreSQL репозиторію
func NewPostgresRepository(db *sql.DB) PostRepository {
	return &postgresRepository{
		db: db,
	}
}

// GetPosts отримує всі пости з бази даних
func (r *postgresRepository) GetPosts() ([]models.Post, error) {
	query := `
		SELECT id, user_id, title, body
		FROM posts
		ORDER BY id DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Body)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// SavePosts зберігає пости в базу даних
func (r *postgresRepository) SavePosts(posts []models.Post) error {
	query := `
		INSERT INTO posts (user_id, title, body)
		VALUES ($1, $2, $3)
		ON CONFLICT (id) DO UPDATE
		SET user_id = EXCLUDED.user_id,
			title = EXCLUDED.title,
			body = EXCLUDED.body
	`

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, post := range posts {
		_, err = stmt.Exec(post.UserID, post.Title, post.Body)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
