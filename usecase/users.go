package usecase

import (
	"contabi-be/models"

	"golang.org/x/crypto/bcrypt"
)

// UsersInteractor implements the UsersUseCase interface
type UsersInteractor struct {
	usersService UsersService
}

// NewUsersUseCase creates a new instance of UsersService
func NewUsersUseCase(usersService UsersService) UsersService {
	return &UsersInteractor{
		usersService: usersService,
	}
}

// GetUsers retrieves all users
func (uu *UsersInteractor) GetUsers() ([]models.User, error) {
	return uu.usersService.GetUsers()
}

// GetUserInfo retrieves user info by ID
func (uu *UsersInteractor) GetUserByID(userID string) (models.User, error) {
	return uu.usersService.GetUserByID(userID)
}

// CreateUser creates a new user
func (uu *UsersInteractor) CreateUser(user models.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)

	return uu.usersService.CreateUser(user)
}

// UpdateUser updates an user
func (uu *UsersInteractor) UpdateUser(user models.User) error {
	return uu.usersService.UpdateUser(user)
}

// UpdateUserRole updates the user role
func (uu *UsersInteractor) UpdateUserRole(user models.User) error {
	return uu.usersService.UpdateUserRole(user)
}

// PutUserPassword updates the user password
func (uu *UsersInteractor) PutUserPassword(user models.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)

	return uu.usersService.PutUserPassword(user)
}

// DeleteUser deletes an user
func (uu *UsersInteractor) DeleteUser(userID string) error {
	return uu.usersService.DeleteUser(userID)
}

// GetRoles gets all the roles
func (uu *UsersInteractor) GetRoles() ([]models.Role, error) {
	return uu.usersService.GetRoles()
}
