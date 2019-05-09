package models

import (
	"github.com/jinzhu/gorm"
	u "github.com/pedrohmachado/brejaria-backend/utils"
)

// ImagemProduto estrutura
type ImagemProduto struct {
	ID            uint   `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	CaminhoImagem string `gorm:"not null"`
	IDProduto     uint   `gorm:"not null"`
}

// Cria relação da imagem com produto
func (imagemProduto *ImagemProduto) Cria() map[string]interface{} {

	db := InitDB()
	defer db.Close()

	// criar tabela caso não exista
	if !db.HasTable(&ImagemProduto{}) {
		db.CreateTable(&ImagemProduto{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&ImagemProduto{})
	}

	temp := &ImagemProduto{}

	err := db.Table("imagem_produtos").Where("id_produto = ?", imagemProduto.IDProduto).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Erro de conexão. Tente novamente")
	}

	db.Create(&imagemProduto)

	resp := u.Message(true, "Registro de imagemProduto efetuado com sucesso")
	resp["imagemProduto"] = imagemProduto
	return resp
}

// GetImagemProduto recupera caminho da imagem pelo id do produto
func GetImagemProduto(IDProduto uint) string {
	db := InitDB()
	defer db.Close()

	imagemProduto := &ImagemProduto{}

	db.Table("imagem_produtos").Where("id_produto = ?", IDProduto).Last(imagemProduto)

	return imagemProduto.CaminhoImagem
}
