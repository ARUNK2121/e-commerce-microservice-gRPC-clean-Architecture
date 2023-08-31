package interfaces

import "users/pkg/utils/models"

type UserUseCase interface {
	UserSignUp(user models.UserSignUp) (int, error)
	UserLogin(userSignUpDetails models.UserLogin) (string, error)
}
