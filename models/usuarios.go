package models

import (
	"os"
	"strings"

	"github.com/jinzhu/gorm"

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
	ID      uint   `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Nome    string `gorm:"not null" json:"nome"`
	Email   string `gorm:"not null" json:"email"`
	Senha   string `gorm:"not null" json:"senha,omitempty"`
	Token   string `gorm:"not null" json:"token,omitempty"`
	Perfil  string `gorm:"not null" json:"perfil"`
	Contato string `gorm:"not null" json:"contato"`
	//Eventos []Evento `gorm:"many2many:usuario_evento;" json:"eventos"`
}

// Valida valida dados de usuario
func (usuario *Usuario) Valida() (map[string]interface{}, bool) {

	if !strings.Contains(usuario.Email, "@") {
		return u.Message(false, "E-mail válido requer '@'"), false
	}

	if len(usuario.Senha) < 6 {
		return u.Message(false, "Senha válida requer no mínimo 6 caracteres"), false
	}

	if usuario.Nome == "" {
		return u.Message(false, "O preenchimento do nome é obrigatório"), false
	}

	temp := &Usuario{}

	db := InitDB()
	defer db.Close()

	// criar tabela caso não exista
	if !db.HasTable(&Usuario{}) {
		db.CreateTable(&Usuario{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Usuario{})
	}

	// bd => checar erros e duplicidade
	err := db.Table("usuarios").Where("email = ?", usuario.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Erro de conexão. Tente novamente"), false
	}

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

	db := InitDB()
	defer db.Close()

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

// Login usuario
func Login(email, senha string) map[string]interface{} {
	usuario := &Usuario{}

	db := InitDB()
	defer db.Close()

	// localiza usuario
	err := db.Table("usuarios").Where("email = ? ", email).First(&usuario).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email não encontrado")
		}
		return u.Message(false, "Falha na conexão. Tente novamente")
	}

	err = bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(senha))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Credenciais de login inválidas. Tente novamente")
	}

	// passou, login com sucesso
	usuario.Senha = ""

	tk := &Token{IDUsuario: usuario.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_SENHA")))
	usuario.Token = tokenString

	resp := u.Message(true, "Usuário logado")
	resp["usuario"] = usuario
	return resp
}

// GetUsuario pelo id
func GetUsuario(ID uint) *Usuario {
	usuario := &Usuario{}

	db := InitDB()
	defer db.Close()

	db.Table("usuarios").Where("id = ?", ID).First(&usuario)

	if usuario.Email == "" { // usuario não foi encontrado
		return nil
	}

	usuario.Senha = ""
	return usuario
}

// AlteraUsuario pelo registro do usuario
func AlteraUsuario(usuario *Usuario, senha, novaSenha string) map[string]interface{} {

	temp := &Usuario{}

	db := InitDB()
	defer db.Close()

	err := db.Table("usuarios").Where("email = ? ", usuario.Email).First(temp).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email não encontrado")
		}
		return u.Message(false, "Falha na conexão. Tente novamente")
	}

	err = bcrypt.CompareHashAndPassword([]byte(temp.Senha), []byte(senha))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Credenciais inválidas. Tente novamente")
	}

	if novaSenha != "" {
		usuario.Senha = novaSenha
	}

	if !strings.Contains(usuario.Email, "@") {
		return u.Message(false, "E-mail válido requer '@'")
	}

	if len(usuario.Senha) < 6 {
		return u.Message(false, "Senha válida requer no mínimo 6 caracteres")
	}

	if usuario.Nome == "" {
		return u.Message(false, "O preenchimento do nome é obrigatório")
	}

	hashedSenha, _ := bcrypt.GenerateFromPassword([]byte(usuario.Senha), bcrypt.DefaultCost)
	usuario.Senha = string(hashedSenha)

	db.Save(&usuario)

	usuario.Senha = ""

	resp := u.Message(true, "Usuário alterado")
	resp["usuario"] = usuario
	return resp
}

// GetProdutores recupera usuarios perfil produtor e geral
func GetProdutores() map[string]interface{} {
	usuarios := make([]*Usuario, 0)

	db := InitDB()
	defer db.Close()

	db.Table("usuarios").Where("perfil = ? OR perfil = ?", "produtor", "geral").Find(&usuarios)

	resp := u.Message(true, "Produtores recuperados com sucesso")
	resp["usuarios"] = usuarios
	return resp
}

// GetProdutor recupera usuario perfil produtor/geral pelo id do usuario
func GetProdutor(IDUsuario uint) map[string]interface{} {

	usuario := &Usuario{}

	db := InitDB()
	defer db.Close()
	db.Table("usuarios").Where("id = ?", IDUsuario).Find(&usuario)

	resp := u.Message(true, "Produtor recuperado com sucesso")
	resp["usuario"] = usuario
	return resp

}
