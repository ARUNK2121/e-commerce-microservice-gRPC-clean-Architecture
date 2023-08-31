package domain

type Admin struct {
	ID       uint   `json:"id" gorm:"unique;not null"`
	Name     string `json:"name" gorm:"validate:required"`
	Username string `json:"email" gorm:"validate:required"`
	Password string `json:"password" gorm:"validate:required"`
}
