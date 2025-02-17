package entity

import "time"

type BorrowedAssets struct {
	BorrowID        int                 `gorm:"primaryKey;autoIncrement" json:"borrowId"`
	AssetID         int                 `gorm:"not null" json:"assetId"`
	UserID          int                 `gorm:"not null" json:"userId"`
	BorrowRequestID int                 `gorm:"not null" json:"borrowRequestId"`
	BorrowDate      time.Time           `gorm:"type:date;not null" json:"borrowDate"`
	ReturnDate      *time.Time          `gorm:"type:date" json:"returnDate"`
	StatusID        int                 `gorm:"not null" json:"statusId"`
	CreatedAt       time.Time           `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt       time.Time           `gorm:"autoUpdateTime" json:"updatedAt"`
	Assets          Assets              `gorm:"foreignKey:AssetID;references:AssetID" json:"assets"`
	Users           Users               `gorm:"foreignKey:UserID;references:UserID" json:"users"`
	Statuses        Statuses            `gorm:"foreignKey:StatusID;references:StatusID" json:"statuses"`
	BorrowRequest   BorrowAssetRequests `gorm:"foreignKey:BorrowRequestID;references:BorrowRequestID" json:"borrowRequest"`
}
