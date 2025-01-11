package controller

import (
	"go-asset-management/service"
	"go-asset-management/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	service service.RoleService
}

func NewRoleController(service service.RoleService) *RoleController {
	return &RoleController{service: service}
}

func (c *RoleController) Create(ctx *gin.Context) {
	var input struct {
		RoleName string `json:"roleName" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid input"))
		return
	}

	err := c.service.CreateRole(input.RoleName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to create role"))
		return
	}

	ctx.JSON(http.StatusOK, util.NewSuccessResponse("Role created successfully", nil))
}

func (c *RoleController) GetAll(ctx *gin.Context) {
	roles, err := c.service.GetAllRoles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to fetch roles"))
		return
	}
	ctx.JSON(http.StatusOK, util.NewSuccessResponse("Roles fetched successfully", roles))
}

func (c *RoleController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid role ID"))
		return
	}

	role, err := c.service.GetRoleByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, util.NewFailedResponse("Role not found"))
		return
	}

	ctx.JSON(http.StatusOK, util.NewSuccessResponse("Role fetched successfully", role))
}

func (c *RoleController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid role ID"))
		return
	}

	err = c.service.DeleteRole(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.NewFailedResponse("Failed to delete role"))
		return
	}

	ctx.JSON(http.StatusOK, util.NewSuccessResponse("Role deleted successfully", nil))
}
