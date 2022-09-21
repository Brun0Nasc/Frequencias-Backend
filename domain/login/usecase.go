package login

import (
	"github.com/Brun0Nasc/Frequencias-Backend/config/database"
	"github.com/Brun0Nasc/Frequencias-Backend/config/services"
	"github.com/Brun0Nasc/Frequencias-Backend/infra/login"
	modelApresentacao "github.com/Brun0Nasc/Frequencias-Backend/domain/login/model"
)

func LoginUsuario(req *modelApresentacao.Login) (*modelApresentacao.Login, error) {
	db := database.Conectar()
	defer db.Close()

	req.Senha = services.SHAR256Encoder(req.Senha)

	loginRepo := login.NovoRepo(db)
	return loginRepo.LoginUsuario(req) 
}