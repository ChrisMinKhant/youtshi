package service

import "v1/model"

type NotifyService struct {
	BusNumber      int
	ArrivedAddress string
}

var websocketService = NewWebsocketService()

func (notifyService *NotifyService) SetBusNumber(busNumber int) {
	notifyService.BusNumber = busNumber
}

func (notifyService *NotifyService) SetArrivedAddress(arriveAddress string) {
	notifyService.ArrivedAddress = arriveAddress
}

// Send notification to client through websocket
func (notifyService *NotifyService) SendNotification() *model.Error {
	return websocketService.PushNotification(notifyService.BusNumber, notifyService.ArrivedAddress)
}
