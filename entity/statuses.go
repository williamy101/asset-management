package entity

import "time"

type Statuses struct {
	StatusID   int       `gorm:"primaryKey;autoIncrement" json:"statusId"`
	StatusName string    `gorm:"type:varchar(100);not null" json:"statusName"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
