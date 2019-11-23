package main

import (
	"httpvshttp2/server/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/serve", routes.ServeData).Methods("POST")
	server := http.Server{}
	server.Addr = ":8080"
	server.Handler = router
	log.Print(server.ListenAndServeTLS("server.crt", "server.key"))
}
