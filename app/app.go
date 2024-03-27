package app

import (
	"fmt"
	"log"
	"net/http"
	"v1/config/handlerprovider"
	"v1/util"
)

/*
* The whole application is started
* from this " app " package.
 */

func StartServer() {

	// predefined port for the application
	serverPort := ":80"

	// Find if there is any custom port
	// defined at " .ev " file.
	if util.GetEvnValue("PORT") != "" {
		serverPort = util.GetEvnValue("SERVER_PORT")
	}

	startRoute()
	showBanner()

	log.Printf("Starting the server and listening at port ::: %s", serverPort)
	http.ListenAndServe(serverPort, nil)
}

func startRoute() {
	// The paths and respective handler functions are mapped
	// in the " handlerprovider " package. Fetching those paths and
	// handler functions and passed to http's function.
	for path, handler := range *handlerprovider.GetHandler() {
		http.HandleFunc(path, handler.Handle)
	}
}

func showBanner() {
	banner := `
	*******************************************************
	*                                                     *
	* __   __  ___   ___  _   _   ____  _____  ___  _   _ *
	*|  \ /  |/ _ \ / _ \| \ | | |  _ \|  ___)/ _ \| \ | |*
	*|   v   | | | | | | |  \| | | |_) ) |_  | |_| |  \| |*
	*| |\_/| | | | | | | |     | |  _ (|  _) |  _  |     |*
	*| |   | | |_| | |_| | |\  | | |_) ) |___| | | | |\  |*
	*|_|   |_|\___/ \___/|_| \_| |____/|_____)_| |_|_| \_|*
	*                                                     *
	*                                                     *
	*******************************************************`
	fmt.Println(banner + "\n")
}
