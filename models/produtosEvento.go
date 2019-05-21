package models

import u "github.com/pedrohmachado/brejaria-backend/utils"

// ProdutoEvento modelo
type ProdutoEvento struct {
	ID        uint `gorm:"AUTO_INCREMENT;primary_key:true" json:"id"`
	IDProduto uint `gorm:"not null" json:"id_produto"`
	IDEvento  uint `gorm:"not null" json:"id_evento"`
}

// AdicionaProdutosEvento adiciona produto a evento
func AdicionaProdutosEvento(IDEvento uint, produtos []*Produto) map[string]interface{} {

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
		db.Create(&produtoEvento)
		produtosEvento = append(produtosEvento, produtoEvento)
		produtoEvento = &ProdutoEvento{}
	}

	resp := u.Message(true, "Produtos relacionados a evento inseridos com sucesso")
	resp["produtosEvento"] = produtosEvento
	return resp

}

// RemoveProdutosEvento exclui produtos relacionado a evento
func RemoveProdutosEvento(IDEvento uint, produtos []*Produto) map[string]interface{} {
	produtosEvento := make([]*ProdutoEvento, 0)

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
