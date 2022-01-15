package v1

import (
	HANDLER "test-case-majoo/api/handler/auths/login"
	MIDDLEWARES "test-case-majoo/api/middleware"

	"test-case-majoo/usecase/auth"

	"github.com/gin-gonic/gin"
)

func RouteLogin(r *gin.RouterGroup, at *auth.Service) {

	login := HANDLER.LoginControllerHandler(at)

	r.POST("/login", login.Login)
	r.POST("/report", MIDDLEWARES.AuthUser(), login.MonthlyReport)
}
