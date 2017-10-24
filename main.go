package main

import (
	"net/http"

	"github.com/marqub/template-go/log"
	"github.com/marqub/template-go/rest"
)

func main() {
	log.Logger().Info("Server started")
	router := rest.NewRouter()
	log.Logger().Fatal(http.ListenAndServe(":8080", router))
	//":"+os.Getenv("PORT")
}
