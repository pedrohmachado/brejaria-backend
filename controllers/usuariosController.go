package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

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

// AlteraUsuario altera valores do usuario pelo id
var AlteraUsuario = func(w http.ResponseWriter, r *http.Request) {

	usuario := &models.Usuario{}

	params := mux.Vars(r)

	novaSenha := params["senha"]

	err := json.NewDecoder(r.Body).Decode(usuario) // decoda a requisição

	if err != nil {
		u.Respond(w, u.Message(false, "Requisição inválida"))
		return
	}

	data := models.AlteraUsuario(usuario, usuario.Senha, novaSenha)
	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}
