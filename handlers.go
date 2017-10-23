package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pmoule/go2hal/hal"
	"github.com/rs/xid"
)

type Resource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var resources = make(map[string]Resource)

func CreateResource(w http.ResponseWriter, r *http.Request) {
	resource :=
		Resource{Name: "Write presentation", ID: xid.New().String()}

	resources[resource.ID] = resource

	bytes, error := hal.NewEncoder().ToJSON(buildHalResource(&resource))

	if error != nil {
		panic(error)
	} else {
		w.Write(bytes)
	}
}

func buildHalResource(resource *Resource) hal.Resource {
	root := hal.NewResourceObject()
	link := &hal.LinkObject{Href: "/resources/" + resource.ID}
	self := hal.NewSelfLinkRelation()
	self.SetLink(link)
	root.AddLink(self)
	root.AddData(*resource)

	return root
}

func GetSingleResource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resource := resources[vars["id"]]
	bytes, error := hal.NewEncoder().ToJSON(buildHalResource(&resource))

	if error != nil {
		panic(error)
	} else {
		w.Write(bytes)
	}
}
