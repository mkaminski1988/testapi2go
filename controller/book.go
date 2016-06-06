package controller

import (
	"github.com/manyminds/api2go"
	"github.com/tanmayambre/testapi2go/model"
)

type Book struct {
	book map[string]*model.Book
}

// FindOne choc
func (controller *Book) FindOne(ID string, r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: nil}, nil
}

// FindAll to satisfy api2go data source interface
func (controller *Book) FindAll(r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: nil}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (controller *Book) Create(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: nil}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (controller *Book) Delete(id string, r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: nil}, nil
}

//Update stores all changes on the user
func (controller *Book) Update(obj interface{}, r api2go.Request) (api2go.Responder, error) {
	return &Response{Res: nil}, nil
}