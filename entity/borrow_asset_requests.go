package entity

import "time"

type BorrowAssetRequests struct {
	BorrowRequestID    int       `gorm:"primaryKey;autoIncrement" json:"borrowRequestId"`
	AssetID            int       `gorm:"not null" json:"assetId"`
	UserID             int       `gorm:"not null" json:"userId"`
	RequestDate        time.Time `gorm:"type:date;not null" json:"requestDate"`
	RequestedStartDate time.Time `gorm:"type:date;not null" json:"requestedStartDate"`
	RequestedEndDate   time.Time `gorm:"type:date;not null" json:"requestedEndDate"`
	StatusID           int       `gorm:"not null" json:"statusId"`
	ApprovedBy         *int      `json:"approvedBy"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	Assets             Assets    `gorm:"foreignKey:AssetID;references:AssetID" json:"assets"`
	Users              Users     `gorm:"foreignKey:UserID;references:UserID" json:"users"`
	Statuses           Statuses  `gorm:"foreignKey:StatusID;references:StatusID" json:"statuses"`
	Approver           *Users    `gorm:"foreignKey:ApprovedBy;references:UserID" json:"approver"`
}
