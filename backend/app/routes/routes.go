package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/igroks/app-task-list/backend/app/controllers"
)

func HandleResquest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()

	routes.Use(cors.Default())

	routes.GET("/metrics")
	routes.GET("/health")
	routes.POST("/task", controllers.Add)
	routes.GET("/task", controllers.List)
	routes.DELETE("/task/:taskId", controllers.Delete)

	return routes
}
