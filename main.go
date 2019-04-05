package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// middleware para autenticação com jwt
	//router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Print(err)
	}
}
