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
	SessionId  int
	connection *websocket.Conn
}

type Manager struct {
}

type NotifyBus struct {
	BusNumber int
	SessionId int
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

	websocketManager.establishConnection(w, r, &NotifyBus{})
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

		if err != nil {
			log.Printf("Error occured at writing message to websocket >>> %v", err.Error())

			connectionList[conn].connection.Close()
		}
	}

	return nil
}

// Establishing the connection and return Client struct
func (websocketManager *Manager) establishConnection(w http.ResponseWriter, r *http.Request, payload *NotifyBus) {
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

	if payload.SessionId != 0 {
		foundFlag := false

		for existedConnection := range connectionList {
			if connectionList[existedConnection].SessionId == payload.SessionId {

				foundFlag = true

				log.Printf("Connection status check >>> %v", connectionList[existedConnection].connection == connection)

				connectionList[existedConnection].busNumber = payload.BusNumber
				connectionList[existedConnection].connection = connection
			}
		}

		createdClient := &Client{
			busNumber:  payload.BusNumber,
			SessionId:  payload.SessionId,
			connection: connection,
		}

		if !foundFlag {
			connectionList = append(connectionList, *createdClient)
		}

		connection.WriteJSON(*createdClient)

	} else {
		createdClient := &Client{
			busNumber:  payload.BusNumber,
			SessionId:  rand.New(rand.NewSource(time.Now().UnixNano())).Int(),
			connection: connection,
		}

		connectionList = append(connectionList, *createdClient)

		connection.WriteJSON(*createdClient)
	}

	log.Printf("Total Connection >>> %v", len(connectionList))

}

func foundPanic(connection *websocket.Conn, errorCode string, errorStatus int) {
	if recoveryStatus := recover(); recoveryStatus != nil {
		connection.WriteJSON(model.NewError().Set(errorCode, errorStatus, recoveryStatus.(string)))
	}
}
