package config

import (
	"go-asset-management/controller"
	"go-asset-management/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, roleController *controller.RoleController, userController *controller.UserController, assetController *controller.AssetController, assetCategoryController *controller.AssetCategoryController, statusController *controller.StatusController, maintenanceController *controller.MaintenanceController) {

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userController.Register)
		userRouter.POST("/login", userController.Login)
	}

	adminUserRouter := router.Group("/users/get")
	adminUserRouter.Use(middleware.AuthMiddleware(1)) // get user cuma bisa diakses admin
	{
		adminUserRouter.GET("/", userController.GetAllUsers)
		adminUserRouter.GET("/:id", userController.GetUserByID)
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

	userStatusRouter := router.Group("/statuses/user", middleware.AuthMiddleware(2))
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
	}

	userAssetRouter := router.Group("/assets/get")
	userAssetRouter.Use(middleware.AuthMiddleware(2)) // getter aset untuk user
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

	userMaintenanceRouter := router.Group("/maintenances/user", middleware.AuthMiddleware(2))
	{
		userMaintenanceRouter.GET("/", maintenanceController.GetMaintenancesByUserID)
	}

	commonMaintenanceRouter := router.Group("/maintenances/update", middleware.AuthMiddleware(1, 2))
	{
		commonMaintenanceRouter.PUT("/:id", maintenanceController.UpdateMaintenance)
	}
}
