package middlewares

import (
	"Dolphin/db"
	"net/http"
)

//next proximo endpoint
/* CheckDB es el middleware que me permite conocer el estado de la base de datos*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Conexion Perdida con la base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
