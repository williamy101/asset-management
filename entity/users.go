package entity

import "time"

type Users struct {
	UserID    int       `gorm:"primaryKey;autoIncrement" json:"userId"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	RoleID    int       `gorm:"not null" json:"roleId"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	Roles     Roles     `gorm:"foreignKey:RoleID;references:RoleID" json:"roles"`
}

type UserDTO struct {
	UserID int    `json:"userId"`
	Name   string `json:"name"`
	RoleID int    `json:"roleId"`
}
