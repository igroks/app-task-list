package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/igroks/sd-project/backend/app/controllers"
)

func HandleResquest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()

	routes.GET("/metrics")
	routes.GET("/health")
	routes.POST("/add", controllers.Add)

	return routes
}
