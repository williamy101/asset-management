package controller

import (
	"go-asset-management/service"
	"go-asset-management/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AssetCategoryController struct {
	service service.AssetCategoryService
}

func NewAssetCategoryController(service service.AssetCategoryService) *AssetCategoryController {
	return &AssetCategoryController{service: service}
}

func (c *AssetCategoryController) Create(ctx *gin.Context) {
	var input struct {
		CategoryName string `json:"categoryName" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid input"))
		return
	}

	err := c.service.Create(input.CategoryName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to create category"))
		return
	}

	ctx.JSON(http.StatusOK, util.NewSuccessResponse("Category created successfully", nil))
}

func (c *AssetCategoryController) GetAll(ctx *gin.Context) {
	categories, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch categories"))
		return
	}
	ctx.JSON(http.StatusOK, util.NewSuccessResponse("Categories fetched successfully", categories))
}

func (c *AssetCategoryController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid category ID"))
		return
	}

	category, err := c.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, util.NewFailedResponse("Category not found"))
		return
	}

	ctx.JSON(http.StatusOK, util.NewSuccessResponse("Category fetched successfully", category))
}

func (c *AssetCategoryController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid category ID"))
		return
	}

	var input struct {
		CategoryName string `json:"categoryName" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid input"))
		return
	}

	err = c.service.Update(id, input.CategoryName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to update category"))
		return
	}

	ctx.JSON(http.StatusOK, util.NewSuccessResponse("Category updated successfully", nil))
}

func (c *AssetCategoryController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid category ID"))
		return
	}

	err = c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to delete category"))
		return
	}

	ctx.JSON(http.StatusOK, util.NewSuccessResponse("Category deleted successfully", nil))
}
