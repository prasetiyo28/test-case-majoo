package auths

type LoginApple struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	DocReferrer  string `json:"docReferrer,omitempty"`
	IsSocial     bool   `json:"isSocial"`
	IsPassEmpty  bool   `json:"isPassEmpty"`
	UserCreated  bool   `json:"userCreated"`
	DeviceKeyId  string `json:"deviceKeyId"`
}

type LoginGoogle struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	DocReferrer  string `json:"docReferrer,omitempty"`
	IsSocial     bool   `json:"isSocial"`
	IsPassEmpty  bool   `json:"isPassEmpty"`
	UserCreated  bool   `json:"userCreated"`
	DeviceKeyId  string `json:"deviceKeyId"`
}
type LoginFacebook struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	DocReferrer  string `json:"docReferrer,omitempty"`
	IsSocial     bool   `json:"isSocial"`
	IsPassEmpty  bool   `json:"isPassEmpty"`
	UserCreated  bool   `json:"userCreated"`
	DeviceKeyId  string `json:"deviceKeyId"`
}
