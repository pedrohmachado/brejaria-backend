package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pedrohmachado/brejaria-backend/controllers"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// middleware para autenticação com jwt
	//router.Use(app.JwtAuthentication)

	router.HandleFunc("/api/usuario/novo", controllers.CriaUsuario)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Print(err)
	}
}
