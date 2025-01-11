package controller

import (
	"go-asset-management/service"
	"go-asset-management/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

// Register user
func (ctrl *UserController) Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
		RoleID   int    `json:"roleId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse(err.Error()))
		return
	}

	err := ctrl.userService.RegisterUser(req.Name, req.Email, req.Password, req.RoleID)
	if err != nil {
		c.JSON(http.StatusConflict, util.NewFailedResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, util.NewSuccessResponse("User registered successfully", nil))
}

// Login user
func (ctrl *UserController) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse(err.Error()))
		return
	}

	token, err := ctrl.userService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, util.NewFailedResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("Login successful", token))
}

func (ctrl *UserController) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.NewFailedResponse("Invalid user ID"))
		return
	}

	user, err := ctrl.userService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, util.NewFailedResponse("User not found"))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("User fetched successfully", user))
}

func (ctrl *UserController) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")

	user, err := ctrl.userService.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, util.NewFailedResponse("User not found"))
		return
	}

	c.JSON(http.StatusOK, util.NewSuccessResponse("User fetched successfully", user))
}
