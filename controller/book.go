package controller

import (
	"github.com/manyminds/api2go"
	"github.com/testapi2go/model"
)

type Book struct {
	book map[string]*model.Book
}

func (controller *Book) FindAll(request api2go.Request) (interface{}, error) {
	return nil, api2go.NewHTTPError(nil, "This functionality has not been implemented", 400)
}

func (controller *Book) FindOne(id string, request api2go.Request) (interface{}, error) {
	return nil, api2go.NewHTTPError(nil, "This functionality has not been implemented", 400)
}

func (controller *Book) FindMultiple([]string, api2go.Request) (interface{}, error) {
	return nil, api2go.NewHTTPError(nil, "This functionality has not been implemented", 400)
}

func (controller *Book) Create(object interface{}, request api2go.Request) (string, error) {
	return "", api2go.NewHTTPError(nil, "This functionality has not been implemented", 400)

}

func (controller *Book) Delete(id string, request api2go.Request) error {
	return api2go.NewHTTPError(nil, "This functionality has not been implemented", 400)

}

func (controller *Book) Update(object interface{}, request api2go.Request) error {
	return api2go.NewHTTPError(nil, "This functionality has not been implemented", 400)
}
