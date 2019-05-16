package app

import (
	"net/http"

	"github.com/pedrohmachado/brejaria-backend/models"

	u "github.com/pedrohmachado/brejaria-backend/utils"
)

// PerfilAuthentication verifica rotas que podem ser acessadas por perfis distintos
var PerfilAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// definir rotas que não são afetadas pela distinção
		notAuth := []string{"/api/usuario/novo", "/api/usuario/login"}
		// definir rotas que produtor NÃO pode acessar
		produtorNotAuth := []string{""}
		// definir rotas que consumidor NÃO pode acessar
		consumidorNotAuth := []string{"/api/produto/novo", "/api/eu/produtos"}

		requestPath := r.URL.Path

		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		// recupera perfil do usuario
		// perfil 1 consumidor
		// perfil 2 produtor
		IDUsuario := r.Context().Value("usuario").(uint)
		usuario := models.GetUsuario(IDUsuario)
		perfil := usuario.Perfil

		// perfil := "consumidor"
		// perfil := "produtor"
		// perfil := "geral"

		// formato da resposta
		response := make(map[string]interface{})

		if perfil == "" {
			response = u.Message(false, "Perfil do usuário não foi encontrado")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		// caso o usuario tenha os dois perfis
		if perfil == "geral" {
			next.ServeHTTP(w, r)
			return
		}

		// caso o usuario for consumidor
		if perfil == "consumidor" {
			for _, value := range consumidorNotAuth {
				if value == requestPath {
					response = u.Message(false, "Usuário com perfil consumidor não autorizado a acessar essa página")
					w.WriteHeader(http.StatusForbidden)
					w.Header().Add("Content-Type", "application/json")
					u.Respond(w, response)
					return
				}
			}
			next.ServeHTTP(w, r)
			return
		}

		// caso o usuario for produtor

		if perfil == "produtor" {
			for _, value := range produtorNotAuth {
				if value == requestPath {
					response = u.Message(false, "Usuário com perfil produtor não autorizado a acessar essa página")
					w.WriteHeader(http.StatusForbidden)
					w.Header().Add("Content-Type", "application/json")
					u.Respond(w, response)
					return
				}
			}
			next.ServeHTTP(w, r)
			return
		}

	})
}
