package service

import "v1/model"

type NotifyService struct {
	BusNumber      int
	ArrivedAddress string
}

var websockerManager = NewManager()

func NewNotifyService() *NotifyService {
	return &NotifyService{}
}

func (notifyService *NotifyService) SetBusNumber(busNumber int) {
	notifyService.BusNumber = busNumber
}

func (notifyService *NotifyService) SetArrivedAddress(arriveAddress string) {
	notifyService.ArrivedAddress = arriveAddress
}

// Send notification to client through websocket
func (notifyService *NotifyService) SendNotification() *model.Error {
	return websockerManager.sendNotification(notifyService.BusNumber, notifyService.ArrivedAddress)
}
