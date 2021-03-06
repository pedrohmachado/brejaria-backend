package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"

	"github.com/pedrohmachado/brejaria-backend/app"

	"github.com/pedrohmachado/brejaria-backend/controllers"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	c := cors.AllowAll()

	// middleware para autenticação com jwt
	router.Use(app.JwtAuthentication)
	// middleware para distinção de perfil
	router.Use(app.PerfilAuthentication)

	router.HandleFunc("/api/usuario/novo", controllers.CriaUsuario).Methods("POST")
	router.HandleFunc("/api/usuario/login", controllers.Autentica).Methods("POST")
	router.HandleFunc("/api/usuario/{senha}/alterar", controllers.AlteraUsuario).Methods("PUT")

	router.HandleFunc("/api/produtores", controllers.GetProdutores).Methods("GET")
	router.HandleFunc("/api/produtor/{id}", controllers.GetProdutor).Methods("GET")
	router.HandleFunc("/api/produtor/{id}/produtos", controllers.GetProdutosProdutor).Methods("GET")
	router.HandleFunc("/api/produtor/{id}/eventos", controllers.GetEventosProdutor).Methods("GET")

	router.HandleFunc("/api/eu", controllers.GetUsuario).Methods("GET")

	router.HandleFunc("/api/produto/novo", controllers.CriaProduto).Methods("POST")
	router.HandleFunc("/api/usuario/{id}/produtos", controllers.GetProdutosParams).Methods("GET")
	router.HandleFunc("/api/eu/produtos", controllers.GetMeusProdutos).Methods("GET")
	router.HandleFunc("/api/produtos", controllers.GetProdutos).Methods("GET")
	router.HandleFunc("/api/produto/{id}", controllers.GetProduto).Methods("GET")
	router.HandleFunc("/api/produto/alterar", controllers.AlteraProduto).Methods("PUT")
	router.HandleFunc("/api/produto/{id}/eventos-ref", controllers.GetEventosRefProduto).Methods("GET")

	router.HandleFunc("/api/produto/{id}/avaliar", controllers.AvaliaProduto).Methods("POST")
	router.HandleFunc("/api/evento/{id}/avaliar", controllers.AvaliaEvento).Methods("POST")
	router.HandleFunc("/api/produto/{id}/avaliacao", controllers.AvaliacaoMediaProduto).Methods("GET")
	router.HandleFunc("/api/evento/{id}/avaliacao", controllers.AvaliacaoMediaEvento).Methods("GET")
	router.HandleFunc("/api/produto/{id}/avaliacao-usuario", controllers.GetAvaliacaoProdutoUsuario).Methods("GET")
	router.HandleFunc("/api/evento/{id}/avaliacao-usuario", controllers.GetAvaliacaoEventoUsuario).Methods("GET")

	router.HandleFunc("/api/evento/{id}/participantes", controllers.GetParticipantesEvento).Methods("GET")
	router.HandleFunc("/api/evento/{id}/produtos", controllers.GetProdutosEvento).Methods("GET")
	router.HandleFunc("/api/eventos/eu", controllers.GetEventosParticipante).Methods("GET")

	router.HandleFunc("/api/evento/novo", controllers.CriaEvento).Methods("POST")
	router.HandleFunc("/api/evento/{id}/participar", controllers.AdicionaParticipante).Methods("POST")
	router.HandleFunc("/api/evento/{id}/remover", controllers.RemoveParticipante).Methods("POST")
	router.HandleFunc("/api/evento/{id}", controllers.GetEvento).Methods("GET")
	router.HandleFunc("/api/usuario/{id}/eventos", controllers.GetEventosUsuarioParams).Methods("GET")
	router.HandleFunc("/api/eu/eventos", controllers.GetMeusEventos).Methods("GET")
	router.HandleFunc("/api/eventos", controllers.GetEventos).Methods("GET")
	router.HandleFunc("/api/produto/{id}/eventos", controllers.GetEventosProduto).Methods("GET")
	router.HandleFunc("/api/evento/alterar", controllers.AlteraEvento).Methods("PUT")
	router.HandleFunc("/api/evento/{id}/adicionar-produtos", controllers.AdicionaProdutosEvento).Methods("POST")
	router.HandleFunc("/api/evento/{id}/remover-produtos", controllers.RemoveProdutosEvento).Methods("POST")
	router.HandleFunc("/api/evento/{id}/produtos-ref", controllers.GetProdutosRefEvento).Methods("GET")
	router.HandleFunc("/api/evento/{id}/criador", controllers.GetCriadorEvento).Methods("GET")

	router.HandleFunc("/api/upload/produto/{id}", controllers.UploadImagemProduto).Methods("POST")
	router.HandleFunc("/api/imagem/produto/{id}", controllers.GetImagemProduto).Methods("GET")

	router.HandleFunc("/api/upload/evento/{id}", controllers.UploadImagemEvento).Methods("POST")
	router.HandleFunc("/api/imagem/evento/{id}", controllers.GetImagemEvento).Methods("GET")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}

	log.Println("Listening on http://localhost:" + port)
	err := http.ListenAndServe(":"+port, c.Handler(router))

	if err != nil {
		fmt.Print(err)
	}
}
