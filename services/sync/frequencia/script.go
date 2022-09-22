package frequencia

import (
	"fmt"
	"time"
)

const (
	// tempoExecucao é a constante responsável por definir o intervalo de tempo de execução da rotina
	tempoExecucao = time.Second * 10
	// tipoRotinaMultipla é a constante que identifica a rotina como tipo de execução múltipla
	tipoRotinaMultipla string = "[rotina_de_multipla_execucao]"
	// tipoRotinaUnica é a constante que identifica a rotina como tipo de execução única
	tipoRotinaUnica string = "[rotina_de_execucao_unica]"
)

// GerarListaFrequencia é responsável por criar o tick de execução da rotina de gerar frequências
func GerarListaFrequencia() (erro error) {
	// * Com essa implementação a execução da goroutine é realizada em loop
	// * Serve para scripts que precisam ficar atualizando dados em tempo real
	var ticker = time.NewTicker(tempoExecucao)
	gerarListaFrequencia(tipoRotinaMultipla)

	// Inicializando loop de acordo com o tempo de execução especificado
	for range ticker.C {
		gerarListaFrequencia(tipoRotinaMultipla)
	}

	//! Nesse caso como não está sendo realizada nenhuma execução real
	//! a variável "erro" não está sendo tratada em momento algum
	//! mas assim que for implementada a regra de negócio é de extrema importância
	//! que os erros sejam tratados da melhor forma possível e retornados para a main

	//! O erro não deve ser estourado em caso de falha, pois o mesmo interromperia a execução da rotina principal (a função main)
	//! Tendo isso em vista o erro deve apenas ser tratado e loggado
	return
}

// GerarListaFrequencia2 é responsável por criar uma lista de frequencia em um determinado período
func GerarListaFrequencia2() (erro error) {
	// * Com essa implementação a execução da goroutine é realizada
	// * apenas uma vez dentro do horário especificado [23h - 02h]
	if horarioValido := validarHorarioExecucao(8,10); horarioValido {
		gerarListaFrequencia(tipoRotinaUnica)
	}

	return
}

// gerarListaFrequencia é responsável por toda a regra de negócio para gerar uma lista de frequencia
func gerarListaFrequencia(tipoRotina string) (erro error) {
	fmt.Println("Gerando lista de frequência...", tipoRotina)
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
