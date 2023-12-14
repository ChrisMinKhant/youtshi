package app

import (
	"fmt"
	"log"
	"net/http"
	"v1/config/handlerprovider"
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
                   .                          +
      +                                                    .
                                ___       .
.                        _.--"~~ __"-.            
                      ,-"     .-~  ~"-\          
         .          .^       /       ( )      . 
               +   {_.---._ /         ~        
                   /    .  Y                   
                  /      \_j                   
   .             Y     ( --l__                 
                 |            "-.              
                 |      (___     \             
         .       |        .)~-.__/             
                 l        _)
.                 \      "l                                
    +              \       \                               
                    \       ^.                             
        .            ^.       "-.                   . 
                       "-._      ~-.___,                   
                 .         "--.._____.^                    
  .                                         .    
	 _             ____ _                 
	| |    __ _   / ___| |__   __ _ _ __  
	| |   / _\ | | |   | '_ \ / _\ | '_ \ 
	| |__| (_| | | |___| | | | (_| | | | |
	|_____\__,_|  \____|_| |_|\__,_|_| |_|`

	fmt.Println(banner + "\n")
}
