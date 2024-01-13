package handler

import (
	"net/http"
	"v1/config/serviceprovider"
	"v1/model"
	"v1/service"
	"v1/util"
)

type NotifyHandler struct {
	BusNumber      int
	ArrivedAddress string
}

var notifyService service.NotifyService

func init() {
	go func() {
		notifyService = serviceprovider.GetService("notifyService").(service.NotifyService)
		websocketService = serviceprovider.GetService("websocketService").(service.WebsocketService)
	}()
}

func (notifyHandler *NotifyHandler) Handle(w http.ResponseWriter, r *http.Request) {
	notifyHandler.notifyHandlerGroup()
	findTheFunction(r.Method, w, r)
}

// Grouping the register under the same api name "/users"
func (notifyHandler *NotifyHandler) notifyHandlerGroup() {
	funcMap["POST"] = notifyHandler.notify
}

func (notifyHandler *NotifyHandler) notify(w http.ResponseWriter, r *http.Request) {
	util.DecodeJson(r.Body, notifyHandler)

	websocketService.PushNotification(notifyHandler.BusNumber, notifyHandler.ArrivedAddress)
	// kafka.Send(notifyHandler.BusNumber, notifyHandler.ArrivedAddress)

	// notifyService.DropMessageToKafka([]any{notifyHandler.BusNumber, notifyHandler.ArrivedAddress})

	successResponse := model.SuccessResponse{}

	successResponse.SetStatus(200)
	successResponse.SetMessage("OK")

	util.ParseResponse(w, successResponse, 200)
}
