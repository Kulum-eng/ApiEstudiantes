package routes

import (
	"github.com/gin-gonic/gin"

	"API-HEXAGONAL/src/estudiantes/dependencies"
)


func SetupShinobiRoutes(router *gin.Engine, deps *dependencies.ShinobiDependencies) {
	
	shinobiGroup := router.Group("/v1/shinobis")
	{
		shinobiGroup.POST("", deps.CreateShinobiController.Run)
		shinobiGroup.PUT("/:id", deps.UpdateShinobiController.Run)
		shinobiGroup.GET("", deps.GetAllShinobisController)
		shinobiGroup.GET("/:id", deps.GetByIdShinobiController)
		shinobiGroup.DELETE("/:id", deps.DeleteShinobiController.Run)
	}
}
