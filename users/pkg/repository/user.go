package repository

import (
	"users/pkg/domain"
	interfaces "users/pkg/repository/interface"
	"users/pkg/utils/models"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{
		DB: DB,
	}
}

func (a *userDatabase) UserSignUp(User models.UserSignUp) (int, error) {

	var id int
	if err := a.DB.Raw("insert into users (name,email,phone,password) values (?, ?, ?, ?) RETURNING id", User.Name, User.Email, User.Phone, User.Password).Scan(&id).Error; err != nil {
		return 0, err
	}

	return id, nil

}

func (a *userDatabase) LoginHandler(User models.UserLogin) (domain.User, error) {

	var adminCompareDetails domain.User
	if err := a.DB.Raw("select * from users where email = ? ", User.Email).Scan(&adminCompareDetails).Error; err != nil {
		return domain.User{}, err
	}

	return adminCompareDetails, nil

}
