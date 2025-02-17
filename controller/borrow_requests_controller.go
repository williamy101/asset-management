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

type BorrowAssetRequestController struct {
	borrowAssetRequestService service.BorrowAssetRequestService
}

func NewBorrowAssetRequestController(borrowAssetRequestService service.BorrowAssetRequestService) *BorrowAssetRequestController {
	return &BorrowAssetRequestController{borrowAssetRequestService: borrowAssetRequestService}
}

func (ctrl *BorrowAssetRequestController) CreateBorrowRequest(c *gin.Context) {
	var input struct {
		AssetID            int    `json:"assetId" binding:"required"`
		RequestedStartDate string `json:"requestedStartDate" binding:"required"`
		RequestedEndDate   string `json:"requestedEndDate" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid input"))
		return
	}

	startDate, err := time.Parse("2006-01-02", input.RequestedStartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid start date format, use YYYY-MM-DD"))
		return
	}

	endDate, err := time.Parse("2006-01-02", input.RequestedEndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid end date format, use YYYY-MM-DD"))
		return
	}

	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, util.NewFailedResponse("Unauthorized"))
		return
	}

	err = ctrl.borrowAssetRequestService.CreateBorrowRequest(input.AssetID, userID.(int), startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Borrow request created successfully", nil))
}

func (ctrl *BorrowAssetRequestController) GetAllBorrowRequests(c *gin.Context) {
	requests, err := ctrl.borrowAssetRequestService.GetAllBorrowRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch borrow requests"))
		return
	}
	c.JSON(http.StatusOK, util.NewSuccessResponse("Borrow requests fetched successfully", requests))
}

func (ctrl *BorrowAssetRequestController) GetBorrowRequestByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid request ID"))
		return
	}

	request, err := ctrl.borrowAssetRequestService.GetBorrowRequestByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, util.NewFailedResponse("Borrow request not found"))
		} else {
			c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch borrow request"))
		}
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Borrow request fetched successfully", request))
}

func (ctrl *BorrowAssetRequestController) GetBorrowRequestsByUserID(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, util.NewFailedResponse("Unauthorized"))
		return
	}

	requests, err := ctrl.borrowAssetRequestService.GetBorrowRequestsByUserID(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch borrow requests"))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Borrow requests fetched successfully", requests))
}

func (ctrl *BorrowAssetRequestController) GetBorrowRequestsByAssetID(c *gin.Context) {
	assetID, err := strconv.Atoi(c.Param("assetId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid asset ID"))
		return
	}

	requests, err := ctrl.borrowAssetRequestService.GetBorrowRequestsByAssetID(assetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch borrow requests"))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Borrow requests fetched successfully", requests))
}

func (ctrl *BorrowAssetRequestController) GetBorrowRequestsByStatus(c *gin.Context) {
	statusID, err := strconv.Atoi(c.Param("statusId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid status ID"))
		return
	}

	requests, err := ctrl.borrowAssetRequestService.GetBorrowRequestsByStatus(statusID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch borrow requests"))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Borrow requests fetched successfully", requests))
}

func (ctrl *BorrowAssetRequestController) ApproveBorrowRequest(c *gin.Context) {
	adminID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, util.NewFailedResponse("Unauthorized"))
		return
	}

	requestID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid request ID"))
		return
	}

	err = ctrl.borrowAssetRequestService.ApproveBorrowRequest(requestID, adminID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Borrow request approved successfully", nil))
}

func (ctrl *BorrowAssetRequestController) RejectBorrowRequest(c *gin.Context) {
	requestID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid request ID"))
		return
	}

	err = ctrl.borrowAssetRequestService.RejectBorrowRequest(requestID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Borrow request rejected successfully", nil))
}

func (ctrl *BorrowAssetRequestController) DeleteBorrowRequest(c *gin.Context) {
	requestID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid request ID"))
		return
	}

	err = ctrl.borrowAssetRequestService.DeleteBorrowRequest(requestID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Borrow request deleted successfully", nil))
}
