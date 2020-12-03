package ControllerSys

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context)   {
	c.JSON(http.StatusOK, gin.H{
		"status": -1,
		"msg":  "",
		"data":   nil,
	})
}


