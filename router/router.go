package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {

	//Initialize the router
	router := gin.Default()

	//Initialize routes
	InitializeRoutes(router)

	//Start the server
	router.Run(":8080")
}
