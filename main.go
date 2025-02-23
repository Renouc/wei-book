package main

import (
	"weibook/internal/web"
)

func main() {
	server := web.RegisterRoutes()

	server.Run(":8080")
}
