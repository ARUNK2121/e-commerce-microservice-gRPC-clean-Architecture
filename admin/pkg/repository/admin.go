package repository

import (
	"admin/pkg/domain"
	interfaces "admin/pkg/repository/interface"
	"admin/pkg/utils/models"

	"gorm.io/gorm"
)

type AdminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &AdminDatabase{
		DB: DB,
	}
}

func (a *AdminDatabase) AdminSignUp(admin models.AdminSignUp) (int, error) {

	var id int
	if err := a.DB.Raw("insert into admins (name,username,password) values (?, ?, ?) RETURNING id", admin.Name, admin.Email, admin.Password).Scan(&id).Error; err != nil {
		return 0, err
	}

	return id, nil

}

func (a *AdminDatabase) LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error) {

	var adminCompareDetails domain.Admin
	if err := a.DB.Raw("select * from admins where username = ? ", adminDetails.Email).Scan(&adminCompareDetails).Error; err != nil {
		return domain.Admin{}, err
	}

	return adminCompareDetails, nil

}
