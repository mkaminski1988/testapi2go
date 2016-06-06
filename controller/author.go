package controller

import (
	"github.com/manyminds/api2go"
	"github.com/tanmayambre/testapi2go/model"

)

type Author struct {
	author map[string]*model.Author
}

// FindOne choc
func (controller *Author) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: nil}, nil
}

// FindAll to satisfy api2go data source interface
func (controller *Author) FindAll(r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: nil}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (controller *Author) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: nil}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (controller *Author) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: nil}, nil
}

//Update stores all changes on the user
func (controller *Author) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: nil}, nil
}