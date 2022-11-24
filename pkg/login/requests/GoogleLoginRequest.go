package login

type GoogleLoginRequest struct {
	Token string `validate:"required" json:"token"`
}
