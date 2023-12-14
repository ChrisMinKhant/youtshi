package handler

import (
	"net/http"
	"v1/model"
	"v1/util"
)

type UserHandler struct {
	Id          string
	Name        string
	PickUpPoint string
}

func (userHandler *UserHandler) Handle(w http.ResponseWriter, r *http.Request) {

	util.DecodeJson(r.Body, userHandler)

	user := model.User{}

	user.SetId(userHandler.Id)
	user.SetName(userHandler.Name)
	user.SetPickUpPoint(userHandler.PickUpPoint)

	user.RegisterNewUser(user)

}
