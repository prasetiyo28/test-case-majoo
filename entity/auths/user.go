package auths

import "time"

type User struct {
	ID        string    `json:"guid"`
	UserName  string    `json:"user_name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int       `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int       `json:"updated_by"`
}

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
