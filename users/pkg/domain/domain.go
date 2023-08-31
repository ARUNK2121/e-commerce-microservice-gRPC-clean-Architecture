package domain

type User struct {
	ID       uint   `json:"id" gorm:"unique;not null"`
	Name     string `json:"name" gorm:"validate:required"`
	Email    string `json:"email" gorm:"validate:required"`
	Password string `json:"password" gorm:"validate:required"`
	Phone    string `json:"phone" gorm:"validate:required"`
	Blocked  bool   `json:"blocked" gorm:"validate:required"`
}
