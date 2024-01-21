package service

import (
	"log"
	"math/rand"
	"net/http"
	"time"
	"v1/model"

	"github.com/gorilla/websocket"
)

// Websocket configuration constant
var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
)

// List of client connection
var connectionList = []Client{}

type Client struct {
	busNumber  int
	sessionId  int
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
func (websockerManager *Manager) sendNotification(busNumber int, message string) *model.Error {

	updateBusInfoError := busService.UpdateBusInfo(busNumber, message)

	if updateBusInfoError != nil {
		return updateBusInfoError
	}

	for conn := range connectionList {

		if connectionList[conn].busNumber != busNumber {
			continue
		}

		log.Printf("Fetched wrote websocket message >>> %v", message)
		err := connectionList[conn].connection.WriteJSON(message)

		log.Print("After writing json to websocket...")
		if err != nil {
			return model.NewError().Set(model.I500, 500, err.Error())
		}
	}

	return nil
}

// Establishing the connection and return Client struct
func (websocketManager *Manager) establishConnection(w http.ResponseWriter, r *http.Request, payload *NotifyBus) Client {
	connection, err := websocketUpgrader.Upgrade(w, r, nil)

	// catch panics and return response to the client
	defer foundPanic(connection, model.I500, 500)

	if err != nil {
		log.Printf("Fetched read error >>> %v", err.Error())
		log.Panicf("Found error while upgrading connection >>> %v", err)
	}

	readError := connection.ReadJSON(payload)

	if readError != nil {
		log.Printf("Fetched read error >>> %v", readError.Error())
		log.Panicf("Found error while reading json message from websocket >>> %v", readError)
	}

	// Check if bus exists
	if status, err := busService.IsBusExist(payload.BusNumber); err == nil {
		if !status {
			// If it isn't, create new bus
			registerNewBusError := busService.RegisterNewBus(payload.BusNumber)

			if registerNewBusError != nil {
				log.Printf("Fetched read error >>> %v", registerNewBusError.Get()...)
				log.Panicf(registerNewBusError.ErrorMessage)
			}
		}
	} else if err != nil {
		log.Printf("Fetched bus exist check error >>> %v", err.Get()...)
		log.Panicf(err.ErrorMessage)
	}

	createdClient := &Client{
		busNumber:  payload.BusNumber,
		sessionId:  rand.New(rand.NewSource(time.Now().UnixNano())).Int(),
		connection: connection,
	}

	connection.WriteJSON(*&createdClient.sessionId)

	return *createdClient
}

func foundPanic(connection *websocket.Conn, errorCode string, errorStatus int) {
	if recoveryStatus := recover(); recoveryStatus != nil {
		connection.WriteJSON(model.NewError().Set(errorCode, errorStatus, recoveryStatus.(string)))
	}
}
