package login

type ReqLogin struct {
	Email *string `json:"email"`
	Senha string `json:"senha"`
}
