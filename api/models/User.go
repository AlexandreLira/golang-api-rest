package models

import "time"

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:20;not null;unique" json:"nickname"`
	Email     string    `gorm:"size:20;not null;unique" json:"email"`
	Password  string    `gorm:"size:60;not null" json:"password"`
	CreateAt  time.Time `gorm:"default:current_timestamp()" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:current_timestamp()" json:"updateAt"`
}
