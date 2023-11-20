package router

import (
	"v1/internal/app"
	"v1/internal/controller"
)

func BootRoutes() {
	app.RegisterRoute("/", controller.HandleTest)
	app.RegisterRoute("/sec", controller.SecondHandleTest)
}
