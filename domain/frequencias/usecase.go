package frequencias

import (
	"fmt"

	"github.com/Brun0Nasc/Frequencias-Backend/config/database"
	"github.com/Brun0Nasc/Frequencias-Backend/infra/frequencias"
)

func NovaFrequencia() (err error) {
	db := database.Conectar()
	defer db.Close()
	frequenciasRepo := frequencias.NovoRepo(db)

	err = frequenciasRepo.NovaFrequencia()
	if err != nil {
		return fmt.Errorf("frequencia não criada \nerr:" + err.Error())
	}
	return
}

func PegarFrequenciaMaisRecente() (res *int, err error) {
	db := database.Conectar()
	defer db.Close()
	frequenciasRepo := frequencias.NovoRepo(db)

	res, err = frequenciasRepo.PegarFrequenciaMaisRecente()

	if err != nil {
		return nil, fmt.Errorf("frequencia não listada \nerr:" + err.Error())
	}

	return
}
