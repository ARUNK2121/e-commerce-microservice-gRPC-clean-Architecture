package usecase

import (
	"admin/pkg/helper"
	interfaces "admin/pkg/repository/interface"
	services "admin/pkg/usecase/interface"
	"admin/pkg/utils/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepository interfaces.AdminRepository
}

func NewAdminUseCase(repo interfaces.AdminRepository) services.AdminUseCase {
	return &adminUseCase{
		adminRepository: repo,
	}
}

func (a *adminUseCase) AdminSignUp(adminDetails models.AdminSignUp) (int, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminDetails.Password), 10)
	if err != nil {
		return 500, errors.New("internal server error")
	}
	adminDetails.Password = string(hashedPassword)

	status, err := a.adminRepository.AdminSignUp(adminDetails)
	if err != nil {
		return 400, err
	}

	return status, nil

}

func (a *adminUseCase) AdminLogin(adminDetails models.AdminLogin) (string, error) {

	// getting details of the admin based on the email provided
	adminCompareDetails, err := a.adminRepository.LoginHandler(adminDetails)
	if err != nil {
		return "", err
	}

	// compare password from database and that provided from admins
	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetails.Password))
	if err != nil {
		return "", err
	}

	token, err := helper.GenerateTokenAdmin(adminCompareDetails)

	if err != nil {
		return "", err
	}

	return token, nil
}
