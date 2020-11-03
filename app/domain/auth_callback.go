package domain

type AuthCallBack struct {
	AuthType AuthType
	State string
	Code string
}