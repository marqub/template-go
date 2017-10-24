package rest

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/marqub/template-go/log"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/pmoule/go2hal/hal"
	"github.com/rs/xid"
)

type Resource struct {
	ID   string `json:"-"` //id won't be serialized
	Name string `json:"name"`
}

var resources = make(map[string]Resource)

func CreateResource(w http.ResponseWriter, r *http.Request) {

	var resource Resource
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &resource); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusPreconditionFailed, Text: "Invalid Resource"}); err != nil {
			panic(err)
		}
		return
	}

	resource.ID = xid.New().String()
	resources[resource.ID] = resource
	log.Logger().Info("Create resource :", resource.ID)

	bytes, error := hal.NewEncoder().ToJSON(buildHalResource(&resource))
	w.Header().Set("Content-Type", "application/hal+json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if error != nil {
		panic(error)
	}
	w.Write(bytes)
}

func buildHalResource(resource *Resource) hal.Resource {
	root := hal.NewResourceObject()
	root.AddData(*resource)

	link := &hal.LinkObject{Href: "/resources/" + resource.ID}
	self := hal.NewSelfLinkRelation()
	self.SetLink(link)
	root.AddLink(self)

	return root
}

func GetSingleResource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resource, ok := resources[vars["resourceId"]] //comes from the routes.go declarations
	log.Logger().WithFields(logrus.Fields{"resid": vars["resourceId"]}).Info("Retrieve resource")

	if !ok {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
		return
	}

	bytes, error := hal.NewEncoder().ToJSON(buildHalResource(&resource))
	w.Header().Set("Content-Type", "application/hal+json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if error != nil {
		panic(error)
	}
	w.Write(bytes)
}
