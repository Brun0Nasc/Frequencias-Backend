package login

import (
	"database/sql"
	"fmt"

	"github.com/Brun0Nasc/Frequencias-Backend/config/database"
	"github.com/Brun0Nasc/Frequencias-Backend/config/services"
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/usuarios/model"
	"github.com/Brun0Nasc/Frequencias-Backend/infra/login"
)

func LoginUsuario(req *modelApresentacao.Usuario) (token string, err error) {
	db := database.Conectar()
	defer db.Close()

	req.Senha = services.SHAR256Encoder(req.Senha)

	loginRepo := login.NovoRepo(db)
	res, err := loginRepo.LoginUsuario(req)
	
	if res == nil && err == nil {
		return "", fmt.Errorf("senha incorreta")
	}

	if err != nil {
		if err == sql.ErrNoRows{
			return "email não encontrado", nil
		}
		return "", fmt.Errorf("não foi possível logar, erro: " + err.Error() )
	}

	// Checks if there is an error in this request
	token, err = services.NewJWTService().GenerateToken(res.ID, res.Tipo, &res.Senha, res.Nome)

	return
}