package router

import (
	"github.com/gin-gonic/gin"
)

func menusRoutes(r *gin.Engine, menusController MenusController) {
	r.GET("/menu/emisors", menusController.GetEmisors)
	r.GET("/menu/supervisors", menusController.GetSupervisors)
	r.GET("/menu/responsibles/:supervisor_id", menusController.GetResponsiblesBySupervisor)
	r.GET("/menu/regimenes", menusController.GetRegimenes)
	r.GET("/menu/accountancy/types", menusController.GetAccountancyTypes)
	r.GET("/menu/accountancy/status", menusController.GetAccountancyStatuses)
}
