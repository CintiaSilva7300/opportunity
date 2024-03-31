package router

import (
	"github.com/CintiaSilva7300/go-opportunity/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	//Initialize the handler
	handler.InitializeHandler()

	v1 := router.Group("/api")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/opening/list", handler.ListOpeningHandler)
	}

}

// http://localhost:8080/api/opening
