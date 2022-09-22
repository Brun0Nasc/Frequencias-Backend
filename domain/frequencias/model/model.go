package frequencias

import "time"

type ReqFrequencia struct {
	ID        *int       `json:"id"`
	UsuarioID *int       `json:"usuario_id"`
	CreatedAt *time.Time `json:"created_at"`
	Nome	  *string	`json:"nome,omitempty"`
}
