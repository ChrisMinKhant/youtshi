package service

import (
	"v1/model"
	"v1/util"
)

type UserService struct {
}

func (userService *UserService) RegisterNewUser(requestedUser model.User) bool {

	queryResult := util.BuildSelectQuery("buses", "id", []any{requestedUser.GetBusId()})

	// Bus exists in the database
	if queryResult.Next() {
		util.BuildCreateQuery("users", []string{"id", "name", "bus_id", "pickUpPoint"},
			[]any{requestedUser.GetId(), requestedUser.GetName(), requestedUser.GetBusId(), requestedUser.GetPickUpPoint()})

		return true
	}

	// Bus doesn't exist in the database
	util.BuildCreateQuery("buses", []string{"id"}, []any{requestedUser.GetBusId()})

	util.BuildCreateQuery("users", []string{"id", "name", "bus_id", "pickUpPoint"},
		[]any{requestedUser.GetId(), requestedUser.GetName(), requestedUser.GetBusId(), requestedUser.GetPickUpPoint()})

	return true
}
