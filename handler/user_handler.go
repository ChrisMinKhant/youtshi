package handler

import (
	"log"
	"net/http"
	"v1/config/serviceprovider"
	"v1/model"
	"v1/service"
	"v1/util"
)

type UserHandler struct {
	Id          string
	Name        string
	BusId       int
	PickUpPoint string
}

var userService service.UserService

func init() {
	go func() {
		userService = serviceprovider.GetService("userService").(service.UserService)
	}()
}

// The main entry point of handler
func (userHandler *UserHandler) Handle(w http.ResponseWriter, r *http.Request) {

	log.Print("Reached Handle Function.")

	userHandler.userHandlerGroup()

	findTheFunction(r.Method, w, r)
}

// Grouping the register under the same api name "/users"
func (userHandler *UserHandler) userHandlerGroup() {
	funcMap["POST"] = userHandler.RegisterUser
}

// Registering new user to the database ( users, buses )
func (userHandler *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {

	util.DecodeJson(r.Body, userHandler)

	user := model.User{}

	user.SetId(userHandler.Id)
	user.SetName(userHandler.Name)
	user.SetBusId(userHandler.BusId)
	user.SetPickUpPoint(userHandler.PickUpPoint)

	status := userService.RegisterNewUser(user)

	Response(&status, w)

}
