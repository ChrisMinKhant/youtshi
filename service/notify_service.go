package service

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
func (notifyService *NotifyService) SendNotification() {
	websocketService.PushNotification(notifyService.BusNumber, notifyService.ArrivedAddress)
}
