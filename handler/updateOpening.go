package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOpeninHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Update Opening",
	})
}
