package controller

import (
	"go-asset-management/service"
	"go-asset-management/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AssetController struct {
	assetService service.AssetService
}

func NewAssetController(assetService service.AssetService) *AssetController {
	return &AssetController{assetService: assetService}
}

func (ctrl *AssetController) CreateAsset(c *gin.Context) {
	var input struct {
		AssetName  string `json:"assetName" binding:"required"`
		CategoryID *int   `json:"categoryID"`
		StatusID   int    `json:"statusId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid input"))
		return
	}

	err := ctrl.assetService.CreateAsset(input.AssetName, input.CategoryID, input.StatusID, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Asset created successfully", nil))
}

func (ctrl *AssetController) GetAllAssets(c *gin.Context) {
	assets, err := ctrl.assetService.GetAllAssets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch assets"))
		return
	}
	c.JSON(http.StatusOK, util.NewSuccessResponse("Assets fetched successfully", assets))
}

func (ctrl *AssetController) GetAssetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid asset ID"))
		return
	}

	asset, err := ctrl.assetService.GetAssetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, util.NewFailedResponse("Asset not found"))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Asset fetched successfully", asset))
}

func (ctrl *AssetController) UpdateAsset(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid asset ID"))
		return
	}

	var input struct {
		AssetName  *string `json:"assetName"`
		CategoryID *int    `json:"categoryID"`
		StatusID   *int    `json:"statusId"`
		UserID     *int    `json:"userId"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid input"))
		return
	}

	err = ctrl.assetService.UpdateAsset(id, input.AssetName, input.CategoryID, input.StatusID, input.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Asset updated successfully", nil))
}

func (ctrl *AssetController) DeleteAsset(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid asset ID"))
		return
	}

	err = ctrl.assetService.DeleteAsset(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.NewFailedResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Asset deleted successfully", nil))
}
