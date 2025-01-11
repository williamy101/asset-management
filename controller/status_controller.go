package controller

import (
	"go-asset-management/service"
	"go-asset-management/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StatusController struct {
	service service.StatusService
}

func NewStatusController(service service.StatusService) *StatusController {
	return &StatusController{service: service}
}

func (c *StatusController) Create(ctx *gin.Context) {
	var input struct {
		StatusName string `json:"statusName" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid input"))
		return
	}

	err := c.service.Create(input.StatusName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to create status"))
		return
	}

	ctx.JSON(http.StatusOK, util.NewSuccessResponse("Status created successfully", nil))
}

func (c *StatusController) GetAll(ctx *gin.Context) {
	statuses, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch statuses"))
		return
	}
	ctx.JSON(http.StatusOK, util.NewSuccessResponse("Status created successfully", statuses))
}

func (c *StatusController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid status ID"))
		return
	}

	status, err := c.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, util.NewFailedResponse("Status not found"))
		return
	}

	ctx.JSON(http.StatusOK, util.NewSuccessResponse("Statuses fetched successfully", status))
}

func (c *StatusController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid status ID"))
		return
	}

	var input struct {
		StatusName string `json:"statusName" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid input"))
		return
	}

	err = c.service.Update(id, input.StatusName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to update status"))
		return
	}

	ctx.JSON(http.StatusOK, util.NewSuccessResponse("Status updated successfully", nil))
}

func (c *StatusController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid status ID"))
		return
	}

	err = c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to delete status"))
		return
	}

	ctx.JSON(http.StatusOK, util.NewSuccessResponse("Status deleted successfully", nil))
}
