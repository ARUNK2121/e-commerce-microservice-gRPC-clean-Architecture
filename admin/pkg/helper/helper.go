package helper

import (
	"admin/pkg/domain"
	"time"

	"github.com/golang-jwt/jwt"
)

type authCustomClaimsAdmin struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateTokenAdmin(admin domain.Admin) (string, error) {

	claims := &authCustomClaimsAdmin{
		Name:  admin.Name,
		Email: admin.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("12345678"))

	if err != nil {
		return "", err
	}

	return tokenString, nil

}
