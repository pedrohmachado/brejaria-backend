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

// AdicionaProdutosEvento adiciona produtos a evento
var AdicionaProdutosEvento = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IDEvento, err := strconv.Atoi(params["id"])
	produtos := make([]*models.Produto, 0)

	if err != nil {
		fmt.Print(err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&produtos)
	if err != nil {
		u.Respond(w, u.Message(false, "Erro enquanto decodificava corpo da requisição"))
	}

	data := models.AdicionaProdutosEvento(uint(IDEvento), produtos)
	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}

// GetProdutosRefEvento recupera produtos de um evento
var GetProdutosRefEvento = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IDEvento, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Print(err)
		return
	}

	data := models.GetProdutosRefEvento(uint(IDEvento))
	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}

// RemoveProdutosEvento remove produto relacionado a evento
var RemoveProdutosEvento = func(w http.ResponseWriter, r *http.Request) {}
