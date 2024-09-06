package handler

import (
	"net/http"
	"student/model"
	service "student/service"

	"github.com/gin-gonic/gin"
)

func ComplainCreate(c *gin.Context) {

	var complain model.Complain
	if err := c.ShouldBindJSON(&complain); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	complainService := service.ComplainRepository{}
	response, err := complainService.ComplainCreate(&complain)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func ComplainList(c *gin.Context) {

	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	complainService := service.ComplainRepository{}
	response, err := complainService.ComplainList(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
