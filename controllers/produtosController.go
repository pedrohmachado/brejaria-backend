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

// CriaProduto com o ID do usuario por contexto
var CriaProduto = func(w http.ResponseWriter, r *http.Request) {

	IDUsuario := r.Context().Value("usuario").(uint)
	produto := &models.Produto{}

	err := json.NewDecoder(r.Body).Decode(produto)
	if err != nil {
		u.Respond(w, u.Message(false, "Erro enquanto decodificava corpo da requisição"))
	}

	produto.IDUsuario = IDUsuario
	resp := produto.Cria()
	u.Respond(w, resp)
}

// GetProdutosParams com o ID do usuario por parametro
var GetProdutosParams = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	IDUsuario, err := strconv.Atoi(params["id"])

	if err != nil {
		u.Respond(w, u.Message(false, "Erro na requisição"))
		return
	}

	data := models.GetProdutosUsuario(uint(IDUsuario))
	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}

// GetMeusProdutos com o ID do usuario por contexto
var GetMeusProdutos = func(w http.ResponseWriter, r *http.Request) {

	IDUsuario := r.Context().Value("usuario").(uint)

	data := models.GetProdutosUsuario(uint(IDUsuario))
	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}

// GetProdutos lista todos os produtos
var GetProdutos = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetProdutos()

	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}

// GetProdutosEvento lista todos os produtos de um evento
var GetProdutosEvento = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	IDEvento, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Print(err)
	}

	data := models.GetProdutosEvento(uint(IDEvento))

	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}
