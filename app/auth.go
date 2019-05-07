package app

import (
	"context"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pedrohmachado/brejaria-backend/models"
	u "github.com/pedrohmachado/brejaria-backend/utils"
)

//JwtAuthentication token
var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/api/usuario/novo", "/api/usuario/login"} // Lista de endpoints que não precisam de autorização
		requestPath := r.URL.Path                                      // caminho atual

		// valida se o caminho atual não precisa de autorização
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		if strings.Contains(requestPath, "/api/imagem/") {
			next.ServeHTTP(w, r)
			return
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization") // pega o token de autorização do header da requisição

		if tokenHeader == "" { // token faltando
			response = u.Message(false, "Token não encontrado")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ") // o token vem no formato `Bearer {token-body}`, checar se está correto
		if len(splitted) != 2 {
			response = u.Message(false, "Inválido/Mal formado token de autenticação")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		tokenPart := splitted[1] // pega a parte do token que realmente importa
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN_SENHA")), nil
		})

		if err != nil { // token mal formado
			response = u.Message(false, "Token de autenticação mal formado")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		if !token.Valid { // token inválido
			response = u.Message(false, "Token não é válido.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		// tudo ok, proceed with the request and set the caller to the user retrieved from the parsed token
		// fmt.Sprintf("Usuario %", tk.IDUsuario) // util para monitoramento
		ctx := context.WithValue(r.Context(), "usuario", tk.IDUsuario)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) // fim do middleware de autenticação
	})
}
