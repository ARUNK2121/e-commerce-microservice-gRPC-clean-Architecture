package interfaces

import (
	"admin/pkg/domain"
	"admin/pkg/utils/models"
)

type AdminRepository interface {
	AdminSignUp(admin models.AdminSignUp) (int, error)
	LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error)
}
