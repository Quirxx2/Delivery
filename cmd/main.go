package main

import (
	dbcon "github.com/Quirxx2/AndersenPromo/internal"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// @title           Andersen Promo API
// @version         1.0
// @description     API server for Andersen Promo Application

// @host      localhost:8080
// @BasePath  /cmd

func main() {

	serverHost := ":8080"
	connString := "postgres://usr:password@localhost:5432/registry"

	c, err := dbcon.NewHandlers(connString)
	if err != nil {
		log.Fatalf("Failed to create Handlers: %v", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", c.HealthCheck).Methods(http.MethodGet)
	router.HandleFunc("/create", c.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/delete/{id}", c.DeleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/update/{id}", c.UpdateUser).Methods(http.MethodPatch)
	router.HandleFunc("/get/{id}", c.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/getall", c.GetUserList).Methods(http.MethodGet)

	srv := &http.Server{
		Handler: router,
		Addr:    serverHost,
	}

	log.Fatal(srv.ListenAndServe())
}
