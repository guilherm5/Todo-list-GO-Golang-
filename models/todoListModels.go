package models

import (
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm"
)

type Todo_List struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Tarefa       string    `json:"tarefa"`
	ResumoTarefa string    `json:"resumo_tarefa"`
	Realizado    bool      `json:"realizado"`
}

type Todo_List_DELL struct {
	Ids []uuid.UUID `json:"ids"`
}
