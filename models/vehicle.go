package models

import "time"

type Vehicle struct {
	Vehicle_ID int       `json:"id" gorm:"primaryKey"`
	Fleet_ID   int       `json:"fleet_id" gorm:"foreignKey"`
	Name       string    `json:"name"`
	Max_Speed  float64   `json:"veloc"`
	CreatedAt  time.Time `json:"created"`
	UpdatedAt  time.Time `json:"updated"`
	DeletedAt  time.Time `json:"deleted"`
}
