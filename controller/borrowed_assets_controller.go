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

type BorrowedAssetController struct {
	borrowedAssetService service.BorrowedAssetService
}

func NewBorrowedAssetController(borrowedAssetService service.BorrowedAssetService) *BorrowedAssetController {
	return &BorrowedAssetController{borrowedAssetService: borrowedAssetService}
}

func (ctrl *BorrowedAssetController) GetAllBorrowedAssets(c *gin.Context) {
	borrowedAssets, err := ctrl.borrowedAssetService.GetAllBorrowedAssets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch borrowed assets"))
		return
	}
	c.JSON(http.StatusOK, util.NewSuccessResponse("Borrowed assets fetched successfully", borrowedAssets))
}

func (ctrl *BorrowedAssetController) GetBorrowedAssetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid borrowed asset ID"))
		return
	}

	borrowedAsset, err := ctrl.borrowedAssetService.GetBorrowedAssetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, util.NewFailedResponse("Borrowed asset not found"))
		} else {
			c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch borrowed asset"))
		}
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Borrowed asset fetched successfully", borrowedAsset))
}

func (ctrl *BorrowedAssetController) GetBorrowedAssetsByUserID(c *gin.Context) {
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, util.NewFailedResponse("Unauthorized"))
		return
	}

	borrowedAssets, err := ctrl.borrowedAssetService.GetBorrowedAssetsByUserID(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch borrowed assets"))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Borrowed assets fetched successfully", borrowedAssets))
}

func (ctrl *BorrowedAssetController) GetBorrowedAssetsByAssetID(c *gin.Context) {
	assetID, err := strconv.Atoi(c.Param("assetId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid asset ID"))
		return
	}

	borrowedAssets, err := ctrl.borrowedAssetService.GetBorrowedAssetsByAssetID(assetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch borrowed assets"))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Borrowed assets fetched successfully", borrowedAssets))
}

func (ctrl *BorrowedAssetController) GetBorrowedAssetsByStatus(c *gin.Context) {
	statusID, err := strconv.Atoi(c.Param("statusId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid status ID"))
		return
	}

	borrowedAssets, err := ctrl.borrowedAssetService.GetBorrowedAssetsByStatus(statusID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch borrowed assets"))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Borrowed assets fetched successfully", borrowedAssets))
}

func (ctrl *BorrowedAssetController) UpdateReturnDate(c *gin.Context) {
	var input struct {
		ReturnDate string `json:"returnDate" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid input"))
		return
	}

	returnDate, err := time.Parse("2006-01-02", input.ReturnDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid date format, use YYYY-MM-DD"))
		return
	}

	borrowID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid borrowed asset ID"))
		return
	}

	err = ctrl.borrowedAssetService.UpdateReturnDate(borrowID, returnDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Return date updated successfully", nil))
}
