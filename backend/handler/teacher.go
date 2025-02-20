package handler

import (
	"net/http"
	"student-api/model"
	service "student-api/service"

	"github.com/gin-gonic/gin"
)

func TeacherRegister(c *gin.Context) {

	// Validate input
	var teacher model.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	teacherService := service.TeacherRepository{}
	response, err := teacherService.TeacherRegister(&teacher)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func TeacherLogIn(c *gin.Context) {

	// Validate input
	var teacher model.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	teacherService := service.TeacherRepository{}
	response, err := teacherService.TeacherLogIn(&teacher)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func GetTeachers(c *gin.Context) {

	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	teacherService := service.TeacherRepository{}
	response, err := teacherService.GetTeachers(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func GetTeacherDetail(c *gin.Context) {

	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	teacherService := service.TeacherRepository{}
	response, err := teacherService.GetTeacherDetail(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func DeleteTeachers(c *gin.Context) {
}

func DeleteTeachersByClass(c *gin.Context) {
}

func DeleteTeacher(c *gin.Context) {
}

func UpdateTeacherSubject(c *gin.Context) {
}

func TeacherAttendance(c *gin.Context) {
}
