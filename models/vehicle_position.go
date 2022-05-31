package models

import "time"

type Vehicle_Position struct {
	Vehicle_Position_ID int       `json:"vp_id" gorm:"primaryKey"`
	Vehicle_ID          int       `json:"vehicle_id" gorm:"foreignKey"`
	Timestamp           int       `json:"timestamp"`
	Latitude            int       `json:"latitude"`
	Longitude           int       `json:"longitude"`
	Current_Speed       int       `json:"current_speed"`
	Max_Speed           float64   `json:"max_speed"`
	CreatedAt           time.Time `json:"created"`
	UpdatedAt           time.Time `json:"updated"`
	DeletedAt           time.Time `json:"deleted"`
}
