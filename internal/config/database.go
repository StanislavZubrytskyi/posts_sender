package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// NewDBConnection створює нове підключення до бази даних
func NewDBConnection() (*sql.DB, error) {
	// Отримуємо параметри підключення з змінних середовища
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "posts_db")

	// Формуємо рядок підключення
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Відкриваємо з'єднання
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Перевіряємо з'єднання
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// getEnv отримує значення змінної середовища або повертає значення за замовчуванням
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
