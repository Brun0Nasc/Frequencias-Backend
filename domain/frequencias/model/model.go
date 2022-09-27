package frequencias

import "time"

type Frequencia struct {
	ID        *int       `json:"id"`
	DataAtual *time.Time `json:"data_atual"`
}

type ListaFrequencias struct {
	Dados []Frequencia `json:"data"`
}
