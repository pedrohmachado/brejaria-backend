package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

// GetProdutores recupera produtores
var GetProdutores = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetProdutores()
	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}

// GetProdutor recupera usuario perfil produtor/geral pelo id do usuario
var GetProdutor = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	IDProdutor, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Print(err)
	}

	data := models.GetProdutor(uint(IDProdutor))
	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}
