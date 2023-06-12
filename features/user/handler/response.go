package handler

type loginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
