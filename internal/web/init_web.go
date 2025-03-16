package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func RegisterRoutes() *gin.Engine {
	server := gin.Default()

	// 处理跨域
	server.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return strings.HasPrefix(origin, "http://localhost:3000")
		},
		MaxAge: 12 * time.Hour,
	}))

	uServer := NewUserHandler()
	uServer.RegisterUserServer(server)

	return server
}
