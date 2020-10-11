package domain


type Requester struct {
	AuthType string
	AuthURL string
}

type RequesterRepository interface {
	GetAuthURL(authType AuthType) string
}