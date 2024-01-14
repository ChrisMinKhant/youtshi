package service

import (
	"log"
	"v1/util"

	"github.com/guregu/null"
)

type BusService struct {
	foundBus Bus
}

type Bus struct {
	id              int
	latest_location null.String
}

func NewBusService() *BusService {
	return &BusService{}
}

func (busService *BusService) IsBusExist(id int) bool {
	queryResult := util.BuildSelectQuery("buses", "id", []any{id})

	if !queryResult.Next() {
		return false
	}

	return true
}

func (busService *BusService) RegisterNewBus(id int) {
	util.BuildCreateQuery("buses", []string{"id", "latest_location"}, []any{id, nil})
}

func (busService *BusService) UpdateBusInfo(id int, latestLocation string) {
	if !busService.IsBusExist(id) {
		log.Fatalf("Resource Not Found.")
	}

	util.BuildUpdateQuery("buses", []string{"latest_location"}, "id", []any{id}, []any{latestLocation})
}
