package sync

import "github.com/Brun0Nasc/Frequencias-Backend/services/sync/frequencia"

// IniciarRotinas é responsável por dar inicio a todas as rotinas que o projeto deve executar
func IniciarRotinas() (erro error) {
	go frequencia.GerarFrequencia()
	return
}
