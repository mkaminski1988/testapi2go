package controller

import (
	"github.com/manyminds/api2go"
	"github.com/testapi2go/model"
)

type Author struct {
	author map[string]*model.Author
}

func (controller *Author) FindAll(request api2go.Request) (interface{}, error) {
	return nil, api2go.NewHTTPError(nil, "This functionality has not been implemented", 400)
}

func (controller *Author) FindOne(id string, request api2go.Request) (interface{}, error) {
	return nil, api2go.NewHTTPError(nil, "This functionality has not been implemented", 400)
}

func (controller *Author) FindMultiple([]string, api2go.Request) (interface{}, error) {
	return nil, api2go.NewHTTPError(nil, "This functionality has not been implemented", 400)
}

func (controller *Author) Create(object interface{}, request api2go.Request) (string, error) {
	return "", api2go.NewHTTPError(nil, "This functionality has not been implemented", 400)

}

func (controller *Author) Delete(id string, request api2go.Request) error {
	return api2go.NewHTTPError(nil, "This functionality has not been implemented", 400)

}

func (controller *Author) Update(object interface{}, request api2go.Request) error {
	return api2go.NewHTTPError(nil, "This functionality has not been implemented", 400)
}
