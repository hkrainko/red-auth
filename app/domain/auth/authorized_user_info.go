package auth

type AuthorizedUserInfo struct {
	ID       string `json:"id"`
	AuthType string `json:"authType"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Gender   string `json:"gender"`
	PhotoUrl string `json:"picture"`
}
