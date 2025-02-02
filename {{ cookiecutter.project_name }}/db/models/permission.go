package models

import "time"

type Permission struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
