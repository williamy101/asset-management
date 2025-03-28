package config

import (
	"go-asset-management/controller"
	"go-asset-management/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	router *gin.Engine,
	roleController *controller.RoleController,
	userController *controller.UserController,
	assetController *controller.AssetController,
	assetCategoryController *controller.AssetCategoryController,
	statusController *controller.StatusController,
	maintenanceController *controller.MaintenanceController,
	maintenanceRequestController *controller.MaintenanceRequestController,
	borrowRequestController *controller.BorrowAssetRequestController,
	borrowedAssetController *controller.BorrowedAssetController,
) {

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userController.Register)
		userRouter.POST("/login", userController.Login)
	}

	adminUserRouter := router.Group("/users/admin")
	adminUserRouter.Use(middleware.AuthMiddleware(1))
	{
		adminUserRouter.GET("/", userController.GetAllUsers)
		adminUserRouter.GET("/:id", userController.GetUserByID)
		adminUserRouter.PUT("/role", userController.UpdateUserRole)
		adminUserRouter.GET("/filter", userController.FilterUsers)
	}

	roleRouter := router.Group("/roles")
	roleRouter.Use(middleware.AuthMiddleware(1))
	{
		roleRouter.POST("/", roleController.Create)
		roleRouter.GET("/", roleController.GetAll)
		roleRouter.GET("/:id", roleController.GetByID)
		roleRouter.DELETE("/:id", roleController.Delete)
	}

	statusRouter := router.Group("/statuses", middleware.AuthMiddleware(1))
	{
		statusRouter.POST("/", statusController.Create)
		statusRouter.GET("/", statusController.GetAll)
		statusRouter.GET("/:id", statusController.GetByID)
		statusRouter.PUT("/:id", statusController.Update)
		statusRouter.DELETE("/:id", statusController.Delete)
	}

	userStatusRouter := router.Group("/statuses/user", middleware.AuthMiddleware(2, 3))
	{
		userStatusRouter.GET("/", statusController.GetAll)
		userStatusRouter.GET("/:id", statusController.GetByID)
	}

	assetCategoryRouter := router.Group("/categories")
	assetCategoryRouter.Use(middleware.AuthMiddleware(1))
	{
		assetCategoryRouter.POST("/", assetCategoryController.Create)
		assetCategoryRouter.GET("/", assetCategoryController.GetAll)
		assetCategoryRouter.GET("/:id", assetCategoryController.GetByID)
		assetCategoryRouter.PUT("/:id", assetCategoryController.Update)
		assetCategoryRouter.DELETE("/:id", assetCategoryController.Delete)
	}

	assetRouter := router.Group("/assets")
	assetRouter.Use(middleware.AuthMiddleware(1))
	{
		assetRouter.POST("/", assetController.CreateAsset)
		assetRouter.GET("/", assetController.GetAllAssets)
		assetRouter.GET("/:id", assetController.GetAssetByID)
		assetRouter.PUT("/:id", assetController.UpdateAsset)
		assetRouter.DELETE("/:id", assetController.DeleteAsset)
		assetRouter.GET("/filter", assetController.FilterAssets)
	}

	userAssetRouter := router.Group("/assets/get")
	userAssetRouter.Use(middleware.AuthMiddleware(2, 3))
	{
		userAssetRouter.GET("/", assetController.GetAllAssets)
		userAssetRouter.GET("/:id", assetController.GetAssetByID)
	}

	adminMaintenanceRouter := router.Group("/maintenances", middleware.AuthMiddleware(1))
	{
		adminMaintenanceRouter.POST("/", maintenanceController.CreateMaintenance)
		adminMaintenanceRouter.GET("/", maintenanceController.GetAllMaintenances)
		adminMaintenanceRouter.GET("/:id", maintenanceController.GetMaintenanceByID)
		adminMaintenanceRouter.DELETE("/:id", maintenanceController.DeleteMaintenance)
		adminMaintenanceRouter.GET("/total-cost", maintenanceController.GetTotalCost)
		adminMaintenanceRouter.GET("/total-cost/:asset_id", maintenanceController.GetTotalCostByAssetID)
	}

	technicianMaintenanceRouter := router.Group("/maintenances/technician", middleware.AuthMiddleware(2))
	{
		technicianMaintenanceRouter.PUT("/:id/start", maintenanceController.StartMaintenance)
		technicianMaintenanceRouter.PUT("/:id/end", maintenanceController.EndMaintenance)
	}

	userMaintenanceRouter := router.Group("/maintenances/user", middleware.AuthMiddleware(2, 3))
	{
		userMaintenanceRouter.GET("/", maintenanceController.GetMaintenancesByWorkerID)
	}

	maintenanceRequestRouter := router.Group("/maintenance-requests")
	{
		maintenanceRequestRouter.Use(middleware.AuthMiddleware(3))
		maintenanceRequestRouter.POST("/", maintenanceRequestController.CreateMaintenanceRequest)
	}

	adminMaintenanceRequestRouter := router.Group("/maintenance-requests/admin")
	adminMaintenanceRequestRouter.Use(middleware.AuthMiddleware(1))
	{
		adminMaintenanceRequestRouter.GET("/", maintenanceRequestController.GetAllMaintenanceRequests)
		adminMaintenanceRequestRouter.PUT("/:id/approve", maintenanceRequestController.ApproveMaintenanceRequest)
		adminMaintenanceRequestRouter.PUT("/:id/reject", maintenanceRequestController.RejectMaintenanceRequest)
	}

	borrowRequestRouter := router.Group("/borrow-requests")
	borrowRequestRouter.Use(middleware.AuthMiddleware(3))
	{
		borrowRequestRouter.POST("/", borrowRequestController.CreateBorrowRequest)
		borrowRequestRouter.GET("/", borrowRequestController.GetBorrowRequestsByUserID)
	}

	adminBorrowRequestRouter := router.Group("/borrow-requests/admin")
	adminBorrowRequestRouter.Use(middleware.AuthMiddleware(1))
	{
		adminBorrowRequestRouter.GET("/", borrowRequestController.GetAllBorrowRequests)
		adminBorrowRequestRouter.GET("/:id", borrowRequestController.GetBorrowRequestByID)
		adminBorrowRequestRouter.PUT("/:id/approve", borrowRequestController.ApproveBorrowRequest)
		adminBorrowRequestRouter.PUT("/:id/reject", borrowRequestController.RejectBorrowRequest)
	}

	borrowedAssetRouter := router.Group("/borrowed-assets")
	borrowedAssetRouter.Use(middleware.AuthMiddleware(1, 2, 3))
	{
		borrowedAssetRouter.GET("/", borrowedAssetController.GetAllBorrowedAssets)
		borrowedAssetRouter.GET("/:id", borrowedAssetController.GetBorrowedAssetByID)
		borrowedAssetRouter.PUT("/:id/return", borrowedAssetController.UpdateReturnDate)
	}
}
