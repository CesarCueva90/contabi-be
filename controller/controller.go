package controller

import (
	"contabi-be/models"

	"github.com/sirupsen/logrus"
)

// LoginUseCase
type LoginUseCase interface {
	Login(login, password string) (models.User, error)
}

// UsersUseCase
type UsersUseCase interface {
	GetUsers() ([]models.User, error)
	GetUserByID(id string) (models.User, error)
	CreateUser(user models.User) error
	UpdateUser(user models.User) error
	UpdateUserRole(user models.User) error
	PutUserPassword(user models.User) error
	DeleteUser(id string) error
	GetRoles() ([]models.Role, error)
}

// Controller
type Controller struct {
	lu     LoginUseCase
	uu     UsersUseCase
	logger *logrus.Logger
}

// NewController creates a new Controllert instance
func NewController(lu LoginUseCase, uu UsersUseCase, logger *logrus.Logger) *Controller {
	return &Controller{
		lu:     lu,
		uu:     uu,
		logger: logger,
	}
}
