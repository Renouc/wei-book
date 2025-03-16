package main

import (
	"fmt"
	"weibook/internal/web"
)

const PORT = "8080"

func main() {
	server := web.RegisterRoutes()

	server.Run(fmt.Sprint(":", PORT))
}
