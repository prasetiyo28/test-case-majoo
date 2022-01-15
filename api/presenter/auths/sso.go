package auths

type GoogleSubscription struct {
	Success    bool   `json:"success" example:"true"`
	Code       int    `json:"code" example:"200"`
	Message    string `json:"message" example:"berhasil menampilkan data user"`
	Subscribed bool   `json:"subscribed"`
}

type SSOLogin struct {
	TokenType    string `json:"tokenType"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn,omitempty"`
}

type AuthorizationCode struct {
	AuthorizationCode string `json:"authorization_code"`
}

type OauthLinking struct {
	AccessToken  string `json:"accessToken"`
	IdToken      string `json:"id_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refreshToken"`
	Key          string `json:"key"`
	RedirectUri  string `json:"redirect_uri"`
	Linked       bool   `json:"linked"`
}

type SSOUserInfoResponse struct {
	GuID      string `json:"guid"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
