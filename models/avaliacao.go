package models

import u "github.com/pedrohmachado/brejaria-backend/utils"

// AvaliacaoProduto modelo
type AvaliacaoProduto struct {
	ID        uint `gorm:"AUTO_INCREMENT"`
	IDUsuario uint
	IDProduto uint
	Avaliacao float32 `json:"avaliacao"`
}

// AvaliacaoEvento modelo
type AvaliacaoEvento struct {
	ID        uint `gorm:"AUTO_INCREMENT"`
	IDUsuario uint
	IDEvento  uint
	Avaliacao float32 `json:"avaliacao"`
}

// AvaliaProduto usuario avalia produto
func (avaliacaoProduto *AvaliacaoProduto) AvaliaProduto() map[string]interface{} {

	db := InitDB()
	defer db.Close()

	// criar tabela caso não exista
	if !db.HasTable(&AvaliacaoProduto{}) {
		db.CreateTable(&AvaliacaoProduto{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&AvaliacaoProduto{})
	}

	temp := GetAvaliacaoProduto(avaliacaoProduto.IDProduto, avaliacaoProduto.IDUsuario)

	if temp.ID > 0 {
		db.Model(&avaliacaoProduto).Where("id = ?", temp.ID).Update("avaliacao", avaliacaoProduto.Avaliacao)

		resp := u.Message(true, "Produto reavaliado com sucesso ")
		resp["avaliacaoProduto"] = avaliacaoProduto
		return resp
	}

	db.Create(&avaliacaoProduto)

	resp := u.Message(true, "Produto avaliado com sucesso")
	resp["avaliacaoProduto"] = avaliacaoProduto
	return resp
}

// GetAvaliacaoProduto recupera avaliacao de um produto pelo id do produto e id do usuario
func GetAvaliacaoProduto(IDProduto, IDUsuario uint) *AvaliacaoProduto {
	avaliacaoProduto := &AvaliacaoProduto{}

	db := InitDB()
	defer db.Close()

	db.Where("id_produto = ? AND id_usuario = ?", IDProduto, IDUsuario).First(&avaliacaoProduto)

	return avaliacaoProduto
}

// GetAvaliacaoMediaProduto recupera a media da avaliacao de um produto pelo id do produto
func GetAvaliacaoMediaProduto(IDProduto uint) map[string]interface{} {
	//avaliacaoProduto := &AvaliacaoProduto{}
	avaliacoesProduto := make([]*AvaliacaoProduto, 0)

	db := InitDB()
	defer db.Close()

	//db.Where("id_produto = ?", IDProduto).First(&avaliacaoProduto)
	db.Where("id_produto = ?", IDProduto).Find(&avaliacoesProduto)

	var total float32

	for _, avaliacao := range avaliacoesProduto {
		total += avaliacao.Avaliacao
	}

	media := total / float32(len(avaliacoesProduto))

	resp := u.Message(true, "Avaliação média computada com sucesso")
	resp["avaliacaoMediaProduto"] = media
	resp["numAvaliacoes"] = len(avaliacoesProduto)
	return resp
}

// GetAvaliacaoMediaEvento recupera a media da avaliacao de um evento pelo id do evento
func GetAvaliacaoMediaEvento(IDEvento uint) map[string]interface{} {
	//avaliacaoProduto := &AvaliacaoProduto{}
	avaliacoesEvento := make([]*AvaliacaoEvento, 0)

	db := InitDB()
	defer db.Close()

	//db.Where("id_produto = ?", IDProduto).First(&avaliacaoProduto)
	db.Where("id_evento = ?", IDEvento).Find(&avaliacoesEvento)

	var total float32

	for _, avaliacao := range avaliacoesEvento {
		total += avaliacao.Avaliacao
	}

	media := total / float32(len(avaliacoesEvento))

	resp := u.Message(true, "Avaliação média computada com sucesso")
	resp["avaliacaoMediaEvento"] = media
	resp["numAvaliacoes"] = len(avaliacoesEvento)
	return resp
}

// GetAvaliacaoEventoUsuario recupera avaliacao do evento que o usuario realizou
func GetAvaliacaoEventoUsuario(IDEvento, IDUsuario uint) map[string]interface{} {
	db := InitDB()
	defer db.Close()

	avaliacaoEvento := &AvaliacaoEvento{}

	//db.Where("id_produto = ?", IDProduto).First(&avaliacaoProduto)
	db.Where("id_evento = ? AND id_usuario = ?", IDEvento, IDUsuario).First(&avaliacaoEvento)

	resp := u.Message(true, "Avaliação do usuário recuperada com sucesso")
	resp["avaliacaoEventoUsuario"] = avaliacaoEvento
	return resp
}

// GetAvaliacaoProdutoUsuario recupera avaliacao do evento que o usuario realizou
func GetAvaliacaoProdutoUsuario(IDProduto, IDUsuario uint) map[string]interface{} {
	db := InitDB()
	defer db.Close()

	avaliacaoProduto := &AvaliacaoProduto{}

	//db.Where("id_produto = ?", IDProduto).First(&avaliacaoProduto)
	db.Where("id_produto = ? AND id_usuario = ?", IDProduto, IDUsuario).First(&avaliacaoProduto)

	resp := u.Message(true, "Avaliação do usuário recuperada com sucesso")
	resp["avaliacaoProdutoUsuario"] = avaliacaoProduto
	return resp
}

// AvaliaEvento usuario avalia evento
func (avaliacaoEvento *AvaliacaoEvento) AvaliaEvento() map[string]interface{} {
	db := InitDB()
	defer db.Close()

	// criar tabela caso não exista
	if !db.HasTable(&AvaliacaoEvento{}) {
		db.CreateTable(&AvaliacaoEvento{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&AvaliacaoEvento{})
	}

	temp := GetAvaliacaoEvento(avaliacaoEvento.IDEvento, avaliacaoEvento.IDUsuario)

	if temp.ID > 0 {
		db.Model(&avaliacaoEvento).Where("id = ?", temp.ID).Update("avaliacao", avaliacaoEvento.Avaliacao)

		resp := u.Message(true, "Evento reavaliado com sucesso ")
		resp["avaliacaoEvento"] = avaliacaoEvento
		return resp
	}

	db.Create(&avaliacaoEvento)

	resp := u.Message(true, "Evento avaliado com sucesso")
	resp["avaliacaoEvento"] = avaliacaoEvento
	return resp
}

// GetAvaliacaoEvento recupera avaliacao de um evento
func GetAvaliacaoEvento(IDEvento, IDUsuario uint) *AvaliacaoEvento {
	avaliacaoEvento := &AvaliacaoEvento{}

	db := InitDB()
	defer db.Close()

	db.Where("id_evento = ? AND id_usuario = ?", IDEvento, IDUsuario).First(&avaliacaoEvento)

	return avaliacaoEvento
}
