package auth

import "time"

type AuthorizedUserInfo struct {
	ID       string    `json:"id"`
	AuthType string    `json:"authType"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
	Gender   string    `json:"gender"`
	PhotoURL string    `json:"picture"`
}
