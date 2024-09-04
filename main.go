package main

import (
	"log"
	"net/http"
	"os"

	"github.com/0xatanda/shantytown/handlers"
	"github.com/0xatanda/shantytown/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port isn't found in the env")
	}

	router := gin.New()

	router.Use(middleware.CORS())
	router.Static("/static", "./static")

	router.GET("/", handlers.Home)

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Server is starting on port: ", portString)
	}
}
