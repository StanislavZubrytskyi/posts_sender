package service

import (
	"posts_sender/internal/models"
	"time"
)

// DateService інтерфейс для роботи з датами
type DateService interface {
	CalculateTimeRemaining(targetDate time.Time) models.TimeRemaining
}

type dateService struct{}

// NewDateService створює новий екземпляр сервісу дат
func NewDateService() DateService {
	return &dateService{}
}

// CalculateTimeRemaining обчислює залишок часу до вказаної дати
func (s *dateService) CalculateTimeRemaining(targetDate time.Time) models.TimeRemaining {
	now := time.Now()
	if now.After(targetDate) {
		return models.TimeRemaining{}
	}

	duration := targetDate.Sub(now)
	
	days := int(duration.Hours()) / 24
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	return models.TimeRemaining{
		Days:    days,
		Hours:   hours,
		Minutes: minutes,
		Seconds: seconds,
	}
}
