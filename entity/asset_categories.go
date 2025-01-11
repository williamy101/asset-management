package entity

import "time"

type AssetCategories struct {
	CategoryID   int       `gorm:"primaryKey;autoIncrement" json:"categoryId"`
	CategoryName string    `gorm:"type:varchar(100);not null" json:"categoryName"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
