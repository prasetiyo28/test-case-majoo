package login

import (
	"test-case-majoo/api/presenter/auths"
	ENTITY_AUTH "test-case-majoo/entity/auths"
	"test-case-majoo/entity/responses"
	"test-case-majoo/pkg/errors"
	"test-case-majoo/pkg/response"
	"test-case-majoo/pkg/success"
	"test-case-majoo/pkg/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *LoginController) Login(context *gin.Context) {
	var input ENTITY_AUTH.LoginRequest

	err := utils.ValidateRequest(utils.BIND_TYPE_JSON, "Login", context, &input)
	if err != nil {
		result := responses.Response{
			Code:    err.Code,
			Message: err.Message,
		}
		response.Response(context, &result)
		return
	}

	user, errGetUser := handler.at.GetUserByID(input.UserName)
	if errGetUser != nil {
		result := responses.Response{
			Code:    errGetUser.Code,
			Message: errGetUser.Message,
		}
		response.Response(context, &result)
		return
	}

	hashPassword := utils.Hashing(input.Password)
	if hashPassword != user.Password {
		result := responses.Response{
			Code:    http.StatusUnauthorized,
			Message: errors.ErrLogin.Error(),
		}
		response.Response(context, &result)
		return
	}

	token, rt, errToken := utils.JwtTokenGenerate(user.UserName, user.ID, 0, 0)
	if errToken != nil {
		if errToken != nil {
			result := responses.Response{
				Code:    errToken.Code,
				Message: errToken.Message,
			}
			response.Response(context, &result)
			return
		}
	}

	result := responses.Response{
		Code:    http.StatusOK,
		Message: success.SuccessLogin,
		Data: auths.Login{
			AccessToken:  token,
			RefreshToken: rt,
		},
	}
	response.Response(context, &result)
}
