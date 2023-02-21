package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/igroks/sd-project/backend/app/controllers"
)

func HandleResquest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()

	routes.Use(cors.Default())

	routes.GET("/metrics")
	routes.GET("/health")
	routes.POST("/:databaseName", controllers.Add)
	routes.GET("/:databaseName", controllers.List)
	routes.DELETE("/:databaseName", controllers.Delete)

	return routes
}
