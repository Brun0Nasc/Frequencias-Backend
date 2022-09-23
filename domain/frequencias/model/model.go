package frequencias

import "time"

type Frequencia struct {
	ID        *int       `json:"id"`
	UsuarioID *int       `json:"usuario_id"`
	CreatedAt *time.Time `json:"created_at"`
	Nome      *string    `json:"nome"`
}

type ListaFrequencias struct {
	Dados []Frequencia `json:"data"`
}
