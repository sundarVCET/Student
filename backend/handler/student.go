package handler

import (
	"fmt"
	"net/http"
	"student/model"
	service "student/service"
	validate "student/validate"

	"github.com/gin-gonic/gin"
)

func StudentRegister(c *gin.Context) {

	// Validate input
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	studentService := service.StudentRepository{}
	response, err := studentService.StudentRegister(&student)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)

}
func StudentLogIn(c *gin.Context) {

	// Validate input
	var student model.StudentLoginRequest
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validate.Validate(student); err != nil {
		fmt.Println("err in validate", err)
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	studentService := service.StudentRepository{}
	response, err := studentService.StudentLogIn(&student)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
func GetStudents(c *gin.Context) {
	/// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	studentService := service.StudentRepository{}
	response, err := studentService.GetStudents(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)

}
func GetStudentDetail(c *gin.Context) {
	/// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	studentService := service.StudentRepository{}
	response, err := studentService.GetStudentDetail(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)

}
func DeleteStudents(c *gin.Context) {
	/// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	studentService := service.StudentRepository{}
	response, err := studentService.DeleteStudents(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)

}
func DeleteStudentsByClass(c *gin.Context) {
	/// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	studentService := service.StudentRepository{}
	response, err := studentService.DeleteStudentsByClass(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)

}
func DeleteStudent(c *gin.Context) {
	/// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	studentService := service.StudentRepository{}
	response, err := studentService.DeleteStudent(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)

}
func UpdateStudent(c *gin.Context) {

	// Validate input
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	studentService := service.StudentRepository{}
	response, err := studentService.UpdateStudent(&student)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func UpdateExamResult(c *gin.Context) {
	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	studentService := service.StudentRepository{}
	response, err := studentService.UpdateExamResult(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func StudentAttendance(c *gin.Context) {
	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	studentService := service.StudentRepository{}
	response, err := studentService.StudentAttendance(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func ClearAllStudentsAttendanceBySubject(c *gin.Context) {
	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	studentService := service.StudentRepository{}
	response, err := studentService.ClearAllStudentsAttendanceBySubject(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func ClearAllStudentsAttendance(c *gin.Context) {
	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	studentService := service.StudentRepository{}
	response, err := studentService.ClearAllStudentsAttendance(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func RemoveStudentAttendanceBySubject(c *gin.Context) {
	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	studentService := service.StudentRepository{}
	response, err := studentService.RemoveStudentAttendanceBySubject(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
func RemoveStudentAttendance(c *gin.Context) {
	// id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	studentService := service.StudentRepository{}
	response, err := studentService.RemoveStudentAttendance(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
