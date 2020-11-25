package auth

type CallBack struct {
	AuthType Type
	State    string
	Code     string
}