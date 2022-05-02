package router

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	router := gin.Default()
	LoginRouter(router)
	return router
}
