package main

import (
	"v1/app"
	"v1/config/handlerprovider"
	"v1/config/serviceprovider"
)

/*
* All the bindings are invoked here.
* Then, the application is started.
 */
func main() {

	// Services are binded first and
	// handlers the second because there will be
	// null pointer exception if the services doesn't exist
	// by the time the handlers are binded.
	serviceprovider.BindService()
	handlerprovider.BindHandler()

	// The application is started here.
	app.StartServer()
}
