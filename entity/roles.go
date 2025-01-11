package entity

import "time"

type Roles struct {
	RoleID    int       `gorm:"primaryKey;autoIncrement" json:"roleId"`
	RoleName  string    `gorm:"type:varchar(100);not null" json:"roleName"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
