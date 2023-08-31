package models

type UserLogin struct {
	Email    string
	Password string
}

type UserDetails struct {
	ID    uint
	Name  string
	Email string
}

type UserSignUp struct {
	Name            string
	Email           string
	Phone           string
	Password        string
	ConfirmPassword string
}
