package handler

import (
	"log"
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

func NewNotifyHandler() *NotifyHandler {
	return &NotifyHandler{}
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

	notifyService := serviceprovider.GetService("notifyService").(*service.NotifyService)

	util.DecodeJson(r.Body, notifyHandler)

	log.Printf("Fetched notify request >>> %v", notifyHandler)

	notifyService.SetBusNumber(notifyHandler.BusNumber)
	notifyService.SetArrivedAddress(notifyHandler.ArrivedAddress)
	err := notifyService.SendNotification()
	if err != nil {
		util.ParseResponse(w, err, err.Status)
	} else {

		successResponse := model.SuccessResponse{}

		successResponse.SetStatus(200)
		successResponse.SetMessage("OK")

		util.ParseResponse(w, successResponse, 200)
	}
}
