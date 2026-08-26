package service

import (
	"time"
	"yoga.example/studio/internal/model"
)

func Age(r model.Record, now time.Time) time.Duration { return now.Sub(r.CreatedAt) }
func IsValidID(id string) bool                        { return len(id) >= 3 }
func RankStatus(status string) int {
	switch status {
	case "confirmed":
		return 3
	case "pending":
		return 2
	case "archived":
		return 1
	default:
		return 0
	}
}
