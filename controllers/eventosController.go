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

// CriaEvento com o ID do usuario por contexto
var CriaEvento = func(w http.ResponseWriter, r *http.Request) {

	IDUsuario := r.Context().Value("usuario").(uint)
	evento := &models.Evento{}

	err := json.NewDecoder(r.Body).Decode(evento)
	if err != nil {
		u.Respond(w, u.Message(false, "Erro enquanto decodificava corpo da requisição"))
	}

	evento.IDUsuario = IDUsuario
	resp := evento.Cria()
	u.Respond(w, resp)

}

// AdicionaParticipante com o ID do participante por contexto e ID do evento por parametro
var AdicionaParticipante = func(w http.ResponseWriter, r *http.Request) {

	IDParticipante := r.Context().Value("usuario").(uint)

	params := mux.Vars(r)
	IDEvento, err := strconv.Atoi(params["id"])

	if err != nil {
		u.Respond(w, u.Message(false, "Erro na requisição"))
		return
	}

	data := models.AdicionaParticipante(uint(IDEvento), IDParticipante)

	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}

// RemoveParticipante com o ID do participante por contexto e ID do evento por parametro
var RemoveParticipante = func(w http.ResponseWriter, r *http.Request) {

	IDParticipante := r.Context().Value("usuario").(uint)

	params := mux.Vars(r)
	IDEvento, err := strconv.Atoi(params["id"])

	if err != nil {
		u.Respond(w, u.Message(false, "Erro na requisição"))
		return
	}

	data := models.RemoveParticipante(uint(IDEvento), IDParticipante)

	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}

// GetEventos todos
var GetEventos = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetEventos()

	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}

// GetEventosUsuarioParams com o ID do usuario por parametro
var GetEventosUsuarioParams = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IDUsuario, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Print(err)
		return
	}

	data := models.GetEventosUsuario(uint(IDUsuario))
	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}

// GetMeusEventos com o ID do usuario por contexto
var GetMeusEventos = func(w http.ResponseWriter, r *http.Request) {
	IDUsuario := r.Context().Value("usuario").(uint)

	data := models.GetEventosUsuario(uint(IDUsuario))
	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}

// GetParticipantesEvento lista todos os participantes de um evento
var GetParticipantesEvento = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IDEvento, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Print(err)
		return
	}

	data := models.GetParticipantesEvento(uint(IDEvento))
	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}

// GetEvento lista um evento pelo seu id
var GetEvento = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IDEvento, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Print(err)
		return
	}

	data := models.GetEvento(uint(IDEvento))
	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetEventosProduto = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IDProduto, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Print(err)
		return
	}

	data := models.GetEventosProduto(uint(IDProduto))
	resp := u.Message(true, "Sucesso")
	resp["data"] = data
	u.Respond(w, resp)
}
