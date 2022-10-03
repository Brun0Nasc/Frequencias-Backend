package usuarios

import "time"

type Usuario struct {
	ID        *int
	Tipo      *string
	Nome      *string
	Email     *string
	Senha     *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	RemovedAt *time.Time
}
