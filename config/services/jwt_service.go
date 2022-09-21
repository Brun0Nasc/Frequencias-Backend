package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	secretKey string
	issure    string
}

func NewJWTService() *jwtService {
	return &jwtService{
		secretKey: "secret-key",
		issure:    "Gerenciador-Projetos",
	}

}

type Claim struct {
	Sum uint `json:"sum"`
	Sub uint `json:"tipo"`
	Set string `json:"nome"`
	jwt.StandardClaims
}

// Token generation function
func (s *jwtService) GenerateToken(id uint, tipo uint, nome string) (string, error) {
	claim := &Claim{
		id,
		tipo,
		nome,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

// Function for validate Token
func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(s.secretKey), nil
	})
	return err == nil
}
