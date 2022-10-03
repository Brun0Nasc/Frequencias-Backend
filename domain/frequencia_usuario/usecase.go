package frequencia_usuario

import (
	"fmt"

	"github.com/Brun0Nasc/Frequencias-Backend/config/database"
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/frequencia_usuario/model"
	"github.com/Brun0Nasc/Frequencias-Backend/infra/frequencia_usuario"
	"github.com/Brun0Nasc/Frequencias-Backend/infra/frequencias"
	"github.com/Brun0Nasc/Frequencias-Backend/utils"
)

func NovaFrequenciaUsuario(req *modelApresentacao.ReqFrequenciaUsuario) (err error) {
	db := database.Conectar()
	defer db.Close()
	frequenciaUsuarioRepo := frequencia_usuario.NovoRepo(db)
	frequenciasRepo := frequencias.NovoRepo(db)

	// * Pegando id da data de frequência mais atual
	idFreq, err := frequenciasRepo.PegarFrequenciaMaisRecente()
	if err != nil {
		return
	}

	req.FrequenciaID = idFreq

	err = frequenciaUsuarioRepo.NovaFrequenciaUsuario(req)
	if err != nil {
		return fmt.Errorf("frequencia de usuario nao registrada // err:" + err.Error())
	}
	return
}

func ListaFrequenciasData(params *utils.RequestParams) (res *modelApresentacao.ListaUsuarioFrequencia, err error) {
	db := database.Conectar()
	defer db.Close()
	frequenciaUsuarioRepo := frequencia_usuario.NovoRepo(db)

	res, err = frequenciaUsuarioRepo.ListaFrequenciasData(params)
	if err != nil {
		return nil, fmt.Errorf("nao foi possível listar frequências da data especificada // err:" + err.Error())
	}
	return
}
