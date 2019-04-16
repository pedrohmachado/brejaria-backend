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

	err := json.NewDecoder(r.Body).Decode(usuario) // decoda a requisição
	if err != nil {
		u.Respond(w, u.Message(false, "Requisição inválida"))
		return
	}

	resp := usuario.Cria() // cria usuario
	u.Respond(w, resp)
}

// Autentica o login do usuario
var Autentica = func(w http.ResponseWriter, r *http.Request) {

	usuario := &models.Usuario{}

	err := json.NewDecoder(r.Body).Decode(usuario) // decoda a requisição

	if err != nil {
		u.Respond(w, u.Message(false, "Requisição inválida"))
		return
	}

	//resp, token := models.Login(usuario.Email, usuario.Senha)
	resp := models.Login(usuario.Email, usuario.Senha)

	u.Respond(w, resp)
}
