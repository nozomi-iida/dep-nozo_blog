package presentation

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/nozomi-iida/nozo_blog/domain/user"
	"github.com/nozomi-iida/nozo_blog/service"
)

type ErrorPresentation struct {
	Error   error
	Message string
	Code    int
}

func ErrorHandler(err error) ErrorPresentation  {
	fmt.Printf("err: %v", err)

	switch err {
	case user.ErrUserNotFound:
		return ErrorPresentation{
			Error: user.ErrUserNotFound,
			Message: user.ErrUserNotFound.Error(),
			Code: http.StatusNotFound,
		}
		case service.ErrUnMatchPassword:
			return ErrorPresentation{
				Error: service.ErrUnMatchPassword,
				Message: service.ErrUnMatchPassword.Error(),
				Code: http.StatusUnauthorized,
			} 
	default:
		se := errors.New("server error")
		return ErrorPresentation{
			Error: se,
			Message: se.Error(),
			Code: http.StatusInternalServerError,
		}
	}
}
