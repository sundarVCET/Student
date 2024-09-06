package handler

import (
	"fmt"
	"net/http"
	"student/model"
	service "student/service"
	validate "student/validate"

	"github.com/gin-gonic/gin"
)

// AdminRegister handles the registration of a new admin.
// @Summary Register a new admin
// @Description Register a new admin with the provided details.
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin body model.Admin true "Admin Registration Details"
// @Success 200 {object} model.Admin "Successfully registered admin"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /AdminReg [post]
func AdminRegister(c *gin.Context) {

	// Validate input
	var admin *model.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validate.Validate(admin); err != nil {
		fmt.Println("err in validate", err)
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	adminService := service.AdminRepository{}
	response, err := adminService.AdminRegister(admin)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
func AdminLogIn(c *gin.Context) {

	// Validate input
	var admin model.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	adminService := service.AdminRepository{}
	response, err := adminService.AdminLogIn(&admin)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)

}
func GetAdminDetail(c *gin.Context) {

	//id  Parameter  from url
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"message": "Id not found"})
		return
	}
	adminService := service.AdminRepository{}
	response, err := adminService.GetAdminDetail(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
