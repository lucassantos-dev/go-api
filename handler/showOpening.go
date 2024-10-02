package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeninHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Show Opening",
	})
}
