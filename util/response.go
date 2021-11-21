package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GoodResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    data,
	})
}

func BadResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, Response{
		Status:  http.StatusBadRequest,
		Message: err.Error(),
		Data:    nil,
	})
}
