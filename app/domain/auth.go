package domain


type AuthType string


type AuthRepository interface {
	Auth(requester Requester) (Resident, error)
	GetAuthUrl(authType AuthType) string
}

type AuthUseCase interface {
	Auth(requester Requester) (Resident, error)
	GetAuthUrl(authType AuthType) string
}