package controllers

import (
	"github.com/gin-gonic/gin"
)
func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello!",
	})
}
