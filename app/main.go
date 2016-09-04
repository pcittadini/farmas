package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/pcittadini/farmas/app/handlers"
	"os"
	"time"
)

func main() {

	router := gin.Default()

	// setup cors
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE, PATCH, OPTIONS",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          500 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	api := router.Group("/")
	{
		api.GET("/", app.Root)
		api.POST("/:key/:value", app.SetKey)
		api.GET("/:key", app.GetKey)
	}

	hostname, _ := os.Hostname()

	fmt.Println("API started on : " + hostname)
	router.Run("0.0.0.0:8080")
}
