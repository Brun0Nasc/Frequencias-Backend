package frequencia_usuario

import "time"

type FrequenciaUsuario struct {
	ID           *int
	UsuarioID    *int
	FrequenciaID *int
	Entrada      *time.Time
	Saida        *time.Time
}
