package usuarios

import "time"

type ReqUsuario struct {
	ID        *int       `json:"id"`
	Tipo      *int       `json:"tipo"`
	Nome      *string    `json:"nome"`
	Email     *string    `json:"email"`
	Senha     string     `json:"senha"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
