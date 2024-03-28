package service

import (
	"log"
	"math/rand"
	"net/http"
	"time"
	"v1/model"
	clientmanager "v1/util/client_manager"

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

/*
* ClientManager for managing websocket connection
 */
var clientManager = clientmanager.NewClientManager()

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

// Send notification to the users with the same bus number
// using websocket connection.
func (websockerManager *Manager) sendNotification(busNumber int, message string) *model.Error {

	// Fetch client from nested linked clients by requested bus number.
	fetchedClientList := clientManager.FetchClientByBusNumber(busNumber)

	// Changed bus arrived address to current arrived address.
	updateBusInfoError := busService.UpdateBusInfo(busNumber, message)

	if updateBusInfoError != nil {
		return updateBusInfoError
	}

	for clientIndex := range fetchedClientList {

		if fetchedClientList[clientIndex].GetBusNumber() != busNumber {
			continue
		}

		log.Printf("Fetched wrote websocket message ::: %v", message)
		err := fetchedClientList[clientIndex].GetConnection().WriteJSON(message)

		if err != nil {
			log.Printf("Error occured at writing message to websocket ::: %v", err.Error())

			fetchedClientList[clientIndex].GetConnection().Close()
		}
	}

	return nil
}

// Establishing the connection and return Client struct
func (websocketManager *Manager) establishConnection(w http.ResponseWriter, r *http.Request, payload *NotifyBus) {

	// Upgrade normal http connection to get websocket connection.
	connection, err := websocketUpgrader.Upgrade(w, r, nil)

	log.Printf("Fetcehd created connection ::: %v", connection.UnderlyingConn().LocalAddr())
	// catch panics and return response to the client
	defer foundPanic(connection, model.I500, 500)

	if err != nil {
		log.Panicf("Found error while upgrading connection ::: %v", err)
	}

	readError := connection.ReadJSON(payload)

	if readError != nil {
		log.Panicf("Found error while reading json message from websocket ::: %v", readError)
	}

	// Bus handling section
	// Check if bus exists
	if status, err := busService.IsBusExist(payload.BusNumber); err == nil {
		if !status {
			// If it isn't, create new bus
			registerNewBusError := busService.RegisterNewBus(payload.BusNumber)

			if registerNewBusError != nil {
				log.Printf("Fetched found error while registering new bus ::: %v", registerNewBusError.Get()...)
				log.Panicf(registerNewBusError.ErrorMessage)
			}
		}
	} else if err != nil {
		log.Printf("Fetched bus exist check error ::: %v", err.Get()...)
		log.Panicf(err.ErrorMessage)
	}

	// Checking if the user is new or not, by using seesion id
	// got from front-end. If the user already existed, the new websocket
	// connection will replace the old one.
	if payload.SessionId != 0 {
		foundFlag := false

		fetchedClientList := clientManager.FetchClientByBusNumber(payload.BusNumber)

		for fetchedClientIndex := range fetchedClientList {

			log.Printf("Fetched session id of found client ::: %v", fetchedClientList[fetchedClientIndex].SessionId)

			if fetchedClientList[fetchedClientIndex].SessionId == payload.SessionId {

				log.Print("Fetched existing session check status ::: [ True ]")

				foundFlag = true

				fetchedClientList[fetchedClientIndex].SetBusNumber(payload.BusNumber)
				fetchedClientList[fetchedClientIndex].SetConnection(connection)

				connection.WriteJSON(fetchedClientList[fetchedClientIndex])
			}
		}

		// Creating new client
		createdClient := clientmanager.NewClient(payload.BusNumber, payload.SessionId, connection)

		// Adding new client to linked client through
		// client manager unless such client already existed.
		if !foundFlag {
			log.Printf("Adding new created client to linked client.")
			clientManager.AddClient(createdClient)
		}

		connection.WriteJSON(createdClient)

	} else {

		generatedSessionId := rand.New(rand.NewSource(time.Now().UnixNano())).Int()

		log.Printf("Fetched generated session id before creating new client ::: %v", generatedSessionId)

		createdClient := clientmanager.NewClient(payload.BusNumber, generatedSessionId, connection)
		clientManager.AddClient(createdClient)

		connection.WriteJSON(*createdClient)
	}
}

func foundPanic(connection *websocket.Conn, errorCode string, errorStatus int) {
	if recoveryStatus := recover(); recoveryStatus != nil {
		connection.WriteJSON(model.NewError().Set(errorCode, errorStatus, recoveryStatus.(string)))
	}
}
