package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucassantos-dev/go-api/handler"
)

func initializerRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeninHandler)

		v1.POST("/opening", handler.CreteOpeninHandler)

		v1.DELETE("/opening", handler.DeleteOpeninHandler)

		v1.PUT("/opening", handler.UpdateOpeninHandler)

		v1.GET("/openings", handler.ListOpeninHandler)

	}
}
