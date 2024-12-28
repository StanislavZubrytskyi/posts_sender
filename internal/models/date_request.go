package models

import "time"

// DateRequest представляє запит з датою
type DateRequest struct {
	TargetDate time.Time `json:"target_date"`
}

// TimeRemaining представляє відповідь з залишком часу
type TimeRemaining struct {
	Days    int `json:"days"`
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
	Seconds int `json:"seconds"`
}
