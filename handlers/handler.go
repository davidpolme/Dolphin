package handlers

import (
	"Dolphin/middlewares"
	"Dolphin/routers"

	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Handlers is a function that set the port 8080, the handler and enable the server to listen */
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.CheckDB(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	/* Cors da permiso, para que nuestra API pueda ser accedida desde todas las direcciones ip*/
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
