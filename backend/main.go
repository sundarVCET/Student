// package main

// import (
// 	"context"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"student-api/config"
// 	db "student-api/database"
// 	"syscall"
// 	"time"

// 	_ "student-api/docs" // This is to ensure docs are generated
// 	router "student-api/router"
// 	validate "student-api/validate"

// 	"github.com/gin-gonic/gin"
// 	"github.com/spf13/viper"
// 	swaggerFiles "github.com/swaggo/files"
// 	ginSwagger "github.com/swaggo/gin-swagger"
// )

// func Init() {
// 	log.Println("Init started")
// }

// // @title SCHOOL_PROJECT
// // @version 1.0
// // @description  school project service
// // @host localhost:8080
// // @BasePath /

// var srv *http.Server
// var sigs chan os.Signal

// func main() {

// 	log.SetFlags(log.LstdFlags | log.Lshortfile)
// 	log.Println("=========SCHOOL PROJECT Starting=====")

// 	config.LoadConfig()
// 	validate.Init()
// 	db.Init()

// 	routerHandler := gin.New()

// 	// Register Swagger handler
// 	routerHandler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

// 	router.SetupRoutes(routerHandler)
// 	//router.AuthRoutes(routerHandler)
// 	routerHandler.Run(viper.GetString("Port"))

// 	srv := &http.Server{
// 		Handler: routerHandler,
// 		Addr:    viper.GetString("Port"),
// 	}
// 	go func() {
// 		panic(srv.ListenAndServe())
// 	}()

// 	// Create channel for shutdown signals.
// 	stop := make(chan os.Signal, 1)
// 	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
// 	go signalHandler()
// 	db.CloseClientDB()
// 	time.Sleep(5 * time.Second)
// 	os.Exit(1)

// }
// func signalHandler() {
// 	sig := <-sigs
// 	log.Println(sig)
// 	if err := srv.Shutdown(context.TODO()); err != nil {
// 		log.Fatalf("Server Shutdown Failed:%+v", err)
// 	}
// 	log.Print("Http Closed")

// }
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"student-api/config"
	db "student-api/database"
	"syscall"
	"time"

	_ "student-api/docs"
	router "student-api/router"
	validate "student-api/validate"

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

	// Load configurations and initialize dependencies
	config.LoadConfig()
	validate.Init()
	db.Init()

	// Setup Gin router
	routerHandler := gin.New()
	routerHandler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.SetupRoutes(routerHandler)

	// Setup HTTP server
	srv := &http.Server{
		Handler: routerHandler,
		Addr:    viper.GetString("Port"),
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server is running on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	// Set up channel to listen for interrupt signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-stop
	log.Println("Shutdown signal received, initiating graceful shutdown...")

	// Create a context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt graceful server shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed: %+v", err)
	}
	log.Println("Server gracefully stopped.")

	// Close database connections
	db.CloseClientDB()
	log.Println("Database connection closed.")

	log.Println("Shutdown complete.")
}
