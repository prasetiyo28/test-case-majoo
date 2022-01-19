package middlewares

import (
	"fmt"
	"test-case-majoo/entity/responses"
	"test-case-majoo/pkg/errors"
	"test-case-majoo/pkg/response"
	"test-case-majoo/pkg/utils"

	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type authHeader struct {
	IDToken string `header:"Authorization"`
}

type AccessToken struct {
	Email string `json:"email"`
}

// AuthUser extracts a user from the Authorization header
// which is of the form "Bearer token"
// It sets the user to the context if the user exists
func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}

		// bind Authorization Header to h and check for validation errors
		if err := c.ShouldBindHeader(&h); err != nil {
			if _, ok := err.(validator.ValidationErrors); ok {
				// we used this type in bind_data to extract desired fields from errs
				// you might consider extracting it

				result := responses.Response{
					Code:    http.StatusBadRequest,
					Message: errors.ErrParamInvalid.Error(),
				}
				response.Response(c, &result)
				c.Abort()
				return
			}

			result := responses.Response{
				Code:    http.StatusInternalServerError,
				Message: errors.ErrInternResp.Error(),
			}
			response.Response(c, &result)
			c.Abort()
			return
		}

		idTokenHeader := strings.Split(h.IDToken, "Bearer ")
		if len(idTokenHeader) < 2 {
			result := responses.Response{
				Code:    http.StatusUnauthorized,
				Message: errors.ErrAuthorizationBearer.Error(),
			}
			response.Response(c, &result)
			c.Abort()
			return
		}

		// validate ID token here
		token, err := utils.ValidateToken(idTokenHeader[1])
		if err != nil {
			result := responses.Response{
				Code:    http.StatusUnauthorized,
				Message: errors.ErrInvalidToken.Error(),
			}
			response.Response(c, &result)
			c.Abort()
			return

		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !token.Valid || !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": nil,
			})
			c.Abort()
			return
		}

		if claims["data"] == nil {
			result := responses.Response{
				Code:    http.StatusUnauthorized,
				Message: errors.ErrInvalidToken.Error(),
			}
			response.Response(c, &result)
			c.Abort()
			return
		}

		params := claims["data"]
		fmt.Println("claims", claims["data"])
		m, ok := params.(map[string]interface{})
		if !ok {
			fmt.Println(ok)
		}
		if m["id"] == nil || m["id"] == "" {
			result := responses.Response{
				Code:    http.StatusUnauthorized,
				Message: errors.ErrInvalidToken.Error(),
			}
			response.Response(c, &result)
			c.Abort()
			return
		}

		c.Set("userGuid", m["id"])
		// c.Set("email", m["email"]) //note: tobedeleted
		c.Next()
	}
}
