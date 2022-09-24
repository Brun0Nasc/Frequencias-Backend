package frequencias

import (
	"fmt"

	"github.com/Brun0Nasc/Frequencias-Backend/config/database"
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/frequencias/model"
	"github.com/Brun0Nasc/Frequencias-Backend/infra/frequencias"
)

func NovaFrequencia(req *modelApresentacao.Frequencia) (err error) {
	db := database.Conectar()
	defer db.Close()
	frequenciasRepo := frequencias.NovoRepo(db)

	err = frequenciasRepo.NovaFrequencia(req)
	if err != nil {
		return fmt.Errorf("frequencia not added \nerr:" + err.Error())
	}
	return
}

func ListarFrequenciasUsuario(idUser *int64) (res *modelApresentacao.ListaFrequencias, err error) {
	db := database.Conectar()
	defer db.Close()
	frequenciasRepo := frequencias.NovoRepo(db)

	res, err = frequenciasRepo.ListarFrequenciasUsuario(idUser)
	if err != nil {
		return nil, fmt.Errorf("frequencia not list \nerr:" + err.Error())
	}
	return
}
