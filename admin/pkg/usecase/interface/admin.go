package interfaces

import "admin/pkg/utils/models"

type AdminUseCase interface {
	AdminSignUp(user models.AdminSignUp) (int, error)
	AdminLogin(userSignUpDetails models.AdminLogin) (string, error)
}
