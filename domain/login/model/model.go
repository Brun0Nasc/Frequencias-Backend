package login

import "time"

type Login struct {
	ID        uint      `json:"id"`
	Tipo      uint      `json:"tipo"`
	Nome      string    `json:"nome"`
	Email     string    `json:"email"`
	Senha     string    `json:"senha"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	RemovedAt time.Time `json:"removed_at"`
}
