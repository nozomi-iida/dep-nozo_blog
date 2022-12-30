package presentation

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/nozomi-iida/nozo_blog/domain/user"
	"github.com/nozomi-iida/nozo_blog/service"
)

type ErrorPresentation struct {
	Error errMessage `json:"error"`
}

type errMessage struct {
	Message string `json:"message"`
	Code int `json:"code"`
}

func NewErrorPresentation(err error) ErrorPresentation  {
	switch err {
	case user.ErrUserNotFound:
		return ErrorPresentation{
			Error: errMessage{
				Message: user.ErrUserNotFound.Error(),
				Code: http.StatusNotFound,
			},
		}
		case service.ErrUnMatchPassword:
			return ErrorPresentation{
				Error: errMessage{
					Message: user.ErrUserNotFound.Error(),
					Code: http.StatusUnauthorized,
				},
			} 
	default:
		se := errors.New("server error")
		return ErrorPresentation{
			Error: errMessage{
				Message: se.Error(),
				Code: http.StatusInternalServerError,
			},
		}
	}
}

func ErrorHandler(w http.ResponseWriter, err error)  {
	ep := NewErrorPresentation(err)
	output, _ := json.MarshalIndent(ep, "", "\t")
	w.WriteHeader(ep.Error.Code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
