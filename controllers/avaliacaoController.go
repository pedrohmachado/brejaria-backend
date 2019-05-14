package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pedrohmachado/brejaria-backend/models"
	u "github.com/pedrohmachado/brejaria-backend/utils"

	"github.com/gorilla/mux"
)

// AvaliaProduto avalia produto
var AvaliaProduto = func(w http.ResponseWriter, r *http.Request) {

	IDUsuario := r.Context().Value("usuario").(uint)

	params := mux.Vars(r)

	IDProduto, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Print(err)
	}

	avaliacaoProduto := &models.AvaliacaoProduto{}

	err = json.NewDecoder(r.Body).Decode(avaliacaoProduto)

	if err != nil {
		u.Respond(w, u.Message(false, "Erro enquanto decodificava corpo da requisição"))
	}

	avaliacaoProduto.IDProduto = uint(IDProduto)
	avaliacaoProduto.IDUsuario = IDUsuario

	data := avaliacaoProduto.AvaliaProduto()

	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)

}

// AvaliacaoMediaProduto recupera media da avaliacao do produto
var AvaliacaoMediaProduto = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	IDProduto, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Print(err)
	}

	data := models.GetAvaliacaoMediaProduto(uint(IDProduto))

	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}

// AvaliacaoMediaEvento recupera media da avaliacao do produto
var AvaliacaoMediaEvento = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	IDEvento, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Print(err)
	}

	data := models.GetAvaliacaoMediaEvento(uint(IDEvento))

	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}

// AvaliaEvento avalia evento
var AvaliaEvento = func(w http.ResponseWriter, r *http.Request) {

	IDUsuario := r.Context().Value("usuario").(uint)

	params := mux.Vars(r)

	IDEvento, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Print(err)
	}

	avaliacaoEvento := &models.AvaliacaoEvento{}

	err = json.NewDecoder(r.Body).Decode(avaliacaoEvento)

	if err != nil {
		u.Respond(w, u.Message(false, "Erro enquanto decodificava corpo da requisição"))
	}

	avaliacaoEvento.IDEvento = uint(IDEvento)
	avaliacaoEvento.IDUsuario = IDUsuario

	data := avaliacaoEvento.AvaliaEvento()

	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)

}

// GetAvaliacaoEventoUsuario todo
var GetAvaliacaoEventoUsuario = func(w http.ResponseWriter, r *http.Request) {
	IDUsuario := r.Context().Value("usuario").(uint)

	params := mux.Vars(r)

	IDEvento, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Print(err)
	}

	data := models.GetAvaliacaoEventoUsuario(uint(IDEvento), IDUsuario)

	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}

// GetAvaliacaoProdutoUsuario todo
var GetAvaliacaoProdutoUsuario = func(w http.ResponseWriter, r *http.Request) {
	IDUsuario := r.Context().Value("usuario").(uint)

	params := mux.Vars(r)

	IDProduto, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Print(err)
	}

	data := models.GetAvaliacaoProdutoUsuario(uint(IDProduto), IDUsuario)

	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}
