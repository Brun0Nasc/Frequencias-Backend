package frequencia

import (
	"errors"
	"fmt"
	"time"

	"github.com/Brun0Nasc/Frequencias-Backend/config/database"
	domain "github.com/Brun0Nasc/Frequencias-Backend/domain/frequencias/model"
	"github.com/Brun0Nasc/Frequencias-Backend/infra/frequencias"
	"github.com/Brun0Nasc/Frequencias-Backend/utils"
)

// GerarListaFrequencia é responsável por criar uma lista de frequencia em um determinado período
func GerarFrequencia() (erro error) {
	if horarioValido := validarHorarioExecucao(0, 1); horarioValido {
		var (
			agora           = time.Now()
			proximaExecucao = time.Date(agora.Year(), agora.Month(), agora.Day()+1, 0, 0, 0, 0, time.Local)
		)

		gerarFrequencia()
		//* Aguardar a próxima execução da rotina de acordo com o informado na variável proximaExecucao
		AguardarHorarioExecucao(agora, proximaExecucao)
		//* Inicializando novamente a rotina após o período de pausa
		go GerarFrequencia()
	}

	return
}

// gerarFrequencia é responsável por toda a regra de negócio para gerar uma frequencia
func gerarFrequencia() (erro error) {
	db := database.Conectar()
	defer db.Close()
	var (
		frequenciasRepo  = frequencias.NovoRepo(db)
		params           = utils.RequestParams{Filters: make(map[string][]string, 0)}
		listaFrequencias *domain.ListaFrequencias
	)

	params.AddFilter("existe_frequencia_hoje", "true")
	// Buscando frequencia criada no dia atual
	if listaFrequencias, erro = frequenciasRepo.ListarFrequencias(&params); erro != nil {
		return
	}

	// Verificando se existe frequencia gerada para o dia atual
	if listaFrequencias != nil && len(listaFrequencias.Dados) > 0 {
		return errors.New("Já foi gerada uma frequência para o dia atual.")
	}

	// Criando uma nova frequência
	if erro = frequenciasRepo.NovaFrequencia(); erro != nil {
		return
	}

	fmt.Println("Executando")
	return
}

// validarHorarioExecucao é responsável por validar se o script está em um horario de execução válido
func validarHorarioExecucao(inicio, fim int) bool {
	var agora = time.Now()

	if inicio <= fim {
		return (agora.Hour() >= inicio && agora.Hour() <= fim)
	}

	return (agora.Hour() >= inicio || agora.Hour() <= fim)
}

// AguardarHorarioExecucao é responsável por pausar a execução da rotina até o horário informado
func AguardarHorarioExecucao(agora, proximaExecucao time.Time) {
	time.Sleep(proximaExecucao.Sub(agora))
}
