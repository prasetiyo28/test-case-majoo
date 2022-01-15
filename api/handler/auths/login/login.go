package login

import "test-case-majoo/usecase/auth"

type LoginController struct {
	at *auth.Service
}

func LoginControllerHandler(at *auth.Service) *LoginController {
	handler := &LoginController{
		at: at,
	}
	return handler
}
