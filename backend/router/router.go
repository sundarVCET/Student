package router

import (
	"net/http"
	handler "student-api/handler"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Cors : handle client origin rules
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

// SetupRoutes : create endpoints
func SetupRoutes(router *gin.Engine) {

	//Default : debug
	gin.SetMode(viper.GetString("GinMode"))

	//r.Use(static.Serve("/", static.LocalFile("./view", true))) //Serve frontend static files, ex: ReactJS

	// router.Use(gin.LoggerWithWriter(middleware.LogWriter()))
	// router.Use(gin.CustomRecovery(middleware.AppRecovery()))

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})
	// Define the paths to exclude from authentication
	//excludedPaths := []string{"/StudentReg", "/StudentLogin", "/AdminReg", "/AdminLogin", "/TeacherReg", "/TeacherLogin"}

	// Apply the custom middleware globally
	//router.Use(middleware.AuthWithExceptions(excludedPaths))
	// Admin

	router.POST("/AdminReg", handler.AdminRegister)
	router.POST("/AdminLogin", handler.AdminLogIn)

	router.GET("/Admin/:id", handler.GetAdminDetail)

	// router.delete("/Admin/:id", deleteAdmin)

	// router.put("/Admin/:id", updateAdmin)

	// Student

	router.POST("/StudentReg", handler.StudentRegister)
	router.POST("/StudentLogin", handler.StudentLogIn)

	router.GET("/Students/:id", handler.GetStudents)
	router.GET("/Student/:id", handler.GetStudentDetail)

	router.DELETE("/Students/:id", handler.DeleteStudents)
	router.DELETE("/StudentsClass/:id", handler.DeleteStudentsByClass)
	router.DELETE("/Student/:id", handler.DeleteStudent)

	router.PUT("/Student/:id", handler.UpdateStudent)

	router.PUT("/UpdateExamResult/:id", handler.UpdateExamResult)

	router.PUT("/StudentAttendance/:id", handler.StudentAttendance)

	router.PUT("/RemoveAllStudentsSubAtten/:id", handler.ClearAllStudentsAttendanceBySubject)
	router.PUT("/RemoveAllStudentsAtten/:id", handler.ClearAllStudentsAttendance)

	router.PUT("/RemoveStudentSubAtten/:id", handler.RemoveStudentAttendanceBySubject)
	router.PUT("/RemoveStudentAtten/:id", handler.RemoveStudentAttendance)

	// Sclass

	router.POST("/SclassCreate", handler.SclassCreate)

	router.GET("/SclassList/:id", handler.SclassList)
	router.GET("/Sclass/:id", handler.GetSclassDetail)
	router.GET("/Sclass/Students/:id", handler.GetSclassStudents)

	router.DELETE("/Sclasses/:id", handler.DeleteSclasses)
	router.DELETE("/Sclass/:id", handler.DeleteSclass)

	// Subject

	router.POST("/SubjectCreate", handler.SubjectCreate)

	router.GET("/AllSubjects/:id", handler.AllSubjects)
	router.GET("/ClassSubjects/:id", handler.ClassSubjects)
	router.GET("/FreeSubjectList/:id", handler.FreeSubjectList)
	router.GET("/Subject/:id", handler.GetSubjectDetail)

	router.DELETE("/Subject/:id", handler.DeleteSubject)
	router.DELETE("/Subjects/:id", handler.DeleteSubjects)
	router.DELETE("/SubjectsClass/:id", handler.DeleteSubjectsByClass)

	// Teacher

	router.POST("/TeacherReg", handler.TeacherRegister)
	router.POST("/TeacherLogin", handler.TeacherLogIn)

	router.GET("/Teachers/:id", handler.GetTeachers)
	router.GET("/Teacher/:id", handler.GetTeacherDetail)

	router.DELETE("/Teachers/:id", handler.DeleteTeachers)
	router.DELETE("/TeachersClass/:id", handler.DeleteTeachersByClass)
	router.DELETE("/Teacher/:id", handler.DeleteTeacher)

	router.PUT("/TeacherSubject", handler.UpdateTeacherSubject)
	router.POST("/TeacherAttendance/:id", handler.TeacherAttendance)

	// Notice

	router.POST("/NoticeCreate", handler.NoticeCreate)

	router.GET("/NoticeList/:id", handler.NoticeList)

	router.DELETE("/Notices/:id", handler.DeleteNotices)
	router.DELETE("/Notice/:id", handler.DeleteNotice)

	router.PUT("/Notice/:id", handler.UpdateNotice)

	// Complain

	router.POST("/ComplainCreate", handler.ComplainCreate)

	router.GET("/ComplainList/:id", handler.ComplainList)

	// Image

	router.POST("/AddImage", handler.AddImage)
	router.GET("/GetImage/:userId", handler.GetImage)

	// secured := router.Group("/secured").Use(middleware.Auth())
	// {
	// 	secured.GET("/ping", handler.Ping)
	// }
}

// func AuthRoutes(incomingRoutes *gin.Engine) {

// 	incomingRoutes.Use(middleware.UserAuthenticate())
// }
