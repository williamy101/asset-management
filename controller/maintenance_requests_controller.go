package controller

import (
	"errors"
	"go-asset-management/service"
	"go-asset-management/util"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MaintenanceRequestController struct {
	maintenanceRequestService service.MaintenanceRequestService
}

func NewMaintenanceRequestController(maintenanceRequestService service.MaintenanceRequestService) *MaintenanceRequestController {
	return &MaintenanceRequestController{maintenanceRequestService: maintenanceRequestService}
}

// **Create a new maintenance request**
func (ctrl *MaintenanceRequestController) CreateMaintenanceRequest(c *gin.Context) {
	var input struct {
		AssetID          int    `json:"assetId" binding:"required"`
		IssueDescription string `json:"issueDescription" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid input"))
		return
	}

	// Get user ID from context (assumed to be set by authentication middleware)
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, util.NewFailedResponse("Unauthorized"))
		return
	}

	err := ctrl.maintenanceRequestService.CreateMaintenanceRequest(input.AssetID, userID.(int), input.IssueDescription)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Maintenance request created successfully", nil))
}

// **Get all maintenance requests**
func (ctrl *MaintenanceRequestController) GetAllMaintenanceRequests(c *gin.Context) {
	requests, err := ctrl.maintenanceRequestService.GetAllMaintenanceRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch maintenance requests"))
		return
	}
	c.JSON(http.StatusOK, util.NewSuccessResponse("Maintenance requests fetched successfully", requests))
}

// **Get a maintenance request by ID**
func (ctrl *MaintenanceRequestController) GetMaintenanceRequestByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid request ID"))
		return
	}

	request, err := ctrl.maintenanceRequestService.GetMaintenanceRequestByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, util.NewFailedResponse("Maintenance request not found"))
		} else {
			c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch maintenance request"))
		}
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Maintenance request fetched successfully", request))
}

// **Get maintenance requests made by a specific user**
func (ctrl *MaintenanceRequestController) GetMaintenanceRequestsByUserID(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, util.NewFailedResponse("Unauthorized"))
		return
	}

	requests, err := ctrl.maintenanceRequestService.GetMaintenanceRequestsByUserID(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch maintenance requests"))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Maintenance requests fetched successfully", requests))
}

// **Get maintenance requests related to a specific asset**
func (ctrl *MaintenanceRequestController) GetMaintenanceRequestsByAssetID(c *gin.Context) {
	assetID, err := strconv.Atoi(c.Param("assetId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid asset ID"))
		return
	}

	requests, err := ctrl.maintenanceRequestService.GetMaintenanceRequestsByAssetID(assetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch maintenance requests"))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Maintenance requests fetched successfully", requests))
}

// **Get maintenance requests filtered by status**
func (ctrl *MaintenanceRequestController) GetMaintenanceRequestsByStatus(c *gin.Context) {
	statusID, err := strconv.Atoi(c.Param("statusId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid status ID"))
		return
	}

	requests, err := ctrl.maintenanceRequestService.GetMaintenanceRequestsByStatus(statusID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch maintenance requests"))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Maintenance requests fetched successfully", requests))
}

func (ctrl *MaintenanceRequestController) ApproveMaintenanceRequest(c *gin.Context) {
	var input struct {
		Worker          int     `json:"worker" binding:"required"`
		Description     string  `json:"description" binding:"required"`
		Cost            float64 `json:"cost" binding:"required"`
		MaintenanceDate string  `json:"maintenanceDate" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid input"))
		return
	}

	maintenanceDate, err := time.Parse("2006-01-02", input.MaintenanceDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid date format, use YYYY-MM-DD"))
		return
	}

	requestID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid request ID"))
		return
	}

	err = ctrl.maintenanceRequestService.ApproveMaintenanceRequest(
		requestID, input.Worker, input.Description, input.Cost, maintenanceDate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Maintenance request approved and maintenance created", nil))
}

// **Reject a maintenance request**
func (ctrl *MaintenanceRequestController) RejectMaintenanceRequest(c *gin.Context) {
	requestID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid request ID"))
		return
	}

	err = ctrl.maintenanceRequestService.RejectMaintenanceRequest(requestID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Maintenance request rejected successfully", nil))
}

// **Delete a maintenance request**
func (ctrl *MaintenanceRequestController) DeleteMaintenanceRequest(c *gin.Context) {
	requestID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid request ID"))
		return
	}

	err = ctrl.maintenanceRequestService.DeleteMaintenanceRequest(requestID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Maintenance request deleted successfully", nil))
}
