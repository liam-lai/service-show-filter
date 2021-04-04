package router

import (
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	c "github.com/liam-lai/service-show-filter/app/controller"
	"github.com/liam-lai/service-show-filter/app/middleware"
)

func Router() *gin.Engine {
	r := gin.New()

	r.Use(logger.SetLogger(logger.Config{
		Logger: middleware.Logger(),
		UTC:    true,
	}))

	r.GET("/status", c.Status)
	r.POST("/", c.DrmEnabled)
	return r
}
