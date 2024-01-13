package service

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Websocket configuration constant
var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

// List of client connection
var connectionList = []Client{}

type Client struct {
	busNumber  int
	connection *websocket.Conn
}

type Manager struct {
}

type NotifyBus struct {
	BusNumber int
}

func NewManager() *Manager {
	return &Manager{}
}

/*
* Entry method to start websocket
* and add connection to client connection list
 */
func (websocketManager *Manager) startWebsocket(w http.ResponseWriter, r *http.Request) {
	log.Printf("Started websocket...")

	connection := websocketManager.establishConnection(w, r, &NotifyBus{})

	connectionList = append(connectionList, connection)

}

// Sending live data to the client
func (websockerManager *Manager) sendNotification(busNumber int, message string) {

	for conn := range connectionList {
		if connectionList[conn].busNumber != busNumber {
			continue
		}

		err := connectionList[conn].connection.WriteJSON(message)

		if err != nil {
			log.Fatalf("Found error while writing json message >>> %v", err)
		}
	}
}

// Establishing the connection and return Client struct
func (websocketManager *Manager) establishConnection(w http.ResponseWriter, r *http.Request, payload *NotifyBus) Client {
	connection, err := websocketUpgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatalf("Found error while upgrading connection >>> %v", err)
	}

	readError := connection.ReadJSON(payload)

	if readError != nil {
		log.Fatalf("Found error while reading json message from websocket >>> %v", readError)
	}

	return Client{
		busNumber:  payload.BusNumber,
		connection: connection,
	}

}
