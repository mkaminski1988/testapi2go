package main

import (
	"github.com/manyminds/api2go"
	"github.com/tanmayambre/testapi2go/controller"
	"github.com/tanmayambre/testapi2go/model"

	"net/http"
)

func main() {
	api := api2go.NewAPI("v1")

	api.AddResource(model.Author{}, &controller.Author{})
	api.AddResource(model.Blog{}, &controller.Blog{})
	api.AddResource(model.Book{}, &controller.Book{})

	http.ListenAndServe(":3000", api.Handler())

}
