package entity

import "time"

type Assets struct {
	AssetID         int             `gorm:"primaryKey;autoIncrement" json:"assetId"`
	AssetName       string          `gorm:"type:varchar(100);not null" json:"assetName"`
	CategoryID      int             `json:"categoryId"`
	StatusID        int             `gorm:"not null" json:"statusId"`
	LastMaintenance *time.Time      `gorm:"type:date" json:"lastMaintenance"`
	NextMaintenance *time.Time      `gorm:"type:date" json:"nextMaintenance"`
	CreatedAt       time.Time       `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt       time.Time       `gorm:"autoUpdateTime" json:"updatedAt"`
	AssetCategories AssetCategories `gorm:"foreignKey:CategoryID;references:CategoryID" json:"assetCategories"`
	Statuses        Statuses        `gorm:"foreignKey:StatusID;references:StatusID" json:"statuses"`
}
