package router

import (
	"github.com/gin-gonic/gin"
)

func usersRoutes(r *gin.Engine, usersController UsersController) {
	r.GET("/users", usersController.GetUsers)
	r.GET("/user/:id", usersController.GetUserByID)
	r.POST("/user", usersController.CreateUser)
	r.PUT("/user", usersController.UpdateUser)
	r.PUT("/user/role", usersController.UpdateUserRole)
	r.PUT("/user/pass", usersController.PutUserPassword)
	r.DELETE("/user/:id", usersController.DeleteUser)
	r.GET("/roles", usersController.GetRoles)
}
