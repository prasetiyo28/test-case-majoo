package auths

type Register struct {
	AccessToken   string `json:"accessToken"`
	RefreshToken  string `json:"refreshToken"`
	DocReferrer   string `json:"docReferrer,omitempty"`
	PasswordCheck string `json:"passwordCheck"`
}

type RegisterWP struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	DocReferrer  string `json:"docReferrer,omitempty"`
}

type RegisterVIP struct {
	FirstName string `json:"firstName"`
	Guid      string `json:"userGuid"`
	Email     string `json:"email"`
}

type RegisteredOn struct {
	RegisteredOn []string `json:"registeredOn"`
}
