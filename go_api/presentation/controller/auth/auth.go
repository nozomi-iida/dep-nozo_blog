package auth

import (
	"encoding/json"
	"net/http"

	"github.com/nozomi-iida/nozo_blog/service"
)

type AuthController struct {
	as *service.AuthService
}

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAuthController(fileString string) (AuthController, error)  {
	as, err := service.NewAuthService(
		service.WithSqliteUserRepository(fileString),
	)
	if err != nil {
		return AuthController{}, err
	}
	return AuthController{as: as}, nil	
}

func (ac *AuthController) SignUpRequest(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var signUpRequest SignUpRequest
	json.Unmarshal(body, &signUpRequest)

	ur, err := ac.as.SignUp(signUpRequest.Username, signUpRequest.Password)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	output, _ := json.MarshalIndent(ur, "", "\t\t")

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
