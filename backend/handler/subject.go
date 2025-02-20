package handler

import (
	"net/http"
	"student-api/model"
	service "student-api/service"

	"github.com/gin-gonic/gin"
)

func SubjectCreate(c *gin.Context) {

	// Validate input
	var subject model.SubjectRequest
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	subjectService := service.SubjectRepository{}
	response, err := subjectService.SubjectCreate(&subject)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func AllSubjects(c *gin.Context) {

	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	subjectService := service.SubjectRepository{}
	response, err := subjectService.AllSubjects(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func ClassSubjects(c *gin.Context) {
	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	subjectService := service.SubjectRepository{}
	response, err := subjectService.ClassSubjects(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func FreeSubjectList(c *gin.Context) {
	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	subjectService := service.SubjectRepository{}
	response, err := subjectService.FreeSubjectList(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)

}

func GetSubjectDetail(c *gin.Context) {

	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	subjectService := service.SubjectRepository{}
	response, err := subjectService.GetSubjectDetail(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func DeleteSubject(c *gin.Context) {

	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	subjectService := service.SubjectRepository{}
	response, err := subjectService.DeleteSubject(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func DeleteSubjects(c *gin.Context) {

	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	subjectService := service.SubjectRepository{}
	response, err := subjectService.DeleteSubjects(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func DeleteSubjectsByClass(c *gin.Context) {

	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	subjectService := service.SubjectRepository{}
	response, err := subjectService.DeleteSubjectsByClass(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
