package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoginController handles the login-related HTTP requests
type LoginController struct {
	loginUseCase LoginUseCase
	logger       *logrus.Logger
}

// NewLoginController creates a new instance of LoginController
func NewLoginController(loginUseCase LoginUseCase, logger *logrus.Logger) *LoginController {
	return &LoginController{
		loginUseCase: loginUseCase,
		logger:       logger,
	}
}

// Login handles the request to log in a user
func (lc *LoginController) Login(g *gin.Context) {
	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Bind JSON to struct
	if err := g.ShouldBindJSON(&credentials); err != nil {
		lc.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Login(): Invalid request format")
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Attempt to log in
	user, err := lc.loginUseCase.Login(credentials.Username, credentials.Password)
	if err != nil || user.ID == "" {
		lc.logger.WithFields(logrus.Fields{
			"username": credentials.Username,
		}).Error("Login(): Invalid username or password")
		g.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	lc.logger.WithFields(logrus.Fields{
		"username": user.Username,
	}).Info("Login successful")

	// Return the user (or a token)
	g.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
