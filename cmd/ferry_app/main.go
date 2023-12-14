package main

import (
	"v1/app"
	"v1/config/handlerprovider"
	"v1/config/serviceprovider"
)

func main() {
	serviceprovider.BindService()
	handlerprovider.BindHandler()
	app.StartServer()
}
