package auths

import "time"

type OauthRequest struct {
	ClientID     string `form:"client_id"  validate:"required" `
	ClientSecret string `form:"client_secret" validate:"required" `
	GrantType    string `form:"grant_type" validate:"required,oneof=authorization_code refresh_token" `
	Code         string `form:"code"`
	RedirectUri  string `form:"redirect_uri"`
	RefreshToken string `form:"refresh_token"`
}

type UserSWG struct {
	GuID         string `json:"guid"`
	Scope        string `form:"scope" json:"scope"`
	RedirectUri  string `form:"redirect_uri" json:"redirect_uri"`
	ResponseType string `form:"response_type" json:"response_type"`
	ClientID     string `form:"client_id" json:"client_id"`
	State        string `form:"state" json:"state"`
	Expired      int    `json:"expired"`
}

type SSOUserInfoRequest struct {
	AccessToken string `form:"access_token"`
}

type OauthClient struct {
	GuID          string    `json:"guid" firestore:"userGuid"`
	OauthClientId string    `json:"oauthClientId" firestore:"oauthClientId"`
	Status        int       `json:"status" firestore:"status"`
	Token         string    `json:"token" firestore:"token"`
	Name          string    `json:"name" firestore:"name"`
	Scope         string    `json:"scope" firestore:"scope"`
	Description   string    `json:"description" firestore:"description"`
	CreatedDate   time.Time `json:"createdDate" firestore:"createdDate"`
	UpdateDate    time.Time `json:"updatedDate" firestore:"updatedDate"`
	AccessToken   string    `json:"access_token"`
	RefreshToken  string    `json:"refresh_token"`
}

type OauthClients []OauthClient

type ConfirmLinking struct {
	ConfirmKeys string `json:"confirmKey" validate:"required" `
}

type AuthLinkingClient struct {
	ClientSecret string `json:"client_secret" firestore:"clientSecret"`
	ClientID     string `json:"client_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Scope        string `json:"scope"`
}
