package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"miniproject/controller"
)

func main(){

	r := mux.NewRouter()
	r.HandleFunc("/contoh", controller.GetAll).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
