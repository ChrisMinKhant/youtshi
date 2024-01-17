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

func (busService *BusService) IsBusExist(id int) (bool, *model.Error) {

	queryResult, err := util.BuildSelectQuery("buses", "id", []any{id})

	if err != nil {
		log.Printf("There is an error at IsBusExist() >>> %v,%v,%v", err.Get()...)
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
