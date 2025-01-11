package entity

import "time"

type Maintenances struct {
	MaintenanceID int       `gorm:"primaryKey;autoIncrement" json:"maintenanceId"`
	AssetID       int       `gorm:"not null" json:"assetId"`
	UserID        int       `gorm:"not null" json:"userId"`
	Description   string    `gorm:"type:text" json:"description"`
	Cost          float64   `gorm:"type:decimal(10,2)" json:"cost"`
	StatusID      int       `gorm:"not null" json:"statusId"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	Assets        Assets    `gorm:"foreignKey:AssetID" json:"assets"`
	Users         Users     `gorm:"foreignKey:UserID" json:"users"`
	Statuses      Statuses  `gorm:"foreignKey:StatusID" json:"statuses"`
}
