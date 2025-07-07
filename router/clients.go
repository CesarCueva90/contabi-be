package router

import (
	"github.com/gin-gonic/gin"
)

func clientsRoutes(r *gin.Engine, clientsController ClientsController) {
	clients := r.Group("/clients")
	{
		// Gets all clients with full info
		clients.GET("/", clientsController.GetClientsInfo)

		// Gets only active clients with full info
		clients.GET("/active", clientsController.GetActiveClientsInfo)

		// Gets full info of a specific client
		clients.GET("/:id", clientsController.GetClientInfo)

		// Creates a new client with assignments
		clients.POST("/", clientsController.CreateClient)

		// Updates the basic info of a client
		clients.PUT("/:id", clientsController.UpdateClient)

		// Deactivates a client (soft delete)
		clients.DELETE("/:id", clientsController.DeactivateClient)

		// Activates a client
		clients.PUT("/:id/activate", clientsController.ActivateClient)

		// Updates the assignments of a specific client (supervisor, responsible, emisor)
		clients.PUT("/:id/assignments", clientsController.UpdateClientAssignments)

		// Get clients with pending payments
		clients.GET("/pending-payments", clientsController.GetClientsWithPendingPayments)

		// Updates the payment info of a specific client
		clients.PUT("/:id/payment", clientsController.UpdateClientPayment)
	}
}
