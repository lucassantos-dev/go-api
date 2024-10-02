package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpeninHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "List Opening",
	})
}
