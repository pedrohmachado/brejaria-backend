package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pedrohmachado/brejaria-backend/app"

	"github.com/pedrohmachado/brejaria-backend/controllers"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// middleware para autenticação com jwt
	router.Use(app.JwtAuthentication)

	router.HandleFunc("/api/usuario/novo", controllers.CriaUsuario).Methods("POST")
	router.HandleFunc("/api/usuario/login", controllers.Autentica).Methods("POST")

	router.HandleFunc("/api/produto/novo", controllers.CriaProduto).Methods("POST")
	// listagem com id por parametro
	router.HandleFunc("/api/usuario/{id}/produtos", controllers.GetMeusProdutosParams).Methods("GET")
	// listagem com id por contexto
	router.HandleFunc("/api/eu/produtos", controllers.GetMeusProdutos).Methods("GET")

	router.HandleFunc("/api/evento/novo", controllers.CriaEvento).Methods("POST")
	// listagem com id por contexto
	router.HandleFunc("/api/eu/eventos", controllers.GetMeusEventos).Methods("GET")
	// listagem com id por parametro
	router.HandleFunc("/api/usuario/{id}/eventos", controllers.GetEventosParams).Methods("GET")
	// listagem de todos os eventos
	router.HandleFunc("/api/eventos", controllers.GetEventos).Methods("GET")
	// adiciona participante a evento
	router.HandleFunc("/api/evento/{id}/participante", controllers.AdicionaParticipante).Methods("POST")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}

	log.Println("Listening on http://localhost:" + port)
	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Print(err)
	}
}
