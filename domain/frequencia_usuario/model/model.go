package frequencia_usuario

import "time"

type ReqFrequenciaUsuario struct {
	ID           *int       `json:"id"`
	UsuarioID    *int       `json:"usuario_id"`
	FrequenciaID *int       `json:"frequencia_id"`
	Entrada      *time.Time `json:"entrada"`
	Saida        *time.Time `json:"saida"`
}

type ResFrequenciaUsuario struct {
	ID             *int       `json:"id"`
	UsuarioID      *int       `json:"usuario_id"`
	Nome           *string    `json:"nome"`
	FrequenciaID   *int       `json:"frequencia_id"`
	DataFrequencia *time.Time `json:"data_frequencia"`
	Entrada        *time.Time `json:"entrada"`
	Saida          *time.Time `json:"saida"`
}

type ListaUsuarioFrequencia struct {
	Dados []ResFrequenciaUsuario `json:"data"`
}
