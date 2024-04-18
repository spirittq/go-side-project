package router

import (
	"sideq/internal/usecase"
	"sideq/pkg/logger"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, l logger.Interface, t usecase.Example) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	h := handler.Group("v1")
	{
		newExampleRoutes(h, t, l)
	}
}
