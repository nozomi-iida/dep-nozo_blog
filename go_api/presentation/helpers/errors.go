package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/nozomi-iida/nozo_blog/domain/user"
	"github.com/nozomi-iida/nozo_blog/service"
)

var (
	ErrBadRequest = errors.New("Bad Request")
	ErrStatusMethodNotAllowed = errors.New(http.StatusText(http.StatusMethodNotAllowed))
	ErrUnauthorized = errors.New("Unauthorized users")
)

type ErrorPresentation struct {
	Error errMessage `json:"error"`
}

type errMessage struct {
	Message string `json:"message"`
	Code int `json:"code"`
	Type string `json:"type"`
}

func newErrMsg(message string, code int) ErrorPresentation  {
	return ErrorPresentation{
		Error: errMessage{
			Message: message,
			Code: code,
			Type: http.StatusText(code),
		},
	}
}

func NewErrorPresentation(err error) ErrorPresentation  {
	switch err {
	case user.ErrUserNotFound:
		return newErrMsg(user.ErrUserNotFound.Error(), http.StatusNotFound)
	case service.ErrUnMatchPassword:
		return newErrMsg(service.ErrUnMatchPassword.Error(), http.StatusUnauthorized)
	case ErrStatusMethodNotAllowed:
		return newErrMsg("method not allowed", http.StatusMethodNotAllowed)
	case ErrUnauthorized:
		return newErrMsg(ErrUnauthorized.Error(), http.StatusUnauthorized)
	default:
		se := errors.New("server error")
		return newErrMsg(se.Error(), http.StatusInternalServerError)
	}
}

func ErrorHandler(w http.ResponseWriter, err error)  {
	ep := NewErrorPresentation(err)
	output, _ := json.MarshalIndent(ep, "", "\t")
	w.WriteHeader(ep.Error.Code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

// 少し冗長な気がする
func IsValid(w http.ResponseWriter, s interface{}) bool {
	var vl = validator.New()
	err := vl.Struct(s)	
	if err == nil {
		return true 
	} else {
		msg := ""
		for _, err := range err.(validator.ValidationErrors) {
			msg = fmt.Sprintf("%v is %v", err.Field(), err.Tag())
		}
		ep := newErrMsg(msg, http.StatusBadRequest)
		output, _ := json.MarshalIndent(ep, "", "\t")
		w.WriteHeader(ep.Error.Code)	
		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
		return false
	}
}

func Validate(w http.ResponseWriter, s interface{}) {
	var vl = validator.New()
	err := vl.Struct(s)	
	if err != nil {
		msg := ""
		for _, err := range err.(validator.ValidationErrors) {
			msg = fmt.Sprintf("%v is %v", err.Field(), err.Tag())
		}
		ep := newErrMsg(msg, http.StatusBadRequest)
		output, _ := json.MarshalIndent(ep, "", "\t")
		w.WriteHeader(ep.Error.Code)	
		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	}
}
