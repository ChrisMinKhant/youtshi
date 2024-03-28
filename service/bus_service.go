package service

import (
	"log"
	"v1/model"
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

// Checking if the requested bus exists or not.
func (busService *BusService) IsBusExist(id int) (bool, *model.Error) {

	//Building the mysql fetch query for fetching data from database.
	queryResult, err := util.BuildSelectQuery("buses", "id", []any{id})

	if err != nil {
		log.Printf("Fetched error at 'IsBusExist()' ::: %v,%v,%v", err.Get()...)
		return false, err
	}

	if !queryResult.Next() {
		return false, nil
	}

	return true, nil
}

func (busService *BusService) RegisterNewBus(id int) *model.Error {
	return util.BuildCreateQuery("buses", []string{"id", "latest_location"}, []any{id, nil})
}

func (busService *BusService) UpdateBusInfo(id int, latestLocation string) *model.Error {

	// Check if the reuqested bus exists, before
	// updating it. This is important to avoid null pointer
	// exception.
	if status, err := busService.IsBusExist(id); err == nil {
		if !status {
			return model.NewError().Set(model.NF404, 404, "There is no bus with such id.")
		}
	} else if err != nil {
		return err
	}

	buildUpdateQueryError := util.BuildUpdateQuery("buses", []string{"latest_location"}, "id", []any{id}, []any{latestLocation})

	if buildUpdateQueryError != nil {
		return buildUpdateQueryError
	}

	return nil
}
