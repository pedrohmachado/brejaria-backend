package models

import (
	"fmt"
	"time"

	u "github.com/pedrohmachado/brejaria-backend/utils"
)

// Evento modelo
type Evento struct {
	ID            uint       `gorm:"AUTO_INCREMENT;primary_key:true" form:"id" json:"id"`
	Nome          string     `gorm:"not null" json:"nome"`
	Descricao     string     `gorm:"not null" json:"descricao"`
	Local         string     `gorm:"not null" json:"local"`
	DataEvento    string     `gorm:"not null" json:"data_evento"`
	DataCriacao   string     `gorm:"not null" json:"data_criacao"`
	Status        string     `gorm:"not null" json:"status"`
	IDUsuario     uint       `gorm:"not null" json:"usuario_id"`
	Participantes []*Usuario `gorm:"many2many:evento_usuarios" json:"participantes"`
}

// Valida dados de entrada do evento
func (evento *Evento) Valida() (map[string]interface{}, bool) {
	if evento.Nome == "" {
		return u.Message(false, "O evento precisa ter um nome"), false
	}

	if evento.Descricao == "" {
		return u.Message(false, "O evento precisa ter uma descrição"), false
	}

	if evento.IDUsuario <= 0 {
		return u.Message(false, "Usuário não foi reconhecido"), false
	}

	if evento.Local == "" {
		return u.Message(false, "O evento precisa ter um local"), false
	}

	if evento.DataEvento == "" {
		return u.Message(false, "O evento precisa ter uma data"), false
	}

	db := InitDB()
	defer db.Close()

	// criar tabela caso não exista
	if !db.HasTable(&Evento{}) {
		db.CreateTable(&Evento{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Evento{})
	}

	return u.Message(true, "Requisição aprovada"), true
}

// Cria evento relacionado a usuario
func (evento *Evento) Cria() map[string]interface{} {

	if resp, ok := evento.Valida(); !ok {
		return resp
	}

	// layout do formato de data/hora
	layout := "2006-01-02 15:04"

	// dados de entrada fixos
	evento.DataCriacao = time.Now().Format(layout)
	evento.Status = "ativo"

	// usuario criador como participante
	participante := GetUsuario(evento.IDUsuario)
	evento.Participantes = append(evento.Participantes, participante)

	//bd inserir e associação
	db := InitDB()
	defer db.Close()

	db.Create(&evento)
	db.Preload("Participantes").Find(&evento)

	resp := u.Message(true, "Evento cadastrado com sucesso")
	resp["evento"] = evento
	return resp
}

// AlteraEvento altera o evento
func AlteraEvento(evento *Evento) map[string]interface{} {
	if resp, ok := evento.Valida(); !ok {
		return resp
	}

	db := InitDB()
	defer db.Close()

	db.Save(&evento).Where("id = ?", evento.ID)

	resp := u.Message(true, "Evento alterado com sucesso")
	resp["evento"] = evento
	return resp
}

// AdicionaParticipante a um evento pelo id do evento e id do participante
func AdicionaParticipante(IDEvento, IDParticipante uint) map[string]interface{} {

	evento := GetEvento(IDEvento)
	participante := GetUsuario(IDParticipante)
	evento.Participantes = append(evento.Participantes, participante)

	db := InitDB()
	defer db.Close()

	db.Preload("Participantes").Find(&evento)
	db.Save(&evento)
	db.Preload("Participantes").Find(&evento)

	resp := u.Message(true, "Participante adicionado com sucesso")
	resp["evento"] = evento
	return resp
}

// RemoveParticipante de um evento pelo id do evento e id do participante
func RemoveParticipante(IDEvento, IDParticipante uint) map[string]interface{} {

	evento := GetEvento(IDEvento)
	participante := &Usuario{}

	db := InitDB()
	defer db.Close()

	db.Table("usuarios").Where("id = ?", IDParticipante).First(&participante)

	if participante.Email == "" { // usuario não foi encontrado
		return nil
	}

	participante.Senha = ""

	//participantes := evento.Participantes

	// db.Preload("Participantes").Find(&evento)
	// remover participante
	// participantes := evento.Participantes
	db.Model(&evento).Association("Participantes").Delete(&participante)

	db.Save(&evento)

	db.Preload("Participantes").Find(&evento)

	resp := u.Message(true, "Participante removido com sucesso")
	resp["evento"] = evento
	return resp
}

// GetEvento localiza o evento pelo seu id
func GetEvento(ID uint) *Evento {
	evento := &Evento{}

	db := InitDB()
	defer db.Close()

	err := db.Table("eventos").Where("id = ?", ID).First(&evento).Error
	db.Preload("Participantes").Find(&evento)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return evento
}

// GetEventosUsuario localiza todos os eventos de um usuario pelo id do usuario
func GetEventosUsuario(IDUsuario uint) []*Evento {
	eventos := make([]*Evento, 0)

	db := InitDB()
	defer db.Close()

	err := db.Preload("Participantes").Where("id_usuario = ?", IDUsuario).Find(&eventos).Error
	//err := db.Table("eventos").Where("id_usuario = ?", IDUsuario).Find(&eventos).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return eventos
}

// GetEventos localiza todos os eventos
func GetEventos() []*Evento {
	//evento := &Evento{}
	eventos := make([]*Evento, 0)

	db := InitDB()
	defer db.Close()

	db.Preload("Participantes").Find(&eventos)
	err := db.Table("eventos").Find(&eventos).Error
	db.Preload("Participantes").Find(&eventos)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return eventos
}

// GetParticipantesEvento lista todos os participantes de um evento pelo id do evento
func GetParticipantesEvento(IDEvento uint) []*Usuario {

	evento := GetEvento(IDEvento)
	participantes := make([]*Usuario, 0)

	db := InitDB()
	defer db.Close()

	db.Preload("Participantes").First(&evento)
	db.Table("usuarios").Joins("inner join evento_usuarios on evento_usuarios.usuario_id = usuarios.id").Where("evento_usuarios.evento_id = ?", IDEvento).Scan(&participantes)
	db.Preload("Participantes").First(&evento)

	return participantes
}

// GetEventosProduto recupera todos os eventos relacionados ao ID do produto
// encontrar o produto, encontrar o dono do produto e encontrar em quantos eventos o dono do produto está inscrito
func GetEventosProduto(IDProduto uint) []*Evento {

	produto := GetProduto(IDProduto)
	eventos := make([]*Evento, 0)

	db := InitDB()
	defer db.Close()

	err := db.Preload("Participantes").Where("id_usuario = ?", produto.IDUsuario).Find(&eventos).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return eventos
}

// GetEventosParticipante recupera todos os eventos em que um participante está inscrito pelo id do usuario
func GetEventosParticipante(IDUsuario uint) map[string]interface{} {

	eventosInscritos := make([]*Evento, 0)

	db := InitDB()
	defer db.Close()

	//err := db.Preload("Participantes").Where("id_usuario = ?", IDUsuario).Find(&eventosInscritos).Error
	db.Table("eventos").Joins("inner join evento_usuarios on evento_usuarios.evento_id = eventos.id").Where("evento_usuarios.usuario_id = ?", IDUsuario).Scan(&eventosInscritos)

	resp := u.Message(true, "Eventos recuperados com sucesso")
	resp["eventosInscritos"] = eventosInscritos
	return resp
}

// GetEventosProdutor recupera eventos que um usuario produtor/geral possui
func GetEventosProdutor(IDUsuario uint) map[string]interface{} {

	eventos := make([]*Evento, 0)

	db := InitDB()
	defer db.Close()

	//err := db.Preload("Participantes").Where("id_usuario = ?", IDUsuario).Find(&eventosInscritos).Error
	db.Table("eventos").Joins("inner join evento_usuarios on evento_usuarios.evento_id = eventos.id").Where("evento_usuarios.usuario_id = ?", IDUsuario).Scan(&eventos)

	resp := u.Message(true, "Eventos recuperados com sucesso")
	resp["eventos"] = eventos
	return resp
}

// GetCriadorEvento recupera dados do criador do evento pelo id do evento
func GetCriadorEvento(IDEvento uint) map[string]interface{} {
	criador := &Usuario{}

	db := InitDB()
	defer db.Close()

	db.Table("usuarios").Joins("inner join eventos on eventos.id_usuario = usuarios.id").Where("eventos.id = ?", IDEvento).Scan(&criador)

	resp := u.Message(true, "Criador do evento recuperado com sucesso")
	resp["criador"] = criador
	return resp
}
