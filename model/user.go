package model

type User struct {
	id          string
	name        string
	busId       int
	pickUpPoint string
}

func (user *User) SetId(id string) {
	user.id = id
}

func (user *User) SetName(name string) {
	user.name = name
}

func (user *User) SetBusId(busId int) {
	user.busId = busId
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

func (user *User) GetBusId() int {
	return user.busId
}

func (user *User) GetPickUpPoint() string {
	return user.pickUpPoint
}

func (user *User) Get() []any {
	return []any{user.id, user.name, user.busId, user.pickUpPoint}
}
