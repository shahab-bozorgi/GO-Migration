package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	
	router.GET("/test", GetTablesHandler) 
	router.GET("/tables", GetTables) 
	router.POST("/query", ApplyQuery)
}
