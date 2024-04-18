package router

import (
	"net/http"
	"sideq/internal/entity"
	"sideq/internal/usecase"
	"sideq/pkg/logger"

	"github.com/gin-gonic/gin"
)

type exampleRoutes struct {
	t usecase.Example
	l logger.Interface
}

type ErrorDetails struct {
	Details string
}

func newExampleRoutes(handler *gin.RouterGroup, t usecase.Example, l logger.Interface) {
	r := exampleRoutes{t, l}

	h := handler.Group("/example")
	{
		h.GET("/", r.GetExamples)
		h.POST("/", r.PostExample)
	}
}

func (r *exampleRoutes) GetExamples(c *gin.Context) {
	examples, err := r.t.GetExamples(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorDetails{Details: err.Error()})
		return
	}

	c.JSON(http.StatusOK, examples)
}

func (r *exampleRoutes) PostExample(c *gin.Context) {

	examplePostRequest := entity.ExamplePostRequest{}
	err := c.ShouldBindJSON(&examplePostRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorDetails{Details: err.Error()})
		return
	}

	example, err := r.t.PostExample(c, examplePostRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorDetails{Details: err.Error()})
		return
	}

	c.JSON(http.StatusOK, example)
}