package models

type AdminLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"min=8,max=20"`
}

type AdminSignUp struct {
	Name            string `json:"name" binding:"required" gorm:"validate:required"`
	Email           string `json:"email" binding:"required" gorm:"validate:required"`
	Password        string `json:"password" binding:"required" gorm:"validate:required"`
	ConfirmPassword string `json:"confirmpassword" binding:"required"`
}
