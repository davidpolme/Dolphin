package routers

import (
	"Dolphin/db"
	"Dolphin/jwt"
	"Dolphin/models"
	"encoding/json"
	"net/http"
)

/* Login function */
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario y/o Contraseña invalidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}
	document, exist := db.TryLogin(t.Email, t.Password)

	if exist == false {
		http.Error(w, "Usuario y contraseña invalidos", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(document)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar  generar el token correspondiente"+err.Error(), 400)
		return
	}

}
