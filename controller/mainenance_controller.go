package controller

import (
	"errors"
	"go-asset-management/service"
	"go-asset-management/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MaintenanceController struct {
	maintenanceService service.MaintenanceService
}

func NewMaintenanceController(maintenanceService service.MaintenanceService) *MaintenanceController {
	return &MaintenanceController{maintenanceService: maintenanceService}
}

func (ctrl *MaintenanceController) CreateMaintenance(c *gin.Context) {
	var input struct {
		AssetID     int     `json:"assetId" binding:"required"`
		UserID      int     `json:"userId" binding:"required"`
		Description string  `json:"description" binding:"required"`
		Cost        float64 `json:"cost" binding:"required"`
		StatusID    int     `json:"statusId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid input"))
		return
	}

	err := ctrl.maintenanceService.CreateMaintenance(input.AssetID, input.UserID, input.Description, input.Cost, input.StatusID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to create maintenance"))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Maintenance Created Successfully", nil))
}

func (ctrl *MaintenanceController) GetAllMaintenances(c *gin.Context) {
	maintenances, err := ctrl.maintenanceService.GetAllMaintenances()
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch maintenances"))
		return
	}
	c.JSON(http.StatusOK, util.NewSuccessResponse("Maintenances fetched successfully", maintenances))
}

func (ctrl *MaintenanceController) GetMaintenanceByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid maintenance ID"))
		return
	}

	maintenance, err := ctrl.maintenanceService.GetMaintenanceByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, util.NewFailedResponse("Maintenance not found"))
		} else {
			c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch maintenance"))
		}
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Maintenance fetched successfully", maintenance))
}

func (ctrl *MaintenanceController) UpdateMaintenance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid maintenance ID"))
		return
	}

	var input struct {
		Description string `json:"description" binding:"required"`
		StatusID    int    `json:"statusId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid input"))
		return
	}

	err = ctrl.maintenanceService.UpdateMaintenance(id, input.Description, input.StatusID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to update maintenance"))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Maintenance updated successfully", nil))
}

func (ctrl *MaintenanceController) DeleteMaintenance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid maintenance ID"))
		return
	}

	err = ctrl.maintenanceService.DeleteMaintenance(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, util.NewFailedResponse("Maintenance not found"))
		} else if err.Error() == "cannot delete completed maintenance" {
			c.JSON(http.StatusBadRequest, util.NewFailedResponse("Cannot delete completed maintenance"))
		} else {
			c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to delete maintenance"))
		}
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Maintenance deleted successfully", nil))
}

func (ctrl *MaintenanceController) GetTotalCost(c *gin.Context) {
	totalCost, err := ctrl.maintenanceService.GetTotalCost()
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to calculate total cost"))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Total cost calculated successfully", gin.H{"total_cost": totalCost}))
}

func (ctrl *MaintenanceController) GetTotalCostByAssetID(c *gin.Context) {
	assetID, err := strconv.Atoi(c.Param("asset_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid asset ID"))
		return
	}

	totalCost, err := ctrl.maintenanceService.GetTotalCostByAssetID(assetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to calculate total cost for asset"))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Total cost calculated successfully", gin.H{"total_cost": totalCost}))
}

func (ctrl *MaintenanceController) GetMaintenancesByUserID(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("User ID not found in context"))
		return
	}

	maintenances, err := ctrl.maintenanceService.GetMaintenancesByUserID(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch maintenances"))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Maintenances fetched successfully", maintenances))
}
