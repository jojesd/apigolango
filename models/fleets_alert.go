package models

import "time"

type Fleets_Alert struct {
	Fleet_Alert_ID int       `json:"id" gorm:"primaryKey"`
	Fleet_ID       int       `json:"fleet_id" gorm:"foreignKey"`
	WebHook        string    `json:"webhook"`
	CreatedAt      time.Time `json:"created"`
	UpdatedAt      time.Time `json:"updated"`
	DeletedAt      time.Time `json:"deleted"`
}
