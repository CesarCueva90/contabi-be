package usecase

import (
	"contabi-be/models"
)

// LoginInteractor implements the LoginUseCase interface
type LoginInteractor struct {
	loginService LoginService
}

// NewLoginUseCase creates a new instance of LoginUseCase
func NewLoginUseCase(loginService LoginService) LoginUseCase {
	return &LoginInteractor{
		loginService: loginService,
	}
}

// Login checks credentials and returns the user
func (li *LoginInteractor) Login(login, password string) (models.User, error) {
	user, err := li.loginService.Login(login, password)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
