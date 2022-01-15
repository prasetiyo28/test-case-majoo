package auths

type Login struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Temp struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	DocReferrer  string `json:"docReferrer,omitempty"`
	DeviceKeyId  string `json:"deviceKeyId"`
	IsSocial     bool   `json:"isSocial"`
}
type Logout struct {
	Success bool   `json:"success" example:"true"`
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"login sukses"`
}
