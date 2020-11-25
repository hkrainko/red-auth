package auth

type AuthorizedUserInfo struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	PhotoUrl string `json:"picture"`
}
