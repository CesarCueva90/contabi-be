package router

import (
	"contabi-be/middleware"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(g *gin.Context)
}

type UsersController interface {
	GetUsers(g *gin.Context)
	GetUserByID(g *gin.Context)
	CreateUser(g *gin.Context)
	UpdateUser(g *gin.Context)
	UpdateUserRole(g *gin.Context)
	PutUserPassword(g *gin.Context)
	DeleteUser(g *gin.Context)
	GetRoles(g *gin.Context)
}

// NewRouter set the API routes and applies the middleware
func NewRouter(
	loginController LoginController,
	usersController UsersController,

	mw *middleware.Middleware,
) *gin.Engine {
	// Creates a new instance of Gin router
	r := gin.Default()

	// Adds the CORS middleware to all routes
	r.Use(mw.CORS())

	// Routes for Login
	loginRoutes(r, loginController)

	// Adds the authentication middleware to the required routes
	r.Use(mw.AuthMiddleware())

	usersRoutes(r, usersController)

	return r
}
