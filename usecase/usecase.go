package usecase

import "contabi-be/models"

// LoginUseCase defines the interface for login-related operations
type LoginUseCase interface {
	Login(login, password string) (models.User, error)
}

// LoginService defines the interface for login-related operations
type LoginService interface {
	Login(login, password string) (models.User, error)
}

type UsersService interface {
	GetUsers() ([]models.User, error)
	GetUserByID(id string) (models.User, error)
	CreateUser(user models.User) error
	UpdateUser(user models.User) error
	UpdateUserRole(user models.User) error
	PutUserPassword(user models.User) error
	DeleteUser(id string) error
	GetRoles() ([]models.Role, error)
}
