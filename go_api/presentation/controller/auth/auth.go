package auth

import (
	"encoding/json"
	"net/http"

	"github.com/nozomi-iida/nozo_blog/presentation"
	"github.com/nozomi-iida/nozo_blog/service"
)

type AuthController struct {
	as *service.AuthService
}

type AuthRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
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
	var authRequest AuthRequest
	json.Unmarshal(body, &authRequest)

	ur, err := ac.as.SignUp(authRequest.Username, authRequest.Password)
	if err != nil {
		presentation.ErrorHandler(w, err)
		return	
	}

	output, _ := json.MarshalIndent(ur, "", "\t")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (ac *AuthController) SignInRequest(w http.ResponseWriter, r *http.Request)  {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var authRequest AuthRequest
	json.Unmarshal(body, &authRequest)
	if !presentation.IsValid(w, authRequest) {
		return
	}
	
	ur, err := ac.as.SignIn(authRequest.Username, authRequest.Password)
	if err != nil {
		presentation.ErrorHandler(w, err)
		return
	}

	output, _ := json.MarshalIndent(ur, "", "\t")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
