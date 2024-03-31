package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	// logger.Infof("Request received: %v", request)
	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("Error creating opening: %v", err)
		// ctx.JSON(http.StatusInternalServerError, gin.H{
		// 	"error": "Error creating opening",
		// })
		return
	}
}
