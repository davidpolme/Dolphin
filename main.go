package main

import (
	"log"
	"Dolphin/db"
	"Dolphin/handlers"
)

func main()  {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la Base de datos")
		return
	}
	handlers.Handlers()
}