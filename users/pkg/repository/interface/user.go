package interfaces

import (
	"users/pkg/domain"
	"users/pkg/utils/models"
)

type UserRepository interface {
	UserSignUp(admin models.UserSignUp) (int, error)
	LoginHandler(adminDetails models.UserLogin) (domain.User, error)
}
