package controller

import (
	"net/http"

	"contabi-be/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// UsersController handles the user-related HTTP requests
type UsersController struct {
	usersUseCase UsersUseCase
	logger       *logrus.Logger
}

// NewUsersController creates a new instance of UsersController
func NewUsersController(usersUseCase UsersUseCase, logger *logrus.Logger) *UsersController {
	return &UsersController{
		usersUseCase: usersUseCase,
		logger:       logger,
	}
}

// GetUsers handles the request to fetch all users
func (uc *UsersController) GetUsers(g *gin.Context) {
	users, err := uc.usersUseCase.GetUsers()
	if err != nil {
		uc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error fetching users")
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}

	g.JSON(http.StatusOK, users)
}

// GetUserByID retrieves user info by ID
func (uc *UsersController) GetUserByID(g *gin.Context) {
	userID := g.Param("id")

	user, err := uc.usersUseCase.GetUserByID(userID)
	if err != nil {
		uc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("GetUserInfo(): error while fetching user info")
		g.JSON(http.StatusBadRequest, gin.H{"error": "error while fetching user info"})
		return
	}

	g.JSON(http.StatusOK, user)
}

// CreateUser creates a new user
func (uc *UsersController) CreateUser(g *gin.Context) {
	var user models.User
	if err := g.ShouldBindJSON(&user); err != nil {
		uc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error binding user data")
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	err := uc.usersUseCase.CreateUser(user)
	if err != nil {
		uc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("CreateUser(): error while creating user")
		g.JSON(http.StatusBadRequest, gin.H{"error": "error while creating user"})
		return
	}

	g.JSON(http.StatusCreated, "user created successfully")
}

// UpdateUser updates an user
func (uc *UsersController) UpdateUser(g *gin.Context) {
	var user models.User
	if err := g.ShouldBindJSON(&user); err != nil {
		uc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error binding user data")
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	err := uc.usersUseCase.UpdateUser(user)
	if err != nil {
		uc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("UpdateUser(): error while updating user")
		g.JSON(http.StatusBadRequest, gin.H{"error": "error while updating user"})
		return
	}

	g.JSON(http.StatusOK, "user updated successfully")
}

// UpdateUserRole updates the user role
func (uc *UsersController) UpdateUserRole(g *gin.Context) {
	var user models.User
	if err := g.ShouldBindJSON(&user); err != nil {
		uc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error binding user data")
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	err := uc.usersUseCase.UpdateUserRole(user)
	if err != nil {
		uc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("UpdateUserRole(): error while updating user role")
		g.JSON(http.StatusBadRequest, gin.H{"error": "error while updating user role"})
		return
	}

	g.JSON(http.StatusOK, "user role updated successfully")
}

// PutUserPassword updates the user password
func (uc *UsersController) PutUserPassword(g *gin.Context) {
	var user models.User
	if err := g.ShouldBindJSON(&user); err != nil {
		uc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error binding user data")
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	err := uc.usersUseCase.PutUserPassword(user)
	if err != nil {
		uc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("PutUserPassword(): error while updating user password")
		g.JSON(http.StatusBadRequest, gin.H{"error": "error while updating user password"})
		return
	}

	g.JSON(http.StatusOK, "user password updated successfully")
}

// DeleteUser deletes an user
func (uc *UsersController) DeleteUser(g *gin.Context) {
	userID := g.Param("id")

	err := uc.usersUseCase.DeleteUser(userID)
	if err != nil {
		uc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("DeleteUser(): error while deleting user")
		g.JSON(http.StatusBadRequest, gin.H{"error": "error while deleting user"})
		return
	}

	g.JSON(http.StatusOK, "user deleted successfully")
}

// DeleteUser deletes an user
func (uc *UsersController) GetRoles(g *gin.Context) {
	roles, err := uc.usersUseCase.GetRoles()
	if err != nil {
		uc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error fetching roles")
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching roles"})
		return
	}

	g.JSON(http.StatusOK, roles)
}
