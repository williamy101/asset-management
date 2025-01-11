package config

import (
	"go-asset-management/controller"
	"go-asset-management/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, roleController *controller.RoleController, userController *controller.UserController, assetController *controller.AssetController, assetCategoryController *controller.AssetCategoryController, statusController *controller.StatusController) {

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userController.Register)
		userRouter.POST("/login", userController.Login)
		userRouter.GET("/:id", userController.GetUserByID)
	}

	roleRouter := router.Group("/roles")
	roleRouter.Use(middleware.AuthMiddleware(1))
	{
		roleRouter.POST("/", roleController.Create)
		roleRouter.GET("/", roleController.GetAll)
		roleRouter.GET("/:id", roleController.GetByID)
		roleRouter.DELETE("/:id", roleController.Delete)
	}

	statusRouter := router.Group("/statuses")
	statusRouter.Use(middleware.AuthMiddleware(1))
	{
		statusRouter.POST("/", statusController.Create)
		statusRouter.GET("/", statusController.GetAll)
		statusRouter.GET("/:id", statusController.GetByID)
		statusRouter.PUT("/:id", statusController.Update)
		statusRouter.DELETE("/:id", statusController.Delete)
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
}
