package interfaces

import "api_gateway/pkg/utils/models"

type AdminClient interface {
	AdminLogin(adminDetails models.AdminLogin) (string, error)
	AdminSignUp(admin models.AdminSignUp) (int, error)
}
