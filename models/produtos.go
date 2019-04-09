package models

import (
	"fmt"

	u "github.com/pedrohmachado/brejaria-backend/utils"
)

// Produto modelo
type Produto struct {
	ID        uint   `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Nome      string `gorm:"not null" json:"nome"`
	Descricao string `gorm:"not null" json:"descricao"`
	IDUsuario uint   `gorm:"not null" json:"usuario_id"`
}

// Valida dados de entrada de produto
func (produto *Produto) Valida() (map[string]interface{}, bool) {
	if produto.Nome == "" {
		return u.Message(false, "O produto precisa ter um nome"), false
	}

	if produto.Descricao == "" {
		return u.Message(false, "O produto precisa ter uma descrição"), false
	}

	if produto.IDUsuario <= 0 {
		return u.Message(false, "Usuário não foi reconhecido"), false
	}

	return u.Message(true, "Requisição aprovada"), true
}

// Cria um produto
func (produto *Produto) Cria() map[string]interface{} {

	if resp, ok := produto.Valida(); !ok {
		return resp
	}

	db := InitDB()
	defer db.Close()

	// criar tabela caso não exista
	if !db.HasTable(&Produto{}) {
		db.CreateTable(&Produto{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Produto{})
	}

	db.Create(&produto)

	resp := u.Message(true, "Produto cadastrado com sucesso")
	resp["produto"] = produto
	return resp
}

// GetProduto localiza o produto pelo seu id
func GetProduto(ID uint) *Produto {
	produto := &Produto{}

	db := InitDB()
	defer db.Close()

	err := db.Table("produtos").Where("id = ?", ID).First(produto).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return produto
}

// GetProdutos localiza todos os produtos de um usuario pelo id do usuario
func GetProdutos(IDUsuario uint) []*Produto {
	produtos := make([]*Produto, 0)

	db := InitDB()
	defer db.Close()

	err := db.Table("produtos").Where("id_usuario = ?", IDUsuario).First(&produtos).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return produtos
}
