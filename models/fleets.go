package models

import "time"

type Fleets struct {
	Fleet_ID  int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Max_Speed float64   `json:"max_speed"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
	DeletedAt time.Time `json:"deleted"`
}
