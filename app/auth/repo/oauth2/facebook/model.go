package facebook

import (
	"red-auth/domain/auth"
	"time"
)

type AuthorizedUserInfo struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Email    string          `json:"email"`
	Birthday string          `json:"birthday"`
	Gender   string          `json:"gender"`
	Picture  facebookPicture `json:"picture"`
}

type facebookPicture struct {
	Data facebookPictureData `json:"data"`
}

type facebookPictureData struct {
	Height       float64 `json:"height"`
	Width        float64 `json:"width"`
	IsSilhouette bool    `json:"is_silhouette"`
	URL          string  `json:"url"`
}

func (g *AuthorizedUserInfo) toUserInfo() *auth.UserInfo {

	t, _ := time.Parse("01/02/2006", g.Birthday)
	gender := ""
	switch g.Gender {
	case "male":
		gender = "M"
	case "female":
		gender = "F"
	default:
		gender = ""
	}

	return &auth.UserInfo{
		ID:       g.ID,
		AuthType: "Facebook",
		Name:     g.Name,
		Email:    g.Email,
		Birthday: t,
		Gender:   gender,
		PhotoURL: g.Picture.Data.URL,
	}
}
