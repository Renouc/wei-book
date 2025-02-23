package web

import "github.com/gin-gonic/gin"

func RegisterRoutes() *gin.Engine {
	server := gin.Default()

	uServer := NewUserHandler()
	uServer.RegisterUserServer(server)

	return server
}
