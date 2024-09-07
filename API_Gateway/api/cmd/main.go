package main

import (
	_ "api_gateway/api/docs"
	"api_gateway/api/router"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// New ...
// @title Project: SMART HOME CONTROL SYSTEM
// @description This swagger UI was created for Exam
// @version 1.0

// @host localhost:9000

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @tags user
func main() {
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	router.UserRouter()
	router.DeviceRouter()

	log.Println("Server is running on port: 7777")
	if err := http.ListenAndServe(":7777", nil); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
