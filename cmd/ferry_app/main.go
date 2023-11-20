package main

import (
	"v1/internal/app"
	"v1/internal/router"
)

func main() {
	router.BootRoutes()
	app.Init()
}
