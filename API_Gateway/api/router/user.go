package router

import (
	"api_gateway/api/handler"
	"api_gateway/grpc_client/user"
	"api_gateway/logs"
	"net/http"
)

func UserRouter() {
	lgs := logs.GetLogger("logs/logger.log")
	uClient := user.UserApiConn("userservice:8081")

	userHabdler := handler.NewUserClient(uClient, lgs)

	// http.Handle("/", mid.JWTMiddleware(http.HandlerFunc(userHabdler.CreateUser)))
	// http.Handle("/", mid.JWTMiddleware(http.HandlerFunc(userHabdler.LoginUser)))
	// http.Handle("/", mid.JWTMiddleware(http.HandlerFunc(userHabdler.GetUserProfile)))

	http.HandleFunc("POST /user/register", userHabdler.CreateUser)
	http.HandleFunc("POST /user/login", userHabdler.LoginUser)
	http.HandleFunc("GET /user/profile", userHabdler.GetUserProfile)
}
