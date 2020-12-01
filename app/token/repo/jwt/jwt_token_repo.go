package jwt

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"red-auth/app/domain/token"
	"time"
)

type Claims struct {
	Account string      `json:"account"`
	Role string         `json:"role"`
	jwt.StandardClaims
}

var jwtSecret = []byte("secret")

type jwtTokenRepo struct {

}

func NewJWTTokenRepo() token.Repo {
	return &jwtTokenRepo{

	}
}

func (j jwtTokenRepo) GenerateAuthToken(ctx context.Context) (string, error) {
	// set claims and sign
	claims := Claims{
		Account:        body.Account,
		Role:           role,
		StandardClaims: jwt.StandardClaims{
			Audience:  body.Account,
			ExpiresAt: now.Add(20 * time.Second).Unix(),
			Id:        jwtId,
			IssuedAt:  now.Unix(),
			Issuer:    "ginJWT",
			NotBefore: now.Add(10 * time.Second).Unix(),
			Subject:   body.Account,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}