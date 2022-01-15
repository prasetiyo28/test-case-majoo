package utils

import (
	"fmt"
	ENTITY_RESPONSE "test-case-majoo/entity/responses"
	"test-case-majoo/pkg/errors"

	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JwtTokenGenerate is a function used to generate access token and refresh token
// Param : email, guid, exp
// Response : token, refresh token, err
func JwtTokenGenerate(username string, guid string, exp int64, rtExp int64) (token string, rt string, err *ENTITY_RESPONSE.Response) {
	expToken := time.Now().Add(time.Minute * 15).Unix() // will uncomment after UAT finish
	if exp != 0 {
		expToken = exp
	}

	rtExpDate := time.Now().Add(time.Hour * 24 * 365).Unix()
	if rtExp != 0 {
		rtExpDate = rtExp
	}

	// create interface and add to jwt mapClaims
	data := map[string]interface{}{
		"id":       guid,
		"username": username,
	}
	signBytes, _ := ioutil.ReadFile(os.Getenv("PRIVATE_KEY"))
	signKey, _ := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	atClaims := &jwt.MapClaims{
		"iss":  os.Getenv("ALLOWED_ORIGIN"),
		"iat":  time.Now().Unix(),
		"exp":  expToken,
		"data": data,
	}

	refreshToken := jwt.New(jwt.GetSigningMethod("RS256"))
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["username"] = username
	rtClaims["id"] = guid
	rtClaims["exp"] = rtExpDate
	rt, errRt := refreshToken.SignedString(signKey)
	if errRt != nil {
		return "", "", &ENTITY_RESPONSE.Response{
			Code:    http.StatusInternalServerError,
			Message: errors.ErrClaimRefreshToken.Error(),
		}
	}

	// add expDate refresh token to access token
	data["rt"] = rtClaims["exp"]
	at := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), atClaims)
	token, errAt := at.SignedString(signKey)
	if errAt != nil {
		return "", "", &ENTITY_RESPONSE.Response{
			Code:    http.StatusInternalServerError,
			Message: errors.ErrClaimAccessToken.Error(),
		}
	}

	return token, rt, nil
}

// ValidateToken is a function used to validate encoded token
// Param : encodedToken
// Response : jwt.Token, ENTITY_RESPONSE.Response
func ValidateToken(encodedToken string) (*jwt.Token, *ENTITY_RESPONSE.Response) {
	verifyBytes, err := ioutil.ReadFile(os.Getenv("PUBLIC_KEY"))
	if err != nil {
		return nil, &ENTITY_RESPONSE.Response{
			Code:    http.StatusInternalServerError,
			Message: errors.ErrPublicKey.Error(),
		}
	}
	verifyKey, _ := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

	if err == nil {
		return token, nil
	}
	return token, &ENTITY_RESPONSE.Response{
		Code:    http.StatusUnauthorized,
		Message: errors.ErrClaimToken.Error(),
	}
}

// GetExpDate is a function used to get expired date token
// Param : encodedToken
// Response : expiredDateToString
func GetExpDate(encodedToken string) string {
	decodeToken, _ := ValidateToken(encodedToken)
	claims := decodeToken.Claims.(jwt.MapClaims)
	expDate := claims["exp"].(float64)
	expDateToString := strconv.FormatFloat(expDate, 'f', -1, 64)

	return expDateToString

}

// GetEmail is a function used to get email in token
// Param : encodedToken
// Response : email
func GetUsername(encodedToken string, tokenType int) string {
	decodeToken, _ := ValidateToken(encodedToken)
	claims := decodeToken.Claims.(jwt.MapClaims)
	if tokenType == TOKEN_TYPE_REFRESH {
		return claims["username"].(string)
	}
	data := claims["data"].(map[string]interface{})
	return data["username"].(string)
}

// GenerateAccessToken is a function used to generate access token
// Param : email, guid, expDate
// Response : token
func GenerateAccessToken(username, guid string, expDate int64, rtExpDate int64) (string, *ENTITY_RESPONSE.Response) {
	token, _, errToken := JwtTokenGenerate(string(username), guid, expDate, rtExpDate)
	if errToken != nil {
		return "", &ENTITY_RESPONSE.Response{
			Code:    http.StatusInternalServerError,
			Message: errors.ErrClaimToken.Error(),
		}
	}
	return token, nil
}

// GetExpDate is a function used to get expired date refresh token
// Param : encodedToken
// Response : expiredDate Refresh Token
func ExpDateRefreshToken(encodedToken string) string { // <-- encodedToken = bearer
	decode, _ := ValidateToken(encodedToken)
	claim := decode.Claims.(jwt.MapClaims)
	dataClaim := claim["data"].(map[string]interface{})
	exp := dataClaim["rt"].(float64)
	expDate := fmt.Sprintf("%.0f", exp)
	return expDate
}
