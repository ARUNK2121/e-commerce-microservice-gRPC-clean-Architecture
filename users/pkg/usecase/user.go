package usecase

import (
	"errors"
	"users/pkg/helper"
	interfaces "users/pkg/repository/interface"
	services "users/pkg/usecase/interface"
	"users/pkg/utils/models"

	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	Repository interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		Repository: repo,
	}
}

func (a *userUseCase) UserSignUp(UserDetails models.UserSignUp) (int, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(UserDetails.Password), 10)
	if err != nil {
		return 500, errors.New("internal server error")
	}
	UserDetails.Password = string(hashedPassword)

	status, err := a.Repository.UserSignUp(UserDetails)
	if err != nil {
		return 400, err
	}

	return status, nil

}

func (a *userUseCase) UserLogin(UserDetails models.UserLogin) (string, error) {

	// getting details of the admin based on the email provided
	Details, err := a.Repository.LoginHandler(UserDetails)
	if err != nil {
		return "", err
	}

	// compare password from database and that provided from admins
	err = bcrypt.CompareHashAndPassword([]byte(Details.Password), []byte(UserDetails.Password))
	if err != nil {
		return "", err
	}

	token, err := helper.GenerateTokenAdmin(Details)

	if err != nil {
		return "", err
	}

	return token, nil
}
