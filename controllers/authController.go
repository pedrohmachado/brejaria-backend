package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/pedrohmachado/brejaria-backend/models"
	u "github.com/pedrohmachado/brejaria-backend/utils"
)

// CriaUsuario cria usuario
var CriaUsuario = func(w http.ResponseWriter, r *http.Request) {

	usuario := &models.Usuario{}
	err := json.NewDecoder(r.Body).Decode(usuario) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := usuario.Cria() //Create account
	u.Respond(w, resp)
}
