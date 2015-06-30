package controller

import (
	"github.com/manyminds/api2go"
	"github.com/testapi2go/model"
)

type Blog struct {
	blog map[string]*model.Blog
}

func (controller *Blog) FindAll(request api2go.Request) (interface{}, error) {
	return nil, api2go.NewHTTPError(nil, "This functionality has not been implemented", 400)
}

func (controller *Blog) FindOne(id string, request api2go.Request) (interface{}, error) {
	return nil, api2go.NewHTTPError(nil, "This functionality has not been implemented", 400)
}

func (controller *Blog) FindMultiple([]string, api2go.Request) (interface{}, error) {
	return nil, api2go.NewHTTPError(nil, "This functionality has not been implemented", 400)
}

func (controller *Blog) Create(object interface{}, request api2go.Request) (string, error) {
	return "", api2go.NewHTTPError(nil, "This functionality has not been implemented", 400)

}

func (controller *Blog) Delete(id string, request api2go.Request) error {
	return api2go.NewHTTPError(nil, "This functionality has not been implemented", 400)

}

func (controller *Blog) Update(object interface{}, request api2go.Request) error {
	return api2go.NewHTTPError(nil, "This functionality has not been implemented", 400)
}
