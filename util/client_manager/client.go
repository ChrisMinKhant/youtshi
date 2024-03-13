package clientmanager

import "github.com/gorilla/websocket"

/*
* Client is the basic struct for
* handling websocket connection.
 */
type Client struct {
	busNumber  int
	SessionId  int
	connection *websocket.Conn
}

func NewClient(requestedBusNumber int, requestedSessionId int, requestedConnection *websocket.Conn) *Client {
	return &Client{
		busNumber:  requestedBusNumber,
		SessionId:  requestedSessionId,
		connection: requestedConnection,
	}
}

func (client *Client) SetBusNumber(busNumber int) {
	client.busNumber = busNumber
}

func (client *Client) SetConnection(connectoin *websocket.Conn) {
	client.connection = connectoin
}

func (client *Client) GetBusNumber() int {
	return client.busNumber
}

func (client *Client) GetConnection() *websocket.Conn {
	return client.connection
}
