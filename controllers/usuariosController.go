package controllers

import (
	"net/http"

	"github.com/pedrohmachado/brejaria-backend/models"
	u "github.com/pedrohmachado/brejaria-backend/utils"
)

// GetUsuario lista usuario pelo id
var GetUsuario = func(w http.ResponseWriter, r *http.Request) {

	IDUsuario := r.Context().Value("usuario").(uint)

	data := models.GetUsuario(uint(IDUsuario))
	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}
