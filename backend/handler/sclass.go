package handler

import (
	"net/http"
	"student/model"
	service "student/service"

	"github.com/gin-gonic/gin"
)

func SclassCreate(c *gin.Context) {

	// Validate input
	var sclass model.SclassRequest
	if err := c.ShouldBindJSON(&sclass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sClassService := service.SClassRepository{}
	response, err := sClassService.SclassCreate(&sclass)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func SclassList(c *gin.Context) {

	//id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	sclassService := service.SClassRepository{}
	response, err := sclassService.SclassList(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)

}

func GetSclassDetail(c *gin.Context) {
	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	sclassService := service.SClassRepository{}
	response, err := sclassService.GetSclassDetail(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func GetSclassStudents(c *gin.Context) {

	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	sclassService := service.SClassRepository{}
	response, err := sclassService.GetSclassStudents(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func DeleteSclasses(c *gin.Context) {
	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	sclassService := service.SClassRepository{}
	response, err := sclassService.DeleteSclasses(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func DeleteSclass(c *gin.Context) {
	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	sclassService := service.SClassRepository{}
	response, err := sclassService.DeleteSclass(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)

}
