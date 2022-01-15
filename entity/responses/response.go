package responses

type Response struct {
	Error   string      `json:"-"`
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Meta    Meta        `json:"meta"`
	Data    interface{} `json:"data"`
}

type Meta struct {
	Cache    bool  `json:"cache"`
	UnixTime int64 `json:"time"`
}
