package auth

type AuthorizedUserInfo struct {
	Token    string    `json:"token"`
	UserInfo
}
