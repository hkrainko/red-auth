package call_back

import (
	"red-auth/domain/auth"
	"time"
)

type SuccessResponse struct {
	ID       string    `json:"id"`
	AuthType string    `json:"authType"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
	Gender   string    `json:"gender"`
	PhotoURL string    `json:"picture"`
}

func NewSuccessResponse(info *auth.UserInfo) *SuccessResponse {
	return &SuccessResponse{
		ID:       info.ID,
		AuthType: info.AuthType,
		Email:    info.Email,
		Birthday: info.Birthday,
		Gender:   info.Gender,
		PhotoURL: info.PhotoURL,
	}
}