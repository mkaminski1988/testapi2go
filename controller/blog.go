package controller

import (
	"github.com/manyminds/api2go"
	"github.com/tanmayambre/testapi2go/model"
)

type Blog struct {
	blog map[string]*model.Blog
}

// FindOne choc
func (controller *Blog) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: nil}, nil
}

// FindAll to satisfy api2go data source interface
func (controller *Blog) FindAll(r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: nil}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (controller *Blog) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: nil}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (controller *Blog) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: nil}, nil
}

//Update stores all changes on the user
func (controller *Blog) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: nil}, nil
}