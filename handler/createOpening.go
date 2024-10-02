package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreteOpeninHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Create Opening",
	})
}
