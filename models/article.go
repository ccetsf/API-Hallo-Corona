package models

import "time"

type Article struct {
	Id          int          `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId      int          `json:"user_id"`
	User        UserResponse `json:"user"`
	Title       string       `json:"title" gorm:"type: varchar(255)"`
	Attache     string       `json:"attache" gorm:"type: varchar(255)"`
	Description string       `json:"description" gorm:"type: text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
