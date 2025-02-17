package main

import (
	"go-asset-management/config"
	"go-asset-management/controller"
	"go-asset-management/middleware"
	"go-asset-management/repository"
	"go-asset-management/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.New()

	userRepo := repository.NewUserRepository(config.DB)
	roleRepo := repository.NewRoleRepository(config.DB)
	statusRepo := repository.NewStatusRepository(config.DB)
	assetCategoryRepo := repository.NewAssetCategoryRepository(config.DB)
	assetRepo := repository.NewAssetRepository(config.DB)
	maintenanceRepo := repository.NewMaintenanceRepository(config.DB)
	maintenanceRequestRepo := repository.NewMaintenanceRequestRepository(config.DB)
	borrowRequestRepo := repository.NewBorrowAssetRequestRepository(config.DB)
	borrowedAssetRepo := repository.NewBorrowedAssetRepository(config.DB)

	userService := service.NewUserService(userRepo, roleRepo)
	roleService := service.NewRoleService(roleRepo)
	assetService := service.NewAssetService(assetRepo, assetCategoryRepo, maintenanceRepo)
	assetCategoryService := service.NewAssetCategoryService(assetCategoryRepo)
	statusService := service.NewStatusService(statusRepo)
	maintenanceService := service.NewMaintenanceService(maintenanceRepo, assetRepo, maintenanceRequestRepo)
	maintenanceRequestService := service.NewMaintenanceRequestService(maintenanceRequestRepo, assetRepo, maintenanceRepo, userRepo)
	borrowRequestService := service.NewBorrowAssetRequestService(borrowRequestRepo, borrowedAssetRepo, assetRepo)
	borrowedAssetService := service.NewBorrowedAssetService(borrowedAssetRepo, borrowRequestRepo, assetRepo)

	userController := controller.NewUserController(userService)
	roleController := controller.NewRoleController(roleService)
	statusController := controller.NewStatusController(statusService)
	assetCategoryController := controller.NewAssetCategoryController(assetCategoryService)
	assetController := controller.NewAssetController(assetService)
	maintenanceController := controller.NewMaintenanceController(maintenanceService)
	maintenanceRequestController := controller.NewMaintenanceRequestController(maintenanceRequestService)
	borrowRequestController := controller.NewBorrowAssetRequestController(borrowRequestService)
	borrowedAssetController := controller.NewBorrowedAssetController(borrowedAssetService)

	config.SetupRouter(
		r,
		roleController,
		userController,
		assetController,
		assetCategoryController,
		statusController,
		maintenanceController,
		maintenanceRequestController,
		borrowRequestController,
		borrowedAssetController,
	)

	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.RecoveryMiddleware())

	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
