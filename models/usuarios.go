package models

import (
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	u "github.com/pedrohmachado/brejaria-backend/utils"
)

// Token modelo
type Token struct {
	IDUsuario uint
	jwt.StandardClaims
}

// Usuario modelo
type Usuario struct {
	ID    uint   `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Email string `gorm:"not null" json: "email"`
	Senha string `gorm:"not null" json: "senha"`
	Token string `gorm:"not null" json: "token; sql:"-"`
}

// Valida valida dados de usuario
func (usuario *Usuario) Valida() (map[string]interface{}, bool) {

	if !strings.Contains(usuario.Email, "@") {
		return u.Message(false, "E-mail válido requer '@'"), false
	}

	if len(usuario.Senha) < 6 {
		return u.Message(false, "Senha válida requer no mínimo 6 caracteres"), false
	}

	temp := &Usuario{}

	// bd => checar erros e duplicidade

	if temp.Email != "" {
		return u.Message(false, "E-mail já foi cadastrado"), false
	}

	return u.Message(false, "Requisição aprovada"), true
}

// Cria usuario
func (usuario *Usuario) Cria() map[string]interface{} {

	if resp, ok := usuario.Valida(); !ok {
		return resp
	}

	hashedSenha, _ := bcrypt.GenerateFromPassword([]byte(usuario.Senha), bcrypt.DefaultCost)
	usuario.Senha = string(hashedSenha)

	// bd => criar registro na tabela de usuarios
	db := InitDB()
	defer db.Close()

	if !db.HasTable(&Usuario{}) {
		db.CreateTable(&Usuario{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Usuario{})
	}

	db.Create(&usuario)

	if usuario.ID <= 0 {
		return u.Message(false, "Falha no cadastro, erro de conexão")
	}

	tk := &Token{IDUsuario: usuario.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_SENHA")))

	usuario.Token = tokenString

	usuario.Senha = ""

	resp := u.Message(true, "Usuário logado")
	resp["usuario"] = usuario
	return resp
}
