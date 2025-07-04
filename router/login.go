package router

import (
	"github.com/gin-gonic/gin"
)

// loginRoutes sets the routes for Login
func loginRoutes(r *gin.Engine, loginController LoginController) {
	r.POST("/login", loginController.Login)
}
