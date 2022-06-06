package models

import "time"

type User struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	FirtsName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
