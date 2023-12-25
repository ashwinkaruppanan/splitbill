package router

import (
	"net/http"
	"splitbill/api/handler"
	"splitbill/api/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, handler *handler.UserHandler) {

	r.Use(middleware.CORSMiddleware())

	publicHandler := r.Group("/api/v1/")

	publicHandler.GET("/check", func(c *gin.Context) {
		c.JSON(http.StatusOK, "working")
	})
	publicHandler.POST("/signup", handler.CreateUser)
	publicHandler.POST("/signin", handler.LoginUser)
}
