package models

import (
	"github.com/jinzhu/gorm"
	u "github.com/pedrohmachado/brejaria-backend/utils"
)

// ProdutoEvento modelo
type ProdutoEvento struct {
	ID        uint `gorm:"AUTO_INCREMENT;primary_key:true" json:"id"`
	IDProduto uint `gorm:"not null" json:"id_produto"`
	IDEvento  uint `gorm:"not null" json:"id_evento"`
}

// AdicionaProdutosEvento adiciona produto a evento
func AdicionaProdutosEvento(IDEvento uint, produtos []*Produto) map[string]interface{} {

	resp := make(map[string]interface{}, 0)
	produtosEvento := make([]*ProdutoEvento, 0)
	produtoEvento := &ProdutoEvento{}

	db := InitDB()
	defer db.Close()

	// criar tabela caso não exista
	if !db.HasTable(&ProdutoEvento{}) {
		db.CreateTable(&ProdutoEvento{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&ProdutoEvento{})
	}

	for _, produto := range produtos {
		produtoEvento.IDEvento = IDEvento
		produtoEvento.IDProduto = produto.ID

		temp := &ProdutoEvento{}
		err := db.Table("produto_eventos").Where("id_produto = ? AND id_evento = ?", produtoEvento.IDProduto, produtoEvento.IDEvento).First(temp).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			resp = u.Message(false, "Erro de conexão. Tente novamente")
		}

		if temp.ID > 0 {
			resp = u.Message(false, "Produto já foi cadastrado")
		} else {
			db.Create(&produtoEvento)
			produtosEvento = append(produtosEvento, produtoEvento)
			produtoEvento = &ProdutoEvento{}
			resp = u.Message(true, "Produtos relacionados a evento inseridos com sucesso")
			resp["produtosEvento"] = produtosEvento
		}
	}

	return resp
}

// RemoveProdutosEvento exclui produtos relacionado a evento
func RemoveProdutosEvento(IDEvento uint, produtos []*Produto) map[string]interface{} {
	produtosEvento := make([]*ProdutoEvento, 0)
	produtoEvento := &ProdutoEvento{}

	db := InitDB()
	defer db.Close()

	for _, produto := range produtos {
		produtoEvento.IDEvento = IDEvento
		produtoEvento.IDProduto = produto.ID
		db.Where("id_evento = ? AND id_produto = ?", produtoEvento.IDEvento, produtoEvento.IDProduto).Delete(&produtoEvento)
		produtosEvento = append(produtosEvento, produtoEvento)
		produtoEvento = &ProdutoEvento{}
	}

	resp := u.Message(true, "Produtos relacionados a evento excluídos com sucesso")
	resp["produtosEvento"] = produtosEvento
	return resp
}

// GetProdutosRefEvento recupera produtos do evento
func GetProdutosRefEvento(IDEvento uint) map[string]interface{} {

	produtos := make([]*Produto, 0)

	db := InitDB()
	defer db.Close()

	db.Table("produtos").Joins("inner join produto_eventos on produto_eventos.id_produto = produtos.id").Where("produto_eventos.id_evento = ?", IDEvento).Scan(&produtos)

	resp := u.Message(true, "Produtos relacionados a evento recuperados com sucesso")
	resp["produtos"] = produtos
	return resp
}

// GetEventosRefProduto recupera eventos do produto
func GetEventosRefProduto(IDProduto uint) map[string]interface{} {
	eventos := make([]*Evento, 0)

	db := InitDB()
	defer db.Close()

	db.Table("eventos").Joins("inner join produto_eventos on produto_eventos.id_evento = eventos.id").Where("produto_eventos.id_produto = ?", IDProduto).Scan(&eventos)

	resp := u.Message(true, "Eventos relacionados a produto recuperados com sucesso")
	resp["eventos"] = eventos
	return resp
}
