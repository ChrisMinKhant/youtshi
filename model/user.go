package model

import (
	"v1/util"
)

type User struct {
	id          string
	name        string
	pickUpPoint string
}

func (user *User) SetId(id string) {
	user.id = id
}

func (user *User) SetName(name string) {
	user.name = name
}

func (user *User) SetPickUpPoint(pickUpPoint string) {
	user.pickUpPoint = pickUpPoint
}

func (user *User) GetId() string {
	return user.id
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) GetPickUpPoint() string {
	return user.pickUpPoint
}

func (user *User) Get() []any {
	return []any{user.id, user.name, user.pickUpPoint}
}

// Register new user to the database.
func (user *User) RegisterNewUser(newUser User) bool {
	user.SetId(newUser.id)
	user.SetName(newUser.name)
	user.SetPickUpPoint(newUser.pickUpPoint)

	util.BuildCreateQuery("Users", []string{"id", "name", "pickUpPoint"}, []any{user.GetId(), user.GetName(), user.GetPickUpPoint()})

	return true
}
