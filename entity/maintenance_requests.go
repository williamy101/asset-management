package entity

import "time"

type MaintenanceRequests struct {
	RequestID        int        `gorm:"primaryKey;autoIncrement" json:"requestId"`
	AssetID          int        `gorm:"not null" json:"assetId"`
	UserID           int        `gorm:"not null" json:"userId"`
	RequestDate      time.Time  `gorm:"type:date;not null" json:"requestDate"`
	IssueDescription string     `gorm:"type:text" json:"issueDescription"`
	StatusID         int        `gorm:"not null" json:"statusId"`
	DecisionDate     *time.Time `gorm:"type:date" json:"decisionDate"`
	MaintenanceDate  *time.Time `gorm:"type:date" json:"maintenanceDate"`
	CreatedAt        time.Time  `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt        time.Time  `gorm:"autoUpdateTime" json:"updatedAt"`
	Assets           Assets     `gorm:"foreignKey:AssetID;references:AssetID" json:"assets"`
	Users            Users      `gorm:"foreignKey:UserID;references:UserID" json:"users"`
	Statuses         Statuses   `gorm:"foreignKey:StatusID;references:StatusID" json:"statuses"`
}
