package main

import (
	"FilteringService/controller"
	"FilteringService/persistence"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
)

var Pool *redis.Pool = persistence.NewPool()

func main() {

	persistence.Pool = Pool

	r := mux.NewRouter()

	r.HandleFunc("/", controller.PostHandler).Methods("POST")
	r.HandleFunc("/", controller.GetHandler).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
