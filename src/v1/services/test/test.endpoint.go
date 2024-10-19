package test

import (
	"net/http"

	devfleError "github.com/GooDu-dev/gd-practical-project-backend/utils/error"

	"github.com/gin-gonic/gin"
)

type Endpoint struct {
	service *TestService
}

func NewEndPoint() *Endpoint {
	service := TestService{}
	return &Endpoint{
		service: service.NewService(),
	}
}

func (e *Endpoint) GetTestHealth(c *gin.Context) {
	var req PingRequest
	if err := c.BindJSON(&req); err != nil {
		status, res := devfleError.GetErrorResponse(err)
		c.JSON(status, res)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong!",
	})
}
