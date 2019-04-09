package models

import (
	"fmt"
	"time"

	u "github.com/pedrohmachado/brejaria-backend/utils"
)

// Evento modelo
type Evento struct {
	ID            uint      `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Nome          string    `gorm:"not null" json:"nome"`
	Descricao     string    `gorm:"not null" json:"descricao"`
	Local         string    `gorm:"not null" json:"local"`
	DataEvento    string    `gorm:"not null" json:"data_evento"`
	DataCriacao   string    `gorm:"not null" json:"data_criacao"`
	Status        string    `gorm:"not null" json:"status"`
	IDUsuario     uint      `gorm:"not null" json:"usuario_id"`
	Participantes []Usuario `gorm:"many2many:usuario_evento;" json:"participantes"`
}

type Participantes struct {
	Participantes []
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
	layout := "02/01/2006 15:04"

	// dados de entrada fixos
	evento.DataCriacao = time.Now().Format(layout)
	evento.Status = "ativo"
	evento.Participantes = append(evento.Participantes, evento.IDUsuario)

	db := InitDB()
	defer db.Close()

	db.Create(&evento)

	resp := u.Message(true, "Evento cadastrado com sucesso")
	resp["evento"] = evento
	return resp
}

// AdicionaParticipante a um evento pelo id do evento e id do participante
func AdicionaParticipante(IDEvento, IDParticipante uint) map[string]interface{} {

	evento := GetEvento(IDEvento)
	//evento.IDParticipantes = append(evento.IDParticipantes, IDParticipante)

	db := InitDB()
	defer db.Close()

	err := db.Save(&evento).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	resp := u.Message(true, "Participante adicionado com sucesso")
	resp["evento"] = evento
	return resp
}

// GetEvento localiza o evento pelo seu id
func GetEvento(ID uint) *Evento {
	evento := &Evento{}

	db := InitDB()
	defer db.Close()

	err := db.Table("eventos").Where("id = ?", ID).First(evento).Error

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

	err := db.Table("eventos").Where("id_usuario = ?", IDUsuario).First(&eventos).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return eventos
}

// GetEventos localiza todos os eventos pelo status
func GetEventos() []*Evento {
	eventos := make([]*Evento, 0)

	db := InitDB()
	defer db.Close()

	err := db.Find("eventos").Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return eventos
}
