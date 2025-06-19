package models

import "time"

type Group struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	IsPrivate bool      `json:"is_private"` // true = 1-on-1, false = public/group chat
	CreatedAt time.Time `json:"created_at"`
}