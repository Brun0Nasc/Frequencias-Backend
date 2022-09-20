package frequencias

import (
	"fmt"

	"github.com/Brun0Nasc/Frequencias-Backend/config/database"
	"github.com/Brun0Nasc/Frequencias-Backend/infra/frequencias"
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/frequencias/model"
)

func NovaFrequencia(req *modelApresentacao.ReqFrequencia) (err error) {
	db := database.Conectar()
	defer db.Close()
	frequenciasRepo := frequencias.NovoRepo(db)

	err = frequenciasRepo.NovaFrequencia(req)
	if err != nil{
		return fmt.Errorf("frequencia not added \nerr:"+err.Error())
	}
	return
}