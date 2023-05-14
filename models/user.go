package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `gorm:"not null" json:"username"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Photos    []Photo   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

type Photo struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `gorm:"not null" json:"photoUrl"`
	UserID    uint      `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
