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

// **Create Maintenance**
func (ctrl *MaintenanceController) CreateMaintenance(c *gin.Context) {
	var input struct {
		RequestID   int     `json:"requestId" binding:"required"`
		Worker      int     `json:"worker" binding:"required"`
		Description string  `json:"description" binding:"required"`
		Cost        float64 `json:"cost" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid input"))
		return
	}

	err := ctrl.maintenanceService.CreateMaintenance(input.RequestID, input.Worker, input.Description, input.Cost)
	if err != nil {
		if err.Error() == "this asset is already under maintenance, complete it first before scheduling a new one" {
			c.JSON(http.StatusBadRequest, util.NewFailedResponse("This asset already has an active maintenance. Complete it first before scheduling a new one."))
			return
		}
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to create maintenance"))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Maintenance Created Successfully", nil))
}

// **Get All Maintenances**
func (ctrl *MaintenanceController) GetAllMaintenances(c *gin.Context) {
	maintenances, err := ctrl.maintenanceService.GetAllMaintenances()
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch maintenances"))
		return
	}
	c.JSON(http.StatusOK, util.NewSuccessResponse("Maintenances fetched successfully", maintenances))
}

// **Get Maintenance By ID**
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

// **Update Maintenance**
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
		if err.Error() == "invalid status ID, must be 3 (In Maintenance), 4 (Scheduled), or 5 (Completed)" {
			c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid status ID"))
			return
		}
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to update maintenance"))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Maintenance updated successfully", nil))
}

// **Delete Maintenance**
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

// **Get Total Maintenance Cost**
func (ctrl *MaintenanceController) GetTotalCost(c *gin.Context) {
	totalCost, err := ctrl.maintenanceService.GetTotalCost()
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to calculate total cost"))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Total cost calculated successfully", gin.H{"total_cost": totalCost}))
}

// **Get Total Cost by Asset ID**
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

// **Get Maintenance History by Worker ID**
func (ctrl *MaintenanceController) GetMaintenancesByWorkerID(c *gin.Context) {
	workerID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Worker ID not found in context"))
		return
	}

	maintenances, err := ctrl.maintenanceService.GetMaintenancesByWorkerID(workerID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch maintenances"))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Maintenances fetched successfully", maintenances))
}
