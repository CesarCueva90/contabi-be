package router

import (
	"github.com/gin-gonic/gin"
)

func accountancyRoutes(r *gin.Engine, accountancyController AccountancyController) {
	r.GET("/accountancy/clients/supervisor/:supervisor_id", accountancyController.GetClientsBySupervisor)
	r.GET("/accountancy/clients/assignments/:supervisor_id", accountancyController.GetClientAssignmentsMatrix)
	r.PUT("/accountancy/client/:client_id/assignments", accountancyController.UpdateClientAssignments)
	r.GET("/accountancy/clients/responsible/:responsible_id", accountancyController.GetClientsByResonsible)
	r.POST("/accountancy/clients/history/record", accountancyController.CreateClientAccountancyStatusWithAssignments)
	r.PUT("/accountancy/client/:client_id/status/:status_id", accountancyController.UpdateClientAccountancyStatusWithAssignments)
	r.GET("/accountancy/client/:client_id/history", accountancyController.GetClientAccountancyHistory)
	r.GET("/accountancy/clients/all", accountancyController.GetAllClients)
}
