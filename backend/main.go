package main

import (
	"log"
	"student/config"
	db "student/database"

	_ "student/docs" // This is to ensure docs are generated
	router "student/router"
	validate "student/validate"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() {
	log.Println("Init started")
}

// @title SCHOOL_PROJECT
// @version 1.0
// @description  school project service
// @host localhost:8080
// @BasePath /

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("=========SCHOOL PROJECT Starting=====")

	config.LoadConfig()
	validate.Init()
	db.Init()

	routerHandler := gin.New()

	// Register Swagger handler
	routerHandler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.SetupRoutes(routerHandler)
	//router.AuthRoutes(routerHandler)
	routerHandler.Run(viper.GetString("Port"))

}
