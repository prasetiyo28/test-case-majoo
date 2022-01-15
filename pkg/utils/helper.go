package utils

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	ENTITY_RESPONSE "test-case-majoo/entity/responses"
	"test-case-majoo/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/id"
	UNIV_TRANS "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ID_TRANS "github.com/go-playground/validator/v10/translations/id"
)

const (
	LETTER_BYTES    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LETTER_IDX_BITS = 6                      // 6 bits to represent a letter index
	LETTER_IDX_MASK = 1<<LETTER_IDX_BITS - 1 // All 1-bits, as many as letterIdxBits
	LETTER_IDX_MAX  = 63 / LETTER_IDX_BITS   // # of letter indices fitting in 63 bits

	BIND_TYPE_JSON  = 1
	BIND_TYPE_PARAM = 2
)

type Repository interface {
	ComparePassword(old string, exist string) bool
	CheckPassword(password string) string
	ValidateRequest(bindType int, funcName string, ctx *gin.Context, input interface{}) *ENTITY_RESPONSE.Response
	GeneratePassword(password string) (string, *ENTITY_RESPONSE.Response)
	GenerateAccessToken(email, guid string, expDate int64) (string, *ENTITY_RESPONSE.Response)
	CreateCookie(userId string, userName string, firstName string, lastName string, guid string, createdDate string, jwtToken string, JwtRefreshToken string, c *gin.Context) *ENTITY_RESPONSE.Response
}

func ValidateRequest(bindType int, funcName string, ctx *gin.Context, input interface{}) *ENTITY_RESPONSE.Response {
	//check request body
	if bindType == BIND_TYPE_JSON {
		if errBind := ctx.ShouldBindJSON(&input); errBind != nil {
			return &ENTITY_RESPONSE.Response{
				Code:    http.StatusBadRequest,
				Message: errors.ErrFormatRequestBody.Error(),
			}
		}
	} else if bindType == BIND_TYPE_PARAM {
		if errBind := ctx.ShouldBindQuery(input); errBind != nil {
			return &ENTITY_RESPONSE.Response{
				Code:    http.StatusBadRequest,
				Message: errors.ErrFormatRequestBody.Error(),
			}
		}
	} else {
		return &ENTITY_RESPONSE.Response{
			Code:    http.StatusBadRequest,
			Message: errors.ErrFormatRequestBody.Error(),
		}
	}

	//validate request body
	validate := validator.New()
	uni := UNIV_TRANS.New(id.New())
	trans, _ := uni.GetTranslator("id")

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("label"), ",", 2)[0]
		if name == "" {
			return ""
		}

		return name
	})

	//Verifier registration translator
	errTranslation := ID_TRANS.RegisterDefaultTranslations(validate, trans)
	if errTranslation != nil {
		return &ENTITY_RESPONSE.Response{
			Code:    http.StatusBadRequest,
			Message: errTranslation.Error(),
		}
	}

	errTranslation = validate.Struct(input)
	msgError := ""

	if errTranslation != nil {
		for _, e := range errTranslation.(validator.ValidationErrors) {
			translatedErr := fmt.Errorf(e.Translate(trans))
			msgError = msgError + translatedErr.Error() + ". "
		}

		return &ENTITY_RESPONSE.Response{
			Code:    http.StatusBadRequest,
			Message: msgError,
		}
	}

	return nil
}
