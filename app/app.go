package app

import (
	"fmt"
	"log"
	"net/http"
	"v1/config/handlerprovider"
	"v1/config/kafka"
	"v1/util"
)

/*
* ' package app ' is the most important package.
 */

/*
* Start the server and listen at the
* port ' 8080 ' by default or customized port
* which is defined at ' .env '.
 */
func StartServer() {
	serverPort := ":80"

	if util.GetEvnValue("SERVER_PORT") != "" {
		log.Println("Custom server port was found! ")
		serverPort = util.GetEvnValue("SERVER_PORT")
	}

	startRoute()

	showBanner()

	// This is deticated go routines for kafka receiver.
	go func() {
		log.Printf("This is deticated go routine.")
		kafka.Receive()
	}()

	log.Printf("Starting the server and listening at port >>> %s", serverPort)
	http.ListenAndServe(serverPort, nil)

}

// Start the actual routing.
func startRoute() {
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
