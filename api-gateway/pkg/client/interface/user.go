package interfaces

import "api_gateway/pkg/utils/models"

type UserClient interface {
	UserLogin(userDetails models.UserLogin) (string, error)
	UserSignUp(user models.UserSignUp) (int, error)
}
