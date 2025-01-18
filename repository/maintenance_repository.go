package repository

import (
	"go-asset-management/entity"

	"gorm.io/gorm"
)

type MaintenanceRepository interface {
	Create(maintenance *entity.Maintenances) error
	FindAll() ([]entity.Maintenances, error)
	FindByID(maintenanceID int) (*entity.Maintenances, error)
	Update(maintenance *entity.Maintenances) error
	Delete(maintenanceID int) error
	CalculateTotalCost() (float64, error)
	GetTotalCostByAssetID(assetID int) (map[string]interface{}, error)
	FindByUserID(userID int) ([]entity.Maintenances, error)
	FindByAssetID(assetID int) ([]entity.Maintenances, error)
}

type maintenanceRepository struct {
	db *gorm.DB
}

func NewMaintenanceRepository(db *gorm.DB) MaintenanceRepository {
	return &maintenanceRepository{db: db}
}

func (r *maintenanceRepository) Create(maintenance *entity.Maintenances) error {
	return r.db.Create(maintenance).Error
}

func (r *maintenanceRepository) FindAll() ([]entity.Maintenances, error) {
	var maintenances []entity.Maintenances
	err := r.db.
		Preload("Assets").
		Preload("Assets.AssetCategories").
		Preload("Assets.Statuses").
		Preload("Users").
		Preload("Users.Roles").
		Preload("Statuses").
		Find(&maintenances).Error
	if err != nil {
		return nil, err
	}
	return maintenances, nil
}
func (r *maintenanceRepository) FindByID(maintenanceID int) (*entity.Maintenances, error) {
	var maintenance entity.Maintenances
	err := r.db.
		Preload("Assets").
		Preload("Assets.AssetCategories").
		Preload("Assets.Statuses").
		Preload("Users").
		Preload("Users.Roles").
		Preload("Statuses").
		Where("maintenance_id = ?", maintenanceID).
		Find(&maintenance).Error
	if err != nil {
		return nil, err
	}
	return &maintenance, nil
}

func (r *maintenanceRepository) Update(maintenance *entity.Maintenances) error {
	return r.db.Save(maintenance).Error
}

func (r *maintenanceRepository) Delete(maintenanceID int) error {
	return r.db.Delete(&entity.Maintenances{}, maintenanceID).Error
}

func (r *maintenanceRepository) CalculateTotalCost() (float64, error) {
	var totalCost float64
	err := r.db.Model(&entity.Maintenances{}).Select("SUM(cost)").Scan(&totalCost).Error
	if err != nil {
		return 0, err
	}
	return totalCost, nil
}

func (r *maintenanceRepository) GetTotalCostByAssetID(assetID int) (map[string]interface{}, error) {
	var result struct {
		AssetID   int     `json:"assetId"`
		AssetName string  `json:"assetName"`
		TotalCost float64 `json:"totalCost"`
	}

	query := `
        SELECT a.asset_id AS AssetID, a.asset_name AS AssetName, COALESCE(SUM(m.cost), 0) AS TotalCost
        FROM assets a
        LEFT JOIN maintenances m ON a.asset_id = m.asset_id
        WHERE a.asset_id = ?
        GROUP BY a.asset_id, a.asset_name
    `

	err := r.db.Raw(query, assetID).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"assetId":   result.AssetID,
		"assetName": result.AssetName,
		"totalCost": result.TotalCost,
	}, nil
}

func (r *maintenanceRepository) FindByUserID(userID int) ([]entity.Maintenances, error) {
	var maintenances []entity.Maintenances
	err := r.db.Preload("Assets").Preload("Assets.AssetCategories").Preload("Assets.Statuses").
		Preload("Users").Preload("Users.Roles").
		Preload("Statuses").
		Where("user_id = ?", userID).
		Find(&maintenances).Error
	if err != nil {
		return nil, err
	}
	return maintenances, nil
}

func (r *maintenanceRepository) FindByAssetID(assetID int) ([]entity.Maintenances, error) {
	var maintenances []entity.Maintenances
	err := r.db.Preload("Assets").Preload("Assets.AssetCategories").Preload("Assets.Statuses").
		Preload("Users").Preload("Users.Roles").
		Preload("Statuses").
		Where("asset_id = ?", assetID).
		Find(&maintenances).Error
	if err != nil {
		return nil, err
	}
	return maintenances, nil
}
